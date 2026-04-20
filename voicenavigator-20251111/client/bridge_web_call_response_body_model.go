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
	SetHttpStatusCode(v int32) *BridgeWebCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BridgeWebCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *BridgeWebCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *BridgeWebCallResponseBody
	GetRequestId() *string
}

type BridgeWebCallResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BridgeWebCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *BridgeWebCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BridgeWebCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BridgeWebCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *BridgeWebCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BridgeWebCallResponseBody) SetCode(v string) *BridgeWebCallResponseBody {
	s.Code = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetData(v *BridgeWebCallResponseBodyData) *BridgeWebCallResponseBody {
	s.Data = v
	return s
}

func (s *BridgeWebCallResponseBody) SetHttpStatusCode(v int32) *BridgeWebCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetMessage(v string) *BridgeWebCallResponseBody {
	s.Message = &v
	return s
}

func (s *BridgeWebCallResponseBody) SetParams(v []*string) *BridgeWebCallResponseBody {
	s.Params = v
	return s
}

func (s *BridgeWebCallResponseBody) SetRequestId(v string) *BridgeWebCallResponseBody {
	s.RequestId = &v
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
	// 1515602866088888___5ec8446df1e2495a97969e887ab9e748
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1774794266093
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ServerUrl  *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	// example:
	//
	// f814f3ae-b2a7-44fb-883c-771221274673
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// token
	//
	// example:
	//
	// 57aa3e9b11d2fa5736787cacf0408c1a
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
