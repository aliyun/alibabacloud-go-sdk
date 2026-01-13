// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLink(v string) *CheckAccountTypeResponseBody
	GetLink() *string
	SetMessage(v string) *CheckAccountTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckAccountTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckAccountTypeResponseBody
	GetSuccess() *bool
}

type CheckAccountTypeResponseBody struct {
	// example:
	//
	// http://
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CheckAccountTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountTypeResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountTypeResponseBody) GetLink() *string {
	return s.Link
}

func (s *CheckAccountTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAccountTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAccountTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckAccountTypeResponseBody) SetLink(v string) *CheckAccountTypeResponseBody {
	s.Link = &v
	return s
}

func (s *CheckAccountTypeResponseBody) SetMessage(v string) *CheckAccountTypeResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAccountTypeResponseBody) SetRequestId(v string) *CheckAccountTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAccountTypeResponseBody) SetSuccess(v bool) *CheckAccountTypeResponseBody {
	s.Success = &v
	return s
}

func (s *CheckAccountTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
