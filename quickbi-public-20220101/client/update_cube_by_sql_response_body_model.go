// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCubeBySqlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCubeBySqlResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateCubeBySqlResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateCubeBySqlResponseBody
	GetSuccess() *bool
}

type UpdateCubeBySqlResponseBody struct {
	// example:
	//
	// 617277******************ABA47E31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCubeBySqlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCubeBySqlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCubeBySqlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCubeBySqlResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateCubeBySqlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCubeBySqlResponseBody) SetRequestId(v string) *UpdateCubeBySqlResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCubeBySqlResponseBody) SetResult(v bool) *UpdateCubeBySqlResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateCubeBySqlResponseBody) SetSuccess(v bool) *UpdateCubeBySqlResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCubeBySqlResponseBody) Validate() error {
	return dara.Validate(s)
}
