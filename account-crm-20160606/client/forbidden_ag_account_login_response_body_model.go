// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbiddenAgAccountLoginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ForbiddenAgAccountLoginResponseBody
	GetCode() *string
	SetData(v bool) *ForbiddenAgAccountLoginResponseBody
	GetData() *bool
	SetMessage(v string) *ForbiddenAgAccountLoginResponseBody
	GetMessage() *string
	SetRequestId(v string) *ForbiddenAgAccountLoginResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ForbiddenAgAccountLoginResponseBody
	GetSuccess() *bool
}

type ForbiddenAgAccountLoginResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ForbiddenAgAccountLoginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ForbiddenAgAccountLoginResponseBody) GoString() string {
	return s.String()
}

func (s *ForbiddenAgAccountLoginResponseBody) GetCode() *string {
	return s.Code
}

func (s *ForbiddenAgAccountLoginResponseBody) GetData() *bool {
	return s.Data
}

func (s *ForbiddenAgAccountLoginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ForbiddenAgAccountLoginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ForbiddenAgAccountLoginResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ForbiddenAgAccountLoginResponseBody) SetCode(v string) *ForbiddenAgAccountLoginResponseBody {
	s.Code = &v
	return s
}

func (s *ForbiddenAgAccountLoginResponseBody) SetData(v bool) *ForbiddenAgAccountLoginResponseBody {
	s.Data = &v
	return s
}

func (s *ForbiddenAgAccountLoginResponseBody) SetMessage(v string) *ForbiddenAgAccountLoginResponseBody {
	s.Message = &v
	return s
}

func (s *ForbiddenAgAccountLoginResponseBody) SetRequestId(v string) *ForbiddenAgAccountLoginResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForbiddenAgAccountLoginResponseBody) SetSuccess(v bool) *ForbiddenAgAccountLoginResponseBody {
	s.Success = &v
	return s
}

func (s *ForbiddenAgAccountLoginResponseBody) Validate() error {
	return dara.Validate(s)
}
