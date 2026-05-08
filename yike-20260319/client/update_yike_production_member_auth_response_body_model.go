// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateYikeProductionMemberAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateYikeProductionMemberAuthResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateYikeProductionMemberAuthResponseBody
	GetResult() *bool
}

type UpdateYikeProductionMemberAuthResponseBody struct {
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

func (s UpdateYikeProductionMemberAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateYikeProductionMemberAuthResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateYikeProductionMemberAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateYikeProductionMemberAuthResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateYikeProductionMemberAuthResponseBody) SetRequestId(v string) *UpdateYikeProductionMemberAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateYikeProductionMemberAuthResponseBody) SetResult(v bool) *UpdateYikeProductionMemberAuthResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateYikeProductionMemberAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
