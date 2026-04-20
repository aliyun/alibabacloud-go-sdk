// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmGlobalAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertConfig(v *DescribeCloudGtmGlobalAlertResponseBodyAlertConfig) *DescribeCloudGtmGlobalAlertResponseBody
	GetAlertConfig() *DescribeCloudGtmGlobalAlertResponseBodyAlertConfig
	SetAlertGroup(v *DescribeCloudGtmGlobalAlertResponseBodyAlertGroup) *DescribeCloudGtmGlobalAlertResponseBody
	GetAlertGroup() *DescribeCloudGtmGlobalAlertResponseBodyAlertGroup
	SetRequestId(v string) *DescribeCloudGtmGlobalAlertResponseBody
	GetRequestId() *string
}

type DescribeCloudGtmGlobalAlertResponseBody struct {
	AlertConfig *DescribeCloudGtmGlobalAlertResponseBodyAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	AlertGroup  *DescribeCloudGtmGlobalAlertResponseBodyAlertGroup  `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudGtmGlobalAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmGlobalAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmGlobalAlertResponseBody) GetAlertConfig() *DescribeCloudGtmGlobalAlertResponseBodyAlertConfig {
	return s.AlertConfig
}

func (s *DescribeCloudGtmGlobalAlertResponseBody) GetAlertGroup() *DescribeCloudGtmGlobalAlertResponseBodyAlertGroup {
	return s.AlertGroup
}

func (s *DescribeCloudGtmGlobalAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmGlobalAlertResponseBody) SetAlertConfig(v *DescribeCloudGtmGlobalAlertResponseBodyAlertConfig) *DescribeCloudGtmGlobalAlertResponseBody {
	s.AlertConfig = v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBody) SetAlertGroup(v *DescribeCloudGtmGlobalAlertResponseBodyAlertGroup) *DescribeCloudGtmGlobalAlertResponseBody {
	s.AlertGroup = v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBody) SetRequestId(v string) *DescribeCloudGtmGlobalAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBody) Validate() error {
	if s.AlertConfig != nil {
		if err := s.AlertConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AlertGroup != nil {
		if err := s.AlertGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudGtmGlobalAlertResponseBodyAlertConfig struct {
	AlertConfig []*DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmGlobalAlertResponseBodyAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmGlobalAlertResponseBodyAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfig) GetAlertConfig() []*DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig {
	return s.AlertConfig
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfig) SetAlertConfig(v []*DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) *DescribeCloudGtmGlobalAlertResponseBodyAlertConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfig) Validate() error {
	if s.AlertConfig != nil {
		for _, item := range s.AlertConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig struct {
	DingtalkNotice *bool   `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	EmailNotice    *bool   `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	NoticeType     *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// example:
	//
	// 10
	QpsThreshold *int64 `json:"QpsThreshold,omitempty" xml:"QpsThreshold,omitempty"`
	SmsNotice    *bool  `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
	// example:
	//
	// 50
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) GetDingtalkNotice() *bool {
	return s.DingtalkNotice
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) GetEmailNotice() *bool {
	return s.EmailNotice
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) GetNoticeType() *string {
	return s.NoticeType
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) GetQpsThreshold() *int64 {
	return s.QpsThreshold
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) GetSmsNotice() *bool {
	return s.SmsNotice
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) GetThreshold() *int32 {
	return s.Threshold
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) SetDingtalkNotice(v bool) *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig {
	s.DingtalkNotice = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) SetEmailNotice(v bool) *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig {
	s.EmailNotice = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) SetNoticeType(v string) *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) SetQpsThreshold(v int64) *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig {
	s.QpsThreshold = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) SetSmsNotice(v bool) *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) SetThreshold(v int32) *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertConfigAlertConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmGlobalAlertResponseBodyAlertGroup struct {
	AlertGroup []*string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmGlobalAlertResponseBodyAlertGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmGlobalAlertResponseBodyAlertGroup) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertGroup) GetAlertGroup() []*string {
	return s.AlertGroup
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertGroup) SetAlertGroup(v []*string) *DescribeCloudGtmGlobalAlertResponseBodyAlertGroup {
	s.AlertGroup = v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponseBodyAlertGroup) Validate() error {
	return dara.Validate(s)
}
