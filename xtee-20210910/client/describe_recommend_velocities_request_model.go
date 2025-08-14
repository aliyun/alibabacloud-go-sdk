// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendVelocitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRecommendVelocitiesRequest
	GetLang() *string
	SetCode(v string) *DescribeRecommendVelocitiesRequest
	GetCode() *string
	SetRegId(v string) *DescribeRecommendVelocitiesRequest
	GetRegId() *string
	SetType(v string) *DescribeRecommendVelocitiesRequest
	GetType() *string
}

type DescribeRecommendVelocitiesRequest struct {
	// Sets the language type for requests and responses, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Metric code
	//
	// example:
	//
	// coupon_abuse_detection
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Metric type
	//
	// example:
	//
	// recommend_velocity
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeRecommendVelocitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendVelocitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecommendVelocitiesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRecommendVelocitiesRequest) GetCode() *string {
	return s.Code
}

func (s *DescribeRecommendVelocitiesRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRecommendVelocitiesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeRecommendVelocitiesRequest) SetLang(v string) *DescribeRecommendVelocitiesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecommendVelocitiesRequest) SetCode(v string) *DescribeRecommendVelocitiesRequest {
	s.Code = &v
	return s
}

func (s *DescribeRecommendVelocitiesRequest) SetRegId(v string) *DescribeRecommendVelocitiesRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRecommendVelocitiesRequest) SetType(v string) *DescribeRecommendVelocitiesRequest {
	s.Type = &v
	return s
}

func (s *DescribeRecommendVelocitiesRequest) Validate() error {
	return dara.Validate(s)
}
