// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMmAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMmAppResponseBody
	GetSuccess() *bool
}

type UpdateMmAppResponseBody struct {
	// example:
	//
	// xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppResponseBody) SetRequestId(v string) *UpdateMmAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmAppResponseBody) SetSuccess(v bool) *UpdateMmAppResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMmAppResponseBody) Validate() error {
	return dara.Validate(s)
}
