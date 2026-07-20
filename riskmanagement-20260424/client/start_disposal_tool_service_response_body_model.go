// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDisposalToolServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartDisposalToolServiceResponseBody
	GetCode() *string
	SetMessage(v string) *StartDisposalToolServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartDisposalToolServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartDisposalToolServiceResponseBody
	GetSuccess() *bool
}

type StartDisposalToolServiceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1E0869D6-A5A0-52A6-A924-14070806976C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartDisposalToolServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDisposalToolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDisposalToolServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartDisposalToolServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartDisposalToolServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDisposalToolServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartDisposalToolServiceResponseBody) SetCode(v string) *StartDisposalToolServiceResponseBody {
	s.Code = &v
	return s
}

func (s *StartDisposalToolServiceResponseBody) SetMessage(v string) *StartDisposalToolServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartDisposalToolServiceResponseBody) SetRequestId(v string) *StartDisposalToolServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDisposalToolServiceResponseBody) SetSuccess(v bool) *StartDisposalToolServiceResponseBody {
	s.Success = &v
	return s
}

func (s *StartDisposalToolServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
