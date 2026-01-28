// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeWebCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BridgeWebCallResponseBody
	GetCode() *string
	SetData(v *BridgeWebCallResponseBodyData) *BridgeWebCallResponseBody
	GetData() *BridgeWebCallResponseBodyData
	SetErrorMsg(v string) *BridgeWebCallResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v string) *BridgeWebCallResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *BridgeWebCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BridgeWebCallResponseBody
	GetSuccess() *bool
}

type BridgeWebCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BridgeWebCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// connect timed out
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// CF6D3484-19A1-5C77-863B-AC8B5754D37C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BridgeWebCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BridgeWebCallResponseBody) GoString() string {
	return s.String()
}

func (s *BridgeWebCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *BridgeWebCallResponseBody) GetData() *BridgeWebCallResponseBodyData {
	return s.Data
}

func (s *BridgeWebCallResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *BridgeWebCallResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *BridgeWebCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BridgeWebCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BridgeWebCallResponseBody) SetCode(v string) *BridgeWebCallResponseBody {
	s.Code = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetData(v *BridgeWebCallResponseBodyData) *BridgeWebCallResponseBody {
	s.Data = v
	return s
}

func (s *BridgeWebCallResponseBody) SetErrorMsg(v string) *BridgeWebCallResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetHttpStatusCode(v string) *BridgeWebCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetRequestId(v string) *BridgeWebCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetSuccess(v bool) *BridgeWebCallResponseBody {
	s.Success = &v
	return s
}

func (s *BridgeWebCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BridgeWebCallResponseBodyData struct {
	// example:
	//
	// 894526715106764802
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1744964682422
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// example:
	//
	// i-uf6abxo1tuuwarrtffpp
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// wss://pre-cab-wss.aliyuncs.com:443/audio
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	// example:
	//
	// ws-4b7c263f-9b4c-4b28-baae-a65e9155e380
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 83480f806b48f022313de37b691e167e
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s BridgeWebCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BridgeWebCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *BridgeWebCallResponseBodyData) GetChannelId() *string {
	return s.ChannelId
}

func (s *BridgeWebCallResponseBodyData) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *BridgeWebCallResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BridgeWebCallResponseBodyData) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *BridgeWebCallResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *BridgeWebCallResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *BridgeWebCallResponseBodyData) SetChannelId(v string) *BridgeWebCallResponseBodyData {
	s.ChannelId = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetExpirationTime(v string) *BridgeWebCallResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetInstanceId(v string) *BridgeWebCallResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetServerUrl(v string) *BridgeWebCallResponseBodyData {
	s.ServerUrl = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetSessionId(v string) *BridgeWebCallResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) SetToken(v string) *BridgeWebCallResponseBodyData {
	s.Token = &v
	return s
}

func (s *BridgeWebCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
