// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckProjectHasDependencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckProjectHasDependencyResponseBody
	GetCode() *string
	SetData(v bool) *CheckProjectHasDependencyResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CheckProjectHasDependencyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckProjectHasDependencyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckProjectHasDependencyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckProjectHasDependencyResponseBody
	GetSuccess() *bool
}

type CheckProjectHasDependencyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckProjectHasDependencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckProjectHasDependencyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckProjectHasDependencyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckProjectHasDependencyResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckProjectHasDependencyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckProjectHasDependencyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckProjectHasDependencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckProjectHasDependencyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckProjectHasDependencyResponseBody) SetCode(v string) *CheckProjectHasDependencyResponseBody {
	s.Code = &v
	return s
}

func (s *CheckProjectHasDependencyResponseBody) SetData(v bool) *CheckProjectHasDependencyResponseBody {
	s.Data = &v
	return s
}

func (s *CheckProjectHasDependencyResponseBody) SetHttpStatusCode(v int32) *CheckProjectHasDependencyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckProjectHasDependencyResponseBody) SetMessage(v string) *CheckProjectHasDependencyResponseBody {
	s.Message = &v
	return s
}

func (s *CheckProjectHasDependencyResponseBody) SetRequestId(v string) *CheckProjectHasDependencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckProjectHasDependencyResponseBody) SetSuccess(v bool) *CheckProjectHasDependencyResponseBody {
	s.Success = &v
	return s
}

func (s *CheckProjectHasDependencyResponseBody) Validate() error {
	return dara.Validate(s)
}
