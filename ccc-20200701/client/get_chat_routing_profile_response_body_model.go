// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatRoutingProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetChatRoutingProfileResponseBody
	GetCode() *string
	SetData(v *GetChatRoutingProfileResponseBodyData) *GetChatRoutingProfileResponseBody
	GetData() *GetChatRoutingProfileResponseBodyData
	SetHttpStatusCode(v int32) *GetChatRoutingProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetChatRoutingProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChatRoutingProfileResponseBody
	GetRequestId() *string
}

type GetChatRoutingProfileResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetChatRoutingProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 30C7D235-DDCF-4C7F-A462-5E2598252C2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatRoutingProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatRoutingProfileResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatRoutingProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatRoutingProfileResponseBody) GetData() *GetChatRoutingProfileResponseBodyData {
	return s.Data
}

func (s *GetChatRoutingProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetChatRoutingProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatRoutingProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatRoutingProfileResponseBody) SetCode(v string) *GetChatRoutingProfileResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatRoutingProfileResponseBody) SetData(v *GetChatRoutingProfileResponseBodyData) *GetChatRoutingProfileResponseBody {
	s.Data = v
	return s
}

func (s *GetChatRoutingProfileResponseBody) SetHttpStatusCode(v int32) *GetChatRoutingProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetChatRoutingProfileResponseBody) SetMessage(v string) *GetChatRoutingProfileResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatRoutingProfileResponseBody) SetRequestId(v string) *GetChatRoutingProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatRoutingProfileResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetChatRoutingProfileResponseBodyData struct {
	// example:
	//
	// {
	//
	//     "AllowExceedingLimitWhenTransferring": false,
	//
	//     "ConcurrencyLimit": 4,
	//
	//     "AllowExceedingLimitWhenClaiming": true,
	//
	//     "Enabled": true
	//
	// }
	AgentConcurrencySettings *string `json:"AgentConcurrencySettings,omitempty" xml:"AgentConcurrencySettings,omitempty"`
	// example:
	//
	// {"IdleChatTimeoutSeconds":300}
	ChatSettings *string `json:"ChatSettings,omitempty" xml:"ChatSettings,omitempty"`
	// example:
	//
	// {
	//
	//     "AgentRingTimeoutSeconds": 30,
	//
	//     "Enabled": true,
	//
	//     "MaxNumberOfConversationsAgentCanMiss": 5,
	//
	//     "PostAgentMissingConversionsAction": "Nothing"
	//
	// }
	DistributionSettings *string `json:"DistributionSettings,omitempty" xml:"DistributionSettings,omitempty"`
	// example:
	//
	// Automatic
	RoutingType *string `json:"RoutingType,omitempty" xml:"RoutingType,omitempty"`
}

func (s GetChatRoutingProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatRoutingProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatRoutingProfileResponseBodyData) GetAgentConcurrencySettings() *string {
	return s.AgentConcurrencySettings
}

func (s *GetChatRoutingProfileResponseBodyData) GetChatSettings() *string {
	return s.ChatSettings
}

func (s *GetChatRoutingProfileResponseBodyData) GetDistributionSettings() *string {
	return s.DistributionSettings
}

func (s *GetChatRoutingProfileResponseBodyData) GetRoutingType() *string {
	return s.RoutingType
}

func (s *GetChatRoutingProfileResponseBodyData) SetAgentConcurrencySettings(v string) *GetChatRoutingProfileResponseBodyData {
	s.AgentConcurrencySettings = &v
	return s
}

func (s *GetChatRoutingProfileResponseBodyData) SetChatSettings(v string) *GetChatRoutingProfileResponseBodyData {
	s.ChatSettings = &v
	return s
}

func (s *GetChatRoutingProfileResponseBodyData) SetDistributionSettings(v string) *GetChatRoutingProfileResponseBodyData {
	s.DistributionSettings = &v
	return s
}

func (s *GetChatRoutingProfileResponseBodyData) SetRoutingType(v string) *GetChatRoutingProfileResponseBodyData {
	s.RoutingType = &v
	return s
}

func (s *GetChatRoutingProfileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
