// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFeatureTrialInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeFeatureTrialInfoResponseBody
	GetCode() *string
	SetFeatureEnabled(v bool) *DescribeFeatureTrialInfoResponseBody
	GetFeatureEnabled() *bool
	SetInTrialPeriod(v bool) *DescribeFeatureTrialInfoResponseBody
	GetInTrialPeriod() *bool
	SetMessage(v string) *DescribeFeatureTrialInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeFeatureTrialInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeFeatureTrialInfoResponseBody
	GetSuccess() *bool
	SetTrialExpireTime(v int64) *DescribeFeatureTrialInfoResponseBody
	GetTrialExpireTime() *int64
}

type DescribeFeatureTrialInfoResponseBody struct {
	// The return code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the feature is activated. The value false is returned if the feature is not activated. In this case, InTrialPeriod is false and TrialExpireTime is 0.
	FeatureEnabled *bool `json:"FeatureEnabled,omitempty" xml:"FeatureEnabled,omitempty"`
	// Indicates whether the feature is within the free trial period. The value false is returned if the feature is not activated or the free trial has expired.
	InTrialPeriod *bool `json:"InTrialPeriod,omitempty" xml:"InTrialPeriod,omitempty"`
	// The returned message. The value "successful" is returned upon success. An error message is returned upon failure.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The expiration time of the free trial. The value is a UNIX timestamp, in seconds. The value 0 is returned if the feature is not activated.
	//
	// example:
	//
	// 1584597600
	TrialExpireTime *int64 `json:"TrialExpireTime,omitempty" xml:"TrialExpireTime,omitempty"`
}

func (s DescribeFeatureTrialInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureTrialInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFeatureTrialInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFeatureTrialInfoResponseBody) GetFeatureEnabled() *bool {
	return s.FeatureEnabled
}

func (s *DescribeFeatureTrialInfoResponseBody) GetInTrialPeriod() *bool {
	return s.InTrialPeriod
}

func (s *DescribeFeatureTrialInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeFeatureTrialInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFeatureTrialInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFeatureTrialInfoResponseBody) GetTrialExpireTime() *int64 {
	return s.TrialExpireTime
}

func (s *DescribeFeatureTrialInfoResponseBody) SetCode(v string) *DescribeFeatureTrialInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFeatureTrialInfoResponseBody) SetFeatureEnabled(v bool) *DescribeFeatureTrialInfoResponseBody {
	s.FeatureEnabled = &v
	return s
}

func (s *DescribeFeatureTrialInfoResponseBody) SetInTrialPeriod(v bool) *DescribeFeatureTrialInfoResponseBody {
	s.InTrialPeriod = &v
	return s
}

func (s *DescribeFeatureTrialInfoResponseBody) SetMessage(v string) *DescribeFeatureTrialInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFeatureTrialInfoResponseBody) SetRequestId(v string) *DescribeFeatureTrialInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFeatureTrialInfoResponseBody) SetSuccess(v bool) *DescribeFeatureTrialInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFeatureTrialInfoResponseBody) SetTrialExpireTime(v int64) *DescribeFeatureTrialInfoResponseBody {
	s.TrialExpireTime = &v
	return s
}

func (s *DescribeFeatureTrialInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
