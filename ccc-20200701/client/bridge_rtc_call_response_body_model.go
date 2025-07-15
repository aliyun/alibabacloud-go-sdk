// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeRtcCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BridgeRtcCallResponseBody
	GetCode() *string
	SetData(v *BridgeRtcCallResponseBodyData) *BridgeRtcCallResponseBody
	GetData() *BridgeRtcCallResponseBodyData
	SetHttpStatusCode(v int32) *BridgeRtcCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BridgeRtcCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *BridgeRtcCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *BridgeRtcCallResponseBody
	GetRequestId() *string
}

type BridgeRtcCallResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *BridgeRtcCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string                      `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BridgeRtcCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BridgeRtcCallResponseBody) GoString() string {
	return s.String()
}

func (s *BridgeRtcCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *BridgeRtcCallResponseBody) GetData() *BridgeRtcCallResponseBodyData {
	return s.Data
}

func (s *BridgeRtcCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BridgeRtcCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BridgeRtcCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *BridgeRtcCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BridgeRtcCallResponseBody) SetCode(v string) *BridgeRtcCallResponseBody {
	s.Code = &v
	return s
}

func (s *BridgeRtcCallResponseBody) SetData(v *BridgeRtcCallResponseBodyData) *BridgeRtcCallResponseBody {
	s.Data = v
	return s
}

func (s *BridgeRtcCallResponseBody) SetHttpStatusCode(v int32) *BridgeRtcCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BridgeRtcCallResponseBody) SetMessage(v string) *BridgeRtcCallResponseBody {
	s.Message = &v
	return s
}

func (s *BridgeRtcCallResponseBody) SetParams(v []*string) *BridgeRtcCallResponseBody {
	s.Params = v
	return s
}

func (s *BridgeRtcCallResponseBody) SetRequestId(v string) *BridgeRtcCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *BridgeRtcCallResponseBody) Validate() error {
	return dara.Validate(s)
}

type BridgeRtcCallResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	TokenInfo  *string `json:"TokenInfo,omitempty" xml:"TokenInfo,omitempty"`
}

func (s BridgeRtcCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BridgeRtcCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *BridgeRtcCallResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BridgeRtcCallResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *BridgeRtcCallResponseBodyData) GetTokenInfo() *string {
	return s.TokenInfo
}

func (s *BridgeRtcCallResponseBodyData) SetInstanceId(v string) *BridgeRtcCallResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *BridgeRtcCallResponseBodyData) SetJobId(v string) *BridgeRtcCallResponseBodyData {
	s.JobId = &v
	return s
}

func (s *BridgeRtcCallResponseBodyData) SetTokenInfo(v string) *BridgeRtcCallResponseBodyData {
	s.TokenInfo = &v
	return s
}

func (s *BridgeRtcCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
