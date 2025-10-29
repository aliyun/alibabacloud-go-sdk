// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateComponentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateComponentResponseBody
	GetSuccess() *bool
}

type UpdateComponentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComponentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateComponentResponseBody) SetRequestId(v string) *UpdateComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComponentResponseBody) SetSuccess(v bool) *UpdateComponentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
