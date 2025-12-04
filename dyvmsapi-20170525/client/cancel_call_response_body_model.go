// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelCallResponseBody
	GetCode() *string
	SetMessage(v string) *CancelCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelCallResponseBody
	GetRequestId() *string
	SetStatus(v bool) *CancelCallResponseBody
	GetStatus() *bool
}

type CancelCallResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CancelCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCallResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCallResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *CancelCallResponseBody) SetCode(v string) *CancelCallResponseBody {
	s.Code = &v
	return s
}

func (s *CancelCallResponseBody) SetMessage(v string) *CancelCallResponseBody {
	s.Message = &v
	return s
}

func (s *CancelCallResponseBody) SetRequestId(v string) *CancelCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCallResponseBody) SetStatus(v bool) *CancelCallResponseBody {
	s.Status = &v
	return s
}

func (s *CancelCallResponseBody) Validate() error {
	return dara.Validate(s)
}
