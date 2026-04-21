// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetYikeUserRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetYikeUserRoleResponseBody
	GetRequestId() *string
	SetResult(v bool) *SetYikeUserRoleResponseBody
	GetResult() *bool
}

type SetYikeUserRoleResponseBody struct {
	// RequestId
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetYikeUserRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetYikeUserRoleResponseBody) GoString() string {
	return s.String()
}

func (s *SetYikeUserRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetYikeUserRoleResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SetYikeUserRoleResponseBody) SetRequestId(v string) *SetYikeUserRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetYikeUserRoleResponseBody) SetResult(v bool) *SetYikeUserRoleResponseBody {
	s.Result = &v
	return s
}

func (s *SetYikeUserRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
