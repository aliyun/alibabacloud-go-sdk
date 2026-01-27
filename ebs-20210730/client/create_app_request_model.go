// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateAppRequest
	GetAppName() *string
	SetAppTags(v []*CreateAppRequestAppTags) *CreateAppRequest
	GetAppTags() []*CreateAppRequestAppTags
	SetClientToken(v string) *CreateAppRequest
	GetClientToken() *string
	SetDescription(v string) *CreateAppRequest
	GetDescription() *string
	SetOwner(v string) *CreateAppRequest
	GetOwner() *string
	SetRegionId(v string) *CreateAppRequest
	GetRegionId() *string
	SetReportSendEnabled(v bool) *CreateAppRequest
	GetReportSendEnabled() *bool
	SetSubscribePeriod(v string) *CreateAppRequest
	GetSubscribePeriod() *string
	SetSubscribeStatus(v string) *CreateAppRequest
	GetSubscribeStatus() *string
}

type CreateAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// App1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	AppTags []*CreateAppRequestAppTags `json:"AppTags,omitempty" xml:"AppTags,omitempty" type:"Repeated"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1/0
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// false
	ReportSendEnabled *bool `json:"ReportSendEnabled,omitempty" xml:"ReportSendEnabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Weekly
	SubscribePeriod *string `json:"SubscribePeriod,omitempty" xml:"SubscribePeriod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Subscribe
	SubscribeStatus *string `json:"SubscribeStatus,omitempty" xml:"SubscribeStatus,omitempty"`
}

func (s CreateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppRequest) GetAppTags() []*CreateAppRequestAppTags {
	return s.AppTags
}

func (s *CreateAppRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAppRequest) GetReportSendEnabled() *bool {
	return s.ReportSendEnabled
}

func (s *CreateAppRequest) GetSubscribePeriod() *string {
	return s.SubscribePeriod
}

func (s *CreateAppRequest) GetSubscribeStatus() *string {
	return s.SubscribeStatus
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetAppTags(v []*CreateAppRequestAppTags) *CreateAppRequest {
	s.AppTags = v
	return s
}

func (s *CreateAppRequest) SetClientToken(v string) *CreateAppRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetOwner(v string) *CreateAppRequest {
	s.Owner = &v
	return s
}

func (s *CreateAppRequest) SetRegionId(v string) *CreateAppRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAppRequest) SetReportSendEnabled(v bool) *CreateAppRequest {
	s.ReportSendEnabled = &v
	return s
}

func (s *CreateAppRequest) SetSubscribePeriod(v string) *CreateAppRequest {
	s.SubscribePeriod = &v
	return s
}

func (s *CreateAppRequest) SetSubscribeStatus(v string) *CreateAppRequest {
	s.SubscribeStatus = &v
	return s
}

func (s *CreateAppRequest) Validate() error {
	if s.AppTags != nil {
		for _, item := range s.AppTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppRequestAppTags struct {
	// example:
	//
	// controlledBy
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// ear
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateAppRequestAppTags) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestAppTags) GoString() string {
	return s.String()
}

func (s *CreateAppRequestAppTags) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAppRequestAppTags) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAppRequestAppTags) SetTagKey(v string) *CreateAppRequestAppTags {
	s.TagKey = &v
	return s
}

func (s *CreateAppRequestAppTags) SetTagValue(v string) *CreateAppRequestAppTags {
	s.TagValue = &v
	return s
}

func (s *CreateAppRequestAppTags) Validate() error {
	return dara.Validate(s)
}
