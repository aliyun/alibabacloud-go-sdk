// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddYikeUserCreditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddYikeUserCreditResponseBody
	GetErrorCode() *string
	SetRequestId(v string) *AddYikeUserCreditResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddYikeUserCreditResponseBody
	GetResult() *bool
}

type AddYikeUserCreditResponseBody struct {
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

func (s AddYikeUserCreditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddYikeUserCreditResponseBody) GoString() string {
	return s.String()
}

func (s *AddYikeUserCreditResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddYikeUserCreditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddYikeUserCreditResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddYikeUserCreditResponseBody) SetErrorCode(v string) *AddYikeUserCreditResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddYikeUserCreditResponseBody) SetRequestId(v string) *AddYikeUserCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddYikeUserCreditResponseBody) SetResult(v bool) *AddYikeUserCreditResponseBody {
	s.Result = &v
	return s
}

func (s *AddYikeUserCreditResponseBody) Validate() error {
	return dara.Validate(s)
}
