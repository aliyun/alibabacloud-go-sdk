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
	// The ID of the request.
	//
	// example:
	//
	// 0AF20EB0-EBBC-4B94-9B84-F3BAFAC53EDE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about whether Security Center is authorized to run configuration checks on cloud services.
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
	return dara.Validate(s)
}

type DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization struct {
	// Indicates whether Security Center is authorized to run configuration checks on cloud services.
	//
	// 	- **0**: no. Security Center is not authorized to run configuration checks on cloud services.
	//
	// 	- **1**: yes. Security Center is authorized to run configuration checks on cloud services.
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
