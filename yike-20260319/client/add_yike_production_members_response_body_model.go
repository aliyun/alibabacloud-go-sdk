// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddYikeProductionMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddYikeProductionMembersResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddYikeProductionMembersResponseBody
	GetResult() *bool
}

type AddYikeProductionMembersResponseBody struct {
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

func (s AddYikeProductionMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddYikeProductionMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddYikeProductionMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddYikeProductionMembersResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddYikeProductionMembersResponseBody) SetRequestId(v string) *AddYikeProductionMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddYikeProductionMembersResponseBody) SetResult(v bool) *AddYikeProductionMembersResponseBody {
	s.Result = &v
	return s
}

func (s *AddYikeProductionMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
