// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetAcceptLanguage() *string
	SetAppIds(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetAppIds() *string
	SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetCanaryModel() *int32
	SetDbGrayEnable(v bool) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetDbGrayEnable() *bool
	SetEntryApp(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetEntryApp() *string
	SetId(v int64) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetId() *int64
	SetMessageQueueFilterSide(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetMessageQueueFilterSide() *string
	SetMessageQueueGrayEnable(v bool) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetMessageQueueGrayEnable() *bool
	SetName(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetName() *string
	SetNamespace(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetNamespace() *string
	SetPathsShrink(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetPathsShrink() *string
	SetRecordCanaryDetail(v bool) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetRecordCanaryDetail() *bool
	SetRegion(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetRegion() *string
	SetRouteIdsShrink(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetRouteIdsShrink() *string
	SetStatus(v int32) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetStatus() *int32
	SetSwimVersion(v int32) *CreateOrUpdateSwimmingLaneGroupShrinkRequest
	GetSwimVersion() *int32
}

type CreateOrUpdateSwimmingLaneGroupShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The IDs of applications. Separate application IDs with commas (,).
	//
	// example:
	//
	// hkhon1po62@c3df23522baa898,hkhon1po62@66e5235415730a5,hkhon1po62@958bba95910341f
	AppIds      *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	CanaryModel *int32  `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// Specifies whether to enable database canary release.
	//
	// example:
	//
	// true
	DbGrayEnable *bool `json:"DbGrayEnable,omitempty" xml:"DbGrayEnable,omitempty"`
	// The ingress application.
	//
	// example:
	//
	// Ingress
	EntryApp *string `json:"EntryApp,omitempty" xml:"EntryApp,omitempty"`
	// The ID of the lane group. A value of -1 is used to create a lane group. A value greater than 0 is used to modify the specified lane group.
	//
	// example:
	//
	// 120
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The side for message filtering when the canary release for messaging feature is enabled.
	//
	// example:
	//
	// Server
	MessageQueueFilterSide *string `json:"MessageQueueFilterSide,omitempty" xml:"MessageQueueFilterSide,omitempty"`
	// Specifies whether to enable canary release for messaging.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	MessageQueueGrayEnable *bool `json:"MessageQueueGrayEnable,omitempty" xml:"MessageQueueGrayEnable,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the Microservices Engine (MSE) namespace.
	//
	// example:
	//
	// default
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PathsShrink *string `json:"Paths,omitempty" xml:"Paths,omitempty"`
	// Specifies whether to record request details.
	RecordCanaryDetail *bool `json:"RecordCanaryDetail,omitempty" xml:"RecordCanaryDetail,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RouteIdsShrink *string `json:"RouteIds,omitempty" xml:"RouteIds,omitempty"`
	// The status of the lane group. The value 0 specifies that the lane group is disabled. The value 1 specifies that the lane group is enabled.
	//
	// example:
	//
	// 0
	Status      *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	SwimVersion *int32 `json:"SwimVersion,omitempty" xml:"SwimVersion,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetDbGrayEnable() *bool {
	return s.DbGrayEnable
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetEntryApp() *string {
	return s.EntryApp
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetMessageQueueFilterSide() *string {
	return s.MessageQueueFilterSide
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetMessageQueueGrayEnable() *bool {
	return s.MessageQueueGrayEnable
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetPathsShrink() *string {
	return s.PathsShrink
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetRecordCanaryDetail() *bool {
	return s.RecordCanaryDetail
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetRouteIdsShrink() *string {
	return s.RouteIdsShrink
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) GetSwimVersion() *int32 {
	return s.SwimVersion
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetAcceptLanguage(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetAppIds(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.AppIds = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.CanaryModel = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetDbGrayEnable(v bool) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.DbGrayEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetEntryApp(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.EntryApp = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetId(v int64) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetMessageQueueFilterSide(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.MessageQueueFilterSide = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetMessageQueueGrayEnable(v bool) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.MessageQueueGrayEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetName(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetNamespace(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetPathsShrink(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.PathsShrink = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetRecordCanaryDetail(v bool) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.RecordCanaryDetail = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetRegion(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.Region = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetRouteIdsShrink(v string) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.RouteIdsShrink = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetStatus(v int32) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) SetSwimVersion(v int32) *CreateOrUpdateSwimmingLaneGroupShrinkRequest {
	s.SwimVersion = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
