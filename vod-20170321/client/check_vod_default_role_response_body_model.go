// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVodDefaultRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckVodDefaultRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckVodDefaultRoleResponseBody
	GetSuccess() *bool
}

type CheckVodDefaultRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckVodDefaultRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckVodDefaultRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckVodDefaultRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckVodDefaultRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckVodDefaultRoleResponseBody) SetRequestId(v string) *CheckVodDefaultRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckVodDefaultRoleResponseBody) SetSuccess(v bool) *CheckVodDefaultRoleResponseBody {
	s.Success = &v
	return s
}

func (s *CheckVodDefaultRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
