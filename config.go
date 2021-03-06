package main

import (
    "fmt"
    "os"
)

type jsonConfig struct {
    Token      string
    WechatWork struct {
        DefaultReceiverUserId string
        CorpId                string
        CorpSecret            string
        AgentId               string
    }
}

func loadConfig() (config *jsonConfig) {

    config = &jsonConfig{}
    config.Token = os.Getenv("Token")
    config.WechatWork.CorpSecret = os.Getenv("WechatWorkCorpSecret")
    config.WechatWork.CorpId = os.Getenv("WechatWorkCorpId")
    config.WechatWork.DefaultReceiverUserId = os.Getenv("DefaultReceiverUserId")
    config.WechatWork.AgentId = os.Getenv("WechatWorkAgentId")
    fmt.Printf("%+v\n", config)
    return config
}
