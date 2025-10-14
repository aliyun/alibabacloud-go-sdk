// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewGtmRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *PreviewGtmRecoveryPlanResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *PreviewGtmRecoveryPlanResponseBody
	GetPageSize() *int32
	SetPreviews(v *PreviewGtmRecoveryPlanResponseBodyPreviews) *PreviewGtmRecoveryPlanResponseBody
	GetPreviews() *PreviewGtmRecoveryPlanResponseBodyPreviews
	SetRequestId(v string) *PreviewGtmRecoveryPlanResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *PreviewGtmRecoveryPlanResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *PreviewGtmRecoveryPlanResponseBody
	GetTotalPages() *int32
}

type PreviewGtmRecoveryPlanResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned preview information of the disaster recovery plan.
	Previews *PreviewGtmRecoveryPlanResponseBodyPreviews `json:"Previews,omitempty" xml:"Previews,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 853805EA-3D47-47D5-9A1A-A45C24313ABD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned on all pages.
	//
	// example:
	//
	// 15
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 3
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s PreviewGtmRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *PreviewGtmRecoveryPlanResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PreviewGtmRecoveryPlanResponseBody) GetPreviews() *PreviewGtmRecoveryPlanResponseBodyPreviews {
	return s.Previews
}

func (s *PreviewGtmRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewGtmRecoveryPlanResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *PreviewGtmRecoveryPlanResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetPageNumber(v int32) *PreviewGtmRecoveryPlanResponseBody {
	s.PageNumber = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetPageSize(v int32) *PreviewGtmRecoveryPlanResponseBody {
	s.PageSize = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetPreviews(v *PreviewGtmRecoveryPlanResponseBodyPreviews) *PreviewGtmRecoveryPlanResponseBody {
	s.Previews = v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetRequestId(v string) *PreviewGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetTotalItems(v int32) *PreviewGtmRecoveryPlanResponseBody {
	s.TotalItems = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) SetTotalPages(v int32) *PreviewGtmRecoveryPlanResponseBody {
	s.TotalPages = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBody) Validate() error {
	if s.Previews != nil {
		if err := s.Previews.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewGtmRecoveryPlanResponseBodyPreviews struct {
	Preview []*PreviewGtmRecoveryPlanResponseBodyPreviewsPreview `json:"Preview,omitempty" xml:"Preview,omitempty" type:"Repeated"`
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviews) String() string {
	return dara.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviews) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviews) GetPreview() []*PreviewGtmRecoveryPlanResponseBodyPreviewsPreview {
	return s.Preview
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviews) SetPreview(v []*PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) *PreviewGtmRecoveryPlanResponseBodyPreviews {
	s.Preview = v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviews) Validate() error {
	if s.Preview != nil {
		for _, item := range s.Preview {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreviewGtmRecoveryPlanResponseBodyPreviewsPreview struct {
	// The ID of the GTM instance to which the previewed disaster recovery plan belongs.
	//
	// example:
	//
	// instance-example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the GTM instance to which the previewed disaster recovery plan belongs.
	//
	// example:
	//
	// name-example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The returned information of the switching policies for address pools.
	SwitchInfos *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos `json:"SwitchInfos,omitempty" xml:"SwitchInfos,omitempty" type:"Struct"`
	// The user\\"s domain name or domain name list.
	//
	// example:
	//
	// 30.yyy.com
	UserDomainName *string `json:"UserDomainName,omitempty" xml:"UserDomainName,omitempty"`
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) String() string {
	return dara.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) GetName() *string {
	return s.Name
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) GetSwitchInfos() *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos {
	return s.SwitchInfos
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) GetUserDomainName() *string {
	return s.UserDomainName
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) SetInstanceId(v string) *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview {
	s.InstanceId = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) SetName(v string) *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview {
	s.Name = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) SetSwitchInfos(v *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos) *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview {
	s.SwitchInfos = v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) SetUserDomainName(v string) *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview {
	s.UserDomainName = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreview) Validate() error {
	if s.SwitchInfos != nil {
		if err := s.SwitchInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos struct {
	SwitchInfo []*PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo `json:"SwitchInfo,omitempty" xml:"SwitchInfo,omitempty" type:"Repeated"`
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos) String() string {
	return dara.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos) GetSwitchInfo() []*PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo {
	return s.SwitchInfo
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos) SetSwitchInfo(v []*PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo) *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos {
	s.SwitchInfo = v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfos) Validate() error {
	if s.SwitchInfo != nil {
		for _, item := range s.SwitchInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo struct {
	// The formatted message content.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The name of the switching policy for address pools.
	//
	// example:
	//
	// strategy-name-example-1
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo) String() string {
	return dara.Prettify(s)
}

func (s PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo) GetContent() *string {
	return s.Content
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo) GetStrategyName() *string {
	return s.StrategyName
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo) SetContent(v string) *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo {
	s.Content = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo) SetStrategyName(v string) *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo {
	s.StrategyName = &v
	return s
}

func (s *PreviewGtmRecoveryPlanResponseBodyPreviewsPreviewSwitchInfosSwitchInfo) Validate() error {
	return dara.Validate(s)
}
