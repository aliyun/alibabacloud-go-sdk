// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateProjectResponseBody
	GetSuccess() *bool
}

type UpdateProjectResponseBody struct {
	// example:
	//
	// 20250923101459e68e3d0b0869e5e4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetSuccess(v bool) *UpdateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
