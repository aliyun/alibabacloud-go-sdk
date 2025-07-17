// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenameFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *RenameFunctionResponseBody
	GetSuccess() *string
}

type RenameFunctionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1ED4C97F-BA2A-57C5-BA7C-8853627EXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameFunctionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *RenameFunctionResponseBody) SetRequestId(v string) *RenameFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameFunctionResponseBody) SetSuccess(v string) *RenameFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *RenameFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
