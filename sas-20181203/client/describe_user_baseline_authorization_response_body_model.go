// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBaselineAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserBaselineAuthorizationResponseBody
	GetRequestId() *string
	SetUserBaselineAuthorization(v *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) *DescribeUserBaselineAuthorizationResponseBody
	GetUserBaselineAuthorization() *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization
}

type DescribeUserBaselineAuthorizationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AF20EB0-EBBC-4B94-9B84-F3BAFAC53EDE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The authorization information for cloud baseline configuration check.
	UserBaselineAuthorization *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization `json:"UserBaselineAuthorization,omitempty" xml:"UserBaselineAuthorization,omitempty" type:"Struct"`
}

func (s DescribeUserBaselineAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBaselineAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBaselineAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserBaselineAuthorizationResponseBody) GetUserBaselineAuthorization() *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization {
	return s.UserBaselineAuthorization
}

func (s *DescribeUserBaselineAuthorizationResponseBody) SetRequestId(v string) *DescribeUserBaselineAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBaselineAuthorizationResponseBody) SetUserBaselineAuthorization(v *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) *DescribeUserBaselineAuthorizationResponseBody {
	s.UserBaselineAuthorization = v
	return s
}

func (s *DescribeUserBaselineAuthorizationResponseBody) Validate() error {
	if s.UserBaselineAuthorization != nil {
		if err := s.UserBaselineAuthorization.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization struct {
	// The authorization status of the cloud platform configuration check. Valid values:
	//
	// - **0**: Authorization is disabled. If authorization is disabled, you cannot use the cloud platform configuration check feature.
	//
	// - **1**: Authorization is enabled. If authorization is enabled, you can use the cloud platform configuration check feature.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) GoString() string {
	return s.String()
}

func (s *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) SetStatus(v int32) *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization {
	s.Status = &v
	return s
}

func (s *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) Validate() error {
	return dara.Validate(s)
}
