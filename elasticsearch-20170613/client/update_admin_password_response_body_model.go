// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdminPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAdminPasswordResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateAdminPasswordResponseBody
	GetResult() *bool
}

type UpdateAdminPasswordResponseBody struct {
	// example:
	//
	// 0FA05123-745C-42FD-A69B-AFF48EF9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAdminPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdminPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdminPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAdminPasswordResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateAdminPasswordResponseBody) SetRequestId(v string) *UpdateAdminPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdminPasswordResponseBody) SetResult(v bool) *UpdateAdminPasswordResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateAdminPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
