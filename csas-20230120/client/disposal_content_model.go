// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisposalContent interface {
	dara.Model
	String() string
	GoString() string
	SetAlertContent(v string) *DisposalContent
	GetAlertContent() *string
	SetAlertContentEn(v string) *DisposalContent
	GetAlertContentEn() *string
	SetAlertIntervalSeconds(v int64) *DisposalContent
	GetAlertIntervalSeconds() *int64
	SetAlertTitle(v string) *DisposalContent
	GetAlertTitle() *string
	SetAlertTitleEn(v string) *DisposalContent
	GetAlertTitleEn() *string
	SetNacDemotionPolicyIds(v []*string) *DisposalContent
	GetNacDemotionPolicyIds() []*string
	SetNoticeContent(v string) *DisposalContent
	GetNoticeContent() *string
	SetNoticeContentEn(v string) *DisposalContent
	GetNoticeContentEn() *string
	SetNotifyActions(v []*string) *DisposalContent
	GetNotifyActions() []*string
	SetProhibitActions(v []*string) *DisposalContent
	GetProhibitActions() []*string
	SetProhibitSoftwareIds(v []*string) *DisposalContent
	GetProhibitSoftwareIds() []*string
}

type DisposalContent struct {
	// if can be null:
	// true
	AlertContent         *string `json:"AlertContent,omitempty" xml:"AlertContent,omitempty"`
	AlertContentEn       *string `json:"AlertContentEn,omitempty" xml:"AlertContentEn,omitempty"`
	AlertIntervalSeconds *int64  `json:"AlertIntervalSeconds,omitempty" xml:"AlertIntervalSeconds,omitempty"`
	// if can be null:
	// true
	AlertTitle           *string   `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	AlertTitleEn         *string   `json:"AlertTitleEn,omitempty" xml:"AlertTitleEn,omitempty"`
	NacDemotionPolicyIds []*string `json:"NacDemotionPolicyIds,omitempty" xml:"NacDemotionPolicyIds,omitempty" type:"Repeated"`
	NoticeContent        *string   `json:"NoticeContent,omitempty" xml:"NoticeContent,omitempty"`
	NoticeContentEn      *string   `json:"NoticeContentEn,omitempty" xml:"NoticeContentEn,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// true
	NotifyActions       []*string `json:"NotifyActions,omitempty" xml:"NotifyActions,omitempty" type:"Repeated"`
	ProhibitActions     []*string `json:"ProhibitActions,omitempty" xml:"ProhibitActions,omitempty" type:"Repeated"`
	ProhibitSoftwareIds []*string `json:"ProhibitSoftwareIds,omitempty" xml:"ProhibitSoftwareIds,omitempty" type:"Repeated"`
}

func (s DisposalContent) String() string {
	return dara.Prettify(s)
}

func (s DisposalContent) GoString() string {
	return s.String()
}

func (s *DisposalContent) GetAlertContent() *string {
	return s.AlertContent
}

func (s *DisposalContent) GetAlertContentEn() *string {
	return s.AlertContentEn
}

func (s *DisposalContent) GetAlertIntervalSeconds() *int64 {
	return s.AlertIntervalSeconds
}

func (s *DisposalContent) GetAlertTitle() *string {
	return s.AlertTitle
}

func (s *DisposalContent) GetAlertTitleEn() *string {
	return s.AlertTitleEn
}

func (s *DisposalContent) GetNacDemotionPolicyIds() []*string {
	return s.NacDemotionPolicyIds
}

func (s *DisposalContent) GetNoticeContent() *string {
	return s.NoticeContent
}

func (s *DisposalContent) GetNoticeContentEn() *string {
	return s.NoticeContentEn
}

func (s *DisposalContent) GetNotifyActions() []*string {
	return s.NotifyActions
}

func (s *DisposalContent) GetProhibitActions() []*string {
	return s.ProhibitActions
}

func (s *DisposalContent) GetProhibitSoftwareIds() []*string {
	return s.ProhibitSoftwareIds
}

func (s *DisposalContent) SetAlertContent(v string) *DisposalContent {
	s.AlertContent = &v
	return s
}

func (s *DisposalContent) SetAlertContentEn(v string) *DisposalContent {
	s.AlertContentEn = &v
	return s
}

func (s *DisposalContent) SetAlertIntervalSeconds(v int64) *DisposalContent {
	s.AlertIntervalSeconds = &v
	return s
}

func (s *DisposalContent) SetAlertTitle(v string) *DisposalContent {
	s.AlertTitle = &v
	return s
}

func (s *DisposalContent) SetAlertTitleEn(v string) *DisposalContent {
	s.AlertTitleEn = &v
	return s
}

func (s *DisposalContent) SetNacDemotionPolicyIds(v []*string) *DisposalContent {
	s.NacDemotionPolicyIds = v
	return s
}

func (s *DisposalContent) SetNoticeContent(v string) *DisposalContent {
	s.NoticeContent = &v
	return s
}

func (s *DisposalContent) SetNoticeContentEn(v string) *DisposalContent {
	s.NoticeContentEn = &v
	return s
}

func (s *DisposalContent) SetNotifyActions(v []*string) *DisposalContent {
	s.NotifyActions = v
	return s
}

func (s *DisposalContent) SetProhibitActions(v []*string) *DisposalContent {
	s.ProhibitActions = v
	return s
}

func (s *DisposalContent) SetProhibitSoftwareIds(v []*string) *DisposalContent {
	s.ProhibitSoftwareIds = v
	return s
}

func (s *DisposalContent) Validate() error {
	return dara.Validate(s)
}
