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
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The IDs of the applications. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// xxxxn1po62@492b8f98d4bxxxx,xxxxn1po62@2d8f5d1b7f2xxxx
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// The end-to-end canary release mode. Valid values:
	//
	// - 0: Route by request content.
	//
	// - 1: Route by proportion.
	//
	// example:
	//
	// 1
	CanaryModel *int32 `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// **[Deprecated]*	- Specifies whether to enable canary release for databases.
	//
	// example:
	//
	// true
	DbGrayEnable *bool `json:"DbGrayEnable,omitempty" xml:"DbGrayEnable,omitempty"`
	// The entry application. Format: Source system:ID. Examples:
	//
	// - mse:${MSE application ID}
	//
	// - mse-gw:${MSE cloud-native gateway ID}
	//
	// - apig:${cloud-native API gateway ID}
	//
	// example:
	//
	// mse-gw:gw-2b1407e050f1446xxxcd2a65fd9xxxx
	EntryApp *string `json:"EntryApp,omitempty" xml:"EntryApp,omitempty"`
	// The ID of the swimlane group. A value of -1 indicates that a swimlane group is created. A value greater than 0 indicates that a specified swimlane group is updated.
	//
	// example:
	//
	// 120
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The side where message queue canary release filtering is performed.
	//
	// example:
	//
	// Server
	MessageQueueFilterSide *string `json:"MessageQueueFilterSide,omitempty" xml:"MessageQueueFilterSide,omitempty"`
	// Specifies whether to enable canary release for message queues.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	MessageQueueGrayEnable *bool `json:"MessageQueueGrayEnable,omitempty" xml:"MessageQueueGrayEnable,omitempty"`
	// The name of the swimlane group.
	//
	// This parameter is required.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the MSE namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// **[Deprecated]*	- The paths. This parameter is required when a Java gateway is used as the entry point.
	PathsShrink *string `json:"Paths,omitempty" xml:"Paths,omitempty"`
	// Specifies whether to record request details.
	//
	// example:
	//
	// false
	RecordCanaryDetail *bool `json:"RecordCanaryDetail,omitempty" xml:"RecordCanaryDetail,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The IDs of the routes. If a cloud-native gateway is used as the entry point, you must call this operation to pass the IDs of the selected cloud-native gateway routes before creating a swimlane for the swimlane group.
	RouteIdsShrink *string `json:"RouteIds,omitempty" xml:"RouteIds,omitempty"`
	// The status of the swimlane group. Valid values:
	//
	// - 0: Inactive.
	//
	// - 1: Active.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version of the end-to-end canary release. Valid values:
	//
	// - 0: Not recommended. Routes must be specified.
	//
	// - 2: Recommended. Routes do not need to be specified.
	//
	// example:
	//
	// 2
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
