// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTagResponseBody
	GetSuccess() *bool
}

type UpdateTagResponseBody struct {
	// example:
	//
	// xxxXxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTagResponseBody) SetRequestId(v string) *UpdateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTagResponseBody) SetSuccess(v bool) *UpdateTagResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTagResponseBody) Validate() error {
	return dara.Validate(s)
}
