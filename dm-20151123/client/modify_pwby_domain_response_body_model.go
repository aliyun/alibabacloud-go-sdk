// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPWByDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyPWByDomainResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyPWByDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyPWByDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyPWByDomainResponseBody
	GetSuccess() *bool
}

type ModifyPWByDomainResponseBody struct {
	// Status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 02B2A890-CBD8-4806-9BCA-C93190CE7EF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether it was successful
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPWByDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPWByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPWByDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyPWByDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyPWByDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPWByDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyPWByDomainResponseBody) SetCode(v string) *ModifyPWByDomainResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) SetMessage(v string) *ModifyPWByDomainResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) SetRequestId(v string) *ModifyPWByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) SetSuccess(v bool) *ModifyPWByDomainResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
