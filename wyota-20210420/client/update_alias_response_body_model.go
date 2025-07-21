// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAliasResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAliasResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAliasResponseBody
	GetRequestId() *string
}

type UpdateAliasResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliasResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAliasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAliasResponseBody) SetCode(v string) *UpdateAliasResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAliasResponseBody) SetMessage(v string) *UpdateAliasResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAliasResponseBody) SetRequestId(v string) *UpdateAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
