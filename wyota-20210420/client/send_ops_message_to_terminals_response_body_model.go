// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendOpsMessageToTerminalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendOpsMessageToTerminalsResponseBody
	GetCode() *string
	SetData(v []*SendOpsMessageToTerminalsResponseBodyData) *SendOpsMessageToTerminalsResponseBody
	GetData() []*SendOpsMessageToTerminalsResponseBodyData
	SetHttpStatusCode(v int32) *SendOpsMessageToTerminalsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SendOpsMessageToTerminalsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendOpsMessageToTerminalsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendOpsMessageToTerminalsResponseBody
	GetSuccess() *bool
}

type SendOpsMessageToTerminalsResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*SendOpsMessageToTerminalsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendOpsMessageToTerminalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendOpsMessageToTerminalsResponseBody) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendOpsMessageToTerminalsResponseBody) GetData() []*SendOpsMessageToTerminalsResponseBodyData {
	return s.Data
}

func (s *SendOpsMessageToTerminalsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SendOpsMessageToTerminalsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendOpsMessageToTerminalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendOpsMessageToTerminalsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendOpsMessageToTerminalsResponseBody) SetCode(v string) *SendOpsMessageToTerminalsResponseBody {
	s.Code = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) SetData(v []*SendOpsMessageToTerminalsResponseBodyData) *SendOpsMessageToTerminalsResponseBody {
	s.Data = v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) SetHttpStatusCode(v int32) *SendOpsMessageToTerminalsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) SetMessage(v string) *SendOpsMessageToTerminalsResponseBody {
	s.Message = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) SetRequestId(v string) *SendOpsMessageToTerminalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) SetSuccess(v bool) *SendOpsMessageToTerminalsResponseBody {
	s.Success = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBody) Validate() error {
	return dara.Validate(s)
}

type SendOpsMessageToTerminalsResponseBodyData struct {
	FailReason   *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s SendOpsMessageToTerminalsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendOpsMessageToTerminalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalsResponseBodyData) GetFailReason() *string {
	return s.FailReason
}

func (s *SendOpsMessageToTerminalsResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *SendOpsMessageToTerminalsResponseBodyData) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *SendOpsMessageToTerminalsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SendOpsMessageToTerminalsResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *SendOpsMessageToTerminalsResponseBodyData) SetFailReason(v string) *SendOpsMessageToTerminalsResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBodyData) SetResult(v string) *SendOpsMessageToTerminalsResponseBodyData {
	s.Result = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBodyData) SetSerialNumber(v string) *SendOpsMessageToTerminalsResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBodyData) SetSuccess(v bool) *SendOpsMessageToTerminalsResponseBodyData {
	s.Success = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBodyData) SetUuid(v string) *SendOpsMessageToTerminalsResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
