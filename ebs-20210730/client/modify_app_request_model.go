// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppRequest
	GetAppId() *string
	SetAppName(v string) *ModifyAppRequest
	GetAppName() *string
	SetAppTags(v []*ModifyAppRequestAppTags) *ModifyAppRequest
	GetAppTags() []*ModifyAppRequestAppTags
	SetClientToken(v string) *ModifyAppRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyAppRequest
	GetDescription() *string
	SetOwner(v string) *ModifyAppRequest
	GetOwner() *string
	SetRegionId(v string) *ModifyAppRequest
	GetRegionId() *string
	SetReportSendEnabled(v bool) *ModifyAppRequest
	GetReportSendEnabled() *bool
	SetSubscribePeriod(v string) *ModifyAppRequest
	GetSubscribePeriod() *string
	SetSubscribeStatus(v string) *ModifyAppRequest
	GetSubscribeStatus() *string
}

type ModifyAppRequest struct {
	// example:
	//
	// app-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// App1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	AppTags []*ModifyAppRequestAppTags `json:"AppTags,omitempty" xml:"AppTags,omitempty" type:"Repeated"`
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dev789`
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
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

func (s ModifyAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *ModifyAppRequest) GetAppTags() []*ModifyAppRequestAppTags {
	return s.AppTags
}

func (s *ModifyAppRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAppRequest) GetOwner() *string {
	return s.Owner
}

func (s *ModifyAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAppRequest) GetReportSendEnabled() *bool {
	return s.ReportSendEnabled
}

func (s *ModifyAppRequest) GetSubscribePeriod() *string {
	return s.SubscribePeriod
}

func (s *ModifyAppRequest) GetSubscribeStatus() *string {
	return s.SubscribeStatus
}

func (s *ModifyAppRequest) SetAppId(v string) *ModifyAppRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRequest) SetAppName(v string) *ModifyAppRequest {
	s.AppName = &v
	return s
}

func (s *ModifyAppRequest) SetAppTags(v []*ModifyAppRequestAppTags) *ModifyAppRequest {
	s.AppTags = v
	return s
}

func (s *ModifyAppRequest) SetClientToken(v string) *ModifyAppRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppRequest) SetDescription(v string) *ModifyAppRequest {
	s.Description = &v
	return s
}

func (s *ModifyAppRequest) SetOwner(v string) *ModifyAppRequest {
	s.Owner = &v
	return s
}

func (s *ModifyAppRequest) SetRegionId(v string) *ModifyAppRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAppRequest) SetReportSendEnabled(v bool) *ModifyAppRequest {
	s.ReportSendEnabled = &v
	return s
}

func (s *ModifyAppRequest) SetSubscribePeriod(v string) *ModifyAppRequest {
	s.SubscribePeriod = &v
	return s
}

func (s *ModifyAppRequest) SetSubscribeStatus(v string) *ModifyAppRequest {
	s.SubscribeStatus = &v
	return s
}

func (s *ModifyAppRequest) Validate() error {
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

type ModifyAppRequestAppTags struct {
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ModifyAppRequestAppTags) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRequestAppTags) GoString() string {
	return s.String()
}

func (s *ModifyAppRequestAppTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ModifyAppRequestAppTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ModifyAppRequestAppTags) SetTagKey(v string) *ModifyAppRequestAppTags {
	s.TagKey = &v
	return s
}

func (s *ModifyAppRequestAppTags) SetTagValue(v string) *ModifyAppRequestAppTags {
	s.TagValue = &v
	return s
}

func (s *ModifyAppRequestAppTags) Validate() error {
	return dara.Validate(s)
}
