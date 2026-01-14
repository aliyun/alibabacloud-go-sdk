// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserByMobileAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryUserByMobileAccountResponseBody
	GetRequestId() *string
	SetResult(v *QueryUserByMobileAccountResponseBodyResult) *QueryUserByMobileAccountResponseBody
	GetResult() *QueryUserByMobileAccountResponseBodyResult
	SetSuccess(v bool) *QueryUserByMobileAccountResponseBody
	GetSuccess() *bool
}

type QueryUserByMobileAccountResponseBody struct {
	// example:
	//
	// 46e53***********270
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryUserByMobileAccountResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserByMobileAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserByMobileAccountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserByMobileAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserByMobileAccountResponseBody) GetResult() *QueryUserByMobileAccountResponseBodyResult {
	return s.Result
}

func (s *QueryUserByMobileAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserByMobileAccountResponseBody) SetRequestId(v string) *QueryUserByMobileAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserByMobileAccountResponseBody) SetResult(v *QueryUserByMobileAccountResponseBodyResult) *QueryUserByMobileAccountResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserByMobileAccountResponseBody) SetSuccess(v bool) *QueryUserByMobileAccountResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserByMobileAccountResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryUserByMobileAccountResponseBodyResult struct {
	// example:
	//
	// test
	BoundUserId *string `json:"BoundUserId,omitempty" xml:"BoundUserId,omitempty"`
	// example:
	//
	// test
	ThirdAccountName *string `json:"ThirdAccountName,omitempty" xml:"ThirdAccountName,omitempty"`
}

func (s QueryUserByMobileAccountResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserByMobileAccountResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserByMobileAccountResponseBodyResult) GetBoundUserId() *string {
	return s.BoundUserId
}

func (s *QueryUserByMobileAccountResponseBodyResult) GetThirdAccountName() *string {
	return s.ThirdAccountName
}

func (s *QueryUserByMobileAccountResponseBodyResult) SetBoundUserId(v string) *QueryUserByMobileAccountResponseBodyResult {
	s.BoundUserId = &v
	return s
}

func (s *QueryUserByMobileAccountResponseBodyResult) SetThirdAccountName(v string) *QueryUserByMobileAccountResponseBodyResult {
	s.ThirdAccountName = &v
	return s
}

func (s *QueryUserByMobileAccountResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
