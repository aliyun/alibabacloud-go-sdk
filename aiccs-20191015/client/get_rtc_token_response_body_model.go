// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRtcTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRtcTokenResponseBody
	GetCode() *string
	SetData(v *GetRtcTokenResponseBodyData) *GetRtcTokenResponseBody
	GetData() *GetRtcTokenResponseBodyData
	SetMessage(v string) *GetRtcTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRtcTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRtcTokenResponseBody
	GetSuccess() *bool
}

type GetRtcTokenResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"cleansession":true,"clientId":"GID_VOIP@@@ClientId_2000000001271771_100685295007","conferenceTopic":"cs_alicom_voip_conference_pre","dataTopic":"alicom_voip_data_pre","host":"mqtt-cn-4590mdhb901.mqtt.aliyuncs.com","meetingEventKeepAliveInterval":0,"phoneTopic":"alicom_voip_phone","port":0,"reconnectTimeout":2000,"registerTime":0,"sdkClientPort":8883,"serverId":"GID_VOIP@@@MTEuMjIuMTQ1Ljgy","sgwServerTopic":"alicom_voip_server_pre","tlsport":443,"tokenData":"LzMT+XLFl5s/YWJ/MlDz4t/Lq5HC1iGU1P28HAMaxYzmBSHQsWXgdISJ1ZJ+2cxaU0jwYsoyG8Q8cCIbLZTwwaFHf7gc7pPXbJGYgJWUr5ooKsoHaVvvG34cww7W8woWE1OsmZGFDODvooOIjF1CZSorVrR8OwRdprW99yqhMhkJKh7r5f3HfiQgoJWL8b3A85RrRGCSP057skgQ5rIqVAlx7jDFrOTdLtz+krken8qYvpaVBO9wRfFyWQLvxAgNJNx3Oql/hpzc2o3+xbKGTA/P2siLn6Nee1FYk5ClpXcnvfSTM4BAauuWR+oES10VblKEKTU5R/pfjXj3UKOlZ1+OdGO93WA16BR/l1uRb3cOLqya5pjWM+oSmo0sOR7B0ATLz6K1xA0Pc+p6Mu8hZl+OdGO93WA1kj1L0h9Z6CAZG4ol/BNdIg9z6noy7yFmX450Y73dYDWSPUvSH1noIBkbiiX8E10iT0a6ypXxupQdyyrhh8j1yw6otqnw5AT5Tj5VsVWVfflXeZ8nPt1ydXC9nWeHX7K80O6vtOU9M8Qn5VrhkP0F1umbOoYs3NfM+WYZIQx4pkViQo6qqkxgbEg1l6oHJbmOVzrxYRrDTIxqgE/pb0YVJQ==","useTLS":false,"username":"LTAI27GqAW1VrcQA"}
	Data *GetRtcTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRtcTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRtcTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetRtcTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRtcTokenResponseBody) GetData() *GetRtcTokenResponseBodyData {
	return s.Data
}

func (s *GetRtcTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRtcTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRtcTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRtcTokenResponseBody) SetCode(v string) *GetRtcTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetData(v *GetRtcTokenResponseBodyData) *GetRtcTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetRtcTokenResponseBody) SetMessage(v string) *GetRtcTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetRequestId(v string) *GetRtcTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetSuccess(v bool) *GetRtcTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetRtcTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRtcTokenResponseBodyData struct {
	// example:
	//
	// 139000000
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 200000000*******
	RtcId *string `json:"RtcId,omitempty" xml:"RtcId,omitempty"`
	// example:
	//
	// {“cleansession”:true,”clientId”:”GID_VOIP@@@ClientId_****”,”conferenceTopic”:”cs_alicom_voip_conference”,”host”:”mqtt-cn-4590mdhb901.mqtt.aliyuncs.com”,”meetingEventKeepAliveInterval”:0,”phoneTopic”:”alicom_voip_phone”,”port”:0,”reconnectTimeout”:2000,”registerTime”:0,”sdkClientPort”:8883,”serverId”:”GID_VOIP@@@MTEuMTMuMTM2LjExOA==”,”sgwServerTopic”:”alicom_voip_server_pre”,”tlsport”:443,”tokenData”:”abcdef”,”useTLS”:false}
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetRtcTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRtcTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRtcTokenResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *GetRtcTokenResponseBodyData) GetRtcId() *string {
	return s.RtcId
}

func (s *GetRtcTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *GetRtcTokenResponseBodyData) SetAccountName(v string) *GetRtcTokenResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetRtcTokenResponseBodyData) SetRtcId(v string) *GetRtcTokenResponseBodyData {
	s.RtcId = &v
	return s
}

func (s *GetRtcTokenResponseBodyData) SetToken(v string) *GetRtcTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetRtcTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
