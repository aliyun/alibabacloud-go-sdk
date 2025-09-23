// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReConfigApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReConfigApplicationResponseBody
	GetCode() *string
	SetMessage(v string) *ReConfigApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReConfigApplicationResponseBody
	GetRequestId() *string
}

type ReConfigApplicationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9656C816-1E9A-58D2-86D5-710678D61AF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReConfigApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReConfigApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ReConfigApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReConfigApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReConfigApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReConfigApplicationResponseBody) SetCode(v string) *ReConfigApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ReConfigApplicationResponseBody) SetMessage(v string) *ReConfigApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ReConfigApplicationResponseBody) SetRequestId(v string) *ReConfigApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReConfigApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
