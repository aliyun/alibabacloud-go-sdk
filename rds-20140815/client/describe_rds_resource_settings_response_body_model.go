// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsResourceSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRdsInstanceResourceSettings(v *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings) *DescribeRdsResourceSettingsResponseBody
	GetRdsInstanceResourceSettings() *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings
	SetRequestId(v string) *DescribeRdsResourceSettingsResponseBody
	GetRequestId() *string
}

type DescribeRdsResourceSettingsResponseBody struct {
	// The details about notification settings for an instance.
	RdsInstanceResourceSettings *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings `json:"RdsInstanceResourceSettings,omitempty" xml:"RdsInstanceResourceSettings,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 76364A52-E0AB-5CC8-xxxx-CF1DC482C092
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRdsResourceSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsResourceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsResourceSettingsResponseBody) GetRdsInstanceResourceSettings() *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings {
	return s.RdsInstanceResourceSettings
}

func (s *DescribeRdsResourceSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdsResourceSettingsResponseBody) SetRdsInstanceResourceSettings(v *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings) *DescribeRdsResourceSettingsResponseBody {
	s.RdsInstanceResourceSettings = v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBody) SetRequestId(v string) *DescribeRdsResourceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBody) Validate() error {
	if s.RdsInstanceResourceSettings != nil {
		if err := s.RdsInstanceResourceSettings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings struct {
	RdsInstanceResourceSetting []*DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting `json:"RdsInstanceResourceSetting,omitempty" xml:"RdsInstanceResourceSetting,omitempty" type:"Repeated"`
}

func (s DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings) GoString() string {
	return s.String()
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings) GetRdsInstanceResourceSetting() []*DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	return s.RdsInstanceResourceSetting
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings) SetRdsInstanceResourceSetting(v []*DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings {
	s.RdsInstanceResourceSetting = v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettings) Validate() error {
	if s.RdsInstanceResourceSetting != nil {
		for _, item := range s.RdsInstanceResourceSetting {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting struct {
	// The end date.
	//
	// example:
	//
	// 2024-04-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Specifies whether to pin the notification at the top.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTop *string `json:"IsTop,omitempty" xml:"IsTop,omitempty"`
	// The notification text.
	NoticeBarContent *string `json:"NoticeBarContent,omitempty" xml:"NoticeBarContent,omitempty"`
	// The text of the popup button.
	PoppedUpButtonText *string `json:"PoppedUpButtonText,omitempty" xml:"PoppedUpButtonText,omitempty"`
	// The type of the popup button.
	//
	// 	- BUY
	//
	// 	- RENEW
	//
	// 	- UPGRADE
	PoppedUpButtonType *string `json:"PoppedUpButtonType,omitempty" xml:"PoppedUpButtonType,omitempty"`
	// The link of the popup button.
	//
	// example:
	//
	// https://www.example.com/subscribe
	PoppedUpButtonUrl *string `json:"PoppedUpButtonUrl,omitempty" xml:"PoppedUpButtonUrl,omitempty"`
	// The text of the popup.
	PoppedUpContent *string `json:"PoppedUpContent,omitempty" xml:"PoppedUpContent,omitempty"`
	// The location of the notification.
	ResourceNiche *string `json:"ResourceNiche,omitempty" xml:"ResourceNiche,omitempty"`
	// The effective date.
	//
	// example:
	//
	// 2024-04-21
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GoString() string {
	return s.String()
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GetIsTop() *string {
	return s.IsTop
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GetNoticeBarContent() *string {
	return s.NoticeBarContent
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GetPoppedUpButtonText() *string {
	return s.PoppedUpButtonText
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GetPoppedUpButtonType() *string {
	return s.PoppedUpButtonType
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GetPoppedUpButtonUrl() *string {
	return s.PoppedUpButtonUrl
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GetPoppedUpContent() *string {
	return s.PoppedUpContent
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GetResourceNiche() *string {
	return s.ResourceNiche
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) SetEndDate(v string) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	s.EndDate = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) SetIsTop(v string) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	s.IsTop = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) SetNoticeBarContent(v string) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	s.NoticeBarContent = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) SetPoppedUpButtonText(v string) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	s.PoppedUpButtonText = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) SetPoppedUpButtonType(v string) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	s.PoppedUpButtonType = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) SetPoppedUpButtonUrl(v string) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	s.PoppedUpButtonUrl = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) SetPoppedUpContent(v string) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	s.PoppedUpContent = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) SetResourceNiche(v string) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	s.ResourceNiche = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) SetStartDate(v string) *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting {
	s.StartDate = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponseBodyRdsInstanceResourceSettingsRdsInstanceResourceSetting) Validate() error {
	return dara.Validate(s)
}
