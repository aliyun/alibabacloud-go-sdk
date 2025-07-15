// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAttributeForConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceVpcAttributeForConsoleResponseBody
	GetRequestId() *string
}

type ModifyInstanceVpcAttributeForConsoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D08741CF-BE59-5DA6-B06F-BB65173110C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceVpcAttributeForConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAttributeForConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAttributeForConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceVpcAttributeForConsoleResponseBody) SetRequestId(v string) *ModifyInstanceVpcAttributeForConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}
