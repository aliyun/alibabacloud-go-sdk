// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTerminalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddTerminalsResponseBody
	GetCode() *string
	SetData(v []*AddTerminalsResponseBodyData) *AddTerminalsResponseBody
	GetData() []*AddTerminalsResponseBodyData
	SetHttpStatusCode(v int32) *AddTerminalsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddTerminalsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddTerminalsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTerminalsResponseBody
	GetSuccess() *bool
}

type AddTerminalsResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*AddTerminalsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTerminalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTerminalsResponseBody) GoString() string {
	return s.String()
}

func (s *AddTerminalsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddTerminalsResponseBody) GetData() []*AddTerminalsResponseBodyData {
	return s.Data
}

func (s *AddTerminalsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddTerminalsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddTerminalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTerminalsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTerminalsResponseBody) SetCode(v string) *AddTerminalsResponseBody {
	s.Code = &v
	return s
}

func (s *AddTerminalsResponseBody) SetData(v []*AddTerminalsResponseBodyData) *AddTerminalsResponseBody {
	s.Data = v
	return s
}

func (s *AddTerminalsResponseBody) SetHttpStatusCode(v int32) *AddTerminalsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddTerminalsResponseBody) SetMessage(v string) *AddTerminalsResponseBody {
	s.Message = &v
	return s
}

func (s *AddTerminalsResponseBody) SetRequestId(v string) *AddTerminalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTerminalsResponseBody) SetSuccess(v bool) *AddTerminalsResponseBody {
	s.Success = &v
	return s
}

func (s *AddTerminalsResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddTerminalsResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s AddTerminalsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddTerminalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddTerminalsResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddTerminalsResponseBodyData) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *AddTerminalsResponseBodyData) SetErrorCode(v string) *AddTerminalsResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *AddTerminalsResponseBodyData) SetSerialNumber(v string) *AddTerminalsResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *AddTerminalsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
