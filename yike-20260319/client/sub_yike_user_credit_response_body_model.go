// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubYikeUserCreditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubYikeUserCreditResponseBody
	GetErrorCode() *string
	SetRequestId(v string) *SubYikeUserCreditResponseBody
	GetRequestId() *string
	SetResult(v bool) *SubYikeUserCreditResponseBody
	GetResult() *bool
}

type SubYikeUserCreditResponseBody struct {
	// example:
	//
	// NOT_ENOUGH_ALLOCATE_CREDIT_QUOTA
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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

func (s SubYikeUserCreditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubYikeUserCreditResponseBody) GoString() string {
	return s.String()
}

func (s *SubYikeUserCreditResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubYikeUserCreditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubYikeUserCreditResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SubYikeUserCreditResponseBody) SetErrorCode(v string) *SubYikeUserCreditResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubYikeUserCreditResponseBody) SetRequestId(v string) *SubYikeUserCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubYikeUserCreditResponseBody) SetResult(v bool) *SubYikeUserCreditResponseBody {
	s.Result = &v
	return s
}

func (s *SubYikeUserCreditResponseBody) Validate() error {
	return dara.Validate(s)
}
