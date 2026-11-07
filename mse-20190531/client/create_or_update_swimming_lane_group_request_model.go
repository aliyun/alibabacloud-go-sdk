// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetAcceptLanguage() *string
	SetAppIds(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetAppIds() *string
	SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneGroupRequest
	GetCanaryModel() *int32
	SetDbGrayEnable(v bool) *CreateOrUpdateSwimmingLaneGroupRequest
	GetDbGrayEnable() *bool
	SetEntryApp(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetEntryApp() *string
	SetId(v int64) *CreateOrUpdateSwimmingLaneGroupRequest
	GetId() *int64
	SetMessageQueueFilterSide(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetMessageQueueFilterSide() *string
	SetMessageQueueGrayEnable(v bool) *CreateOrUpdateSwimmingLaneGroupRequest
	GetMessageQueueGrayEnable() *bool
	SetName(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetName() *string
	SetNamespace(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetNamespace() *string
	SetPaths(v []*string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetPaths() []*string
	SetRecordCanaryDetail(v bool) *CreateOrUpdateSwimmingLaneGroupRequest
	GetRecordCanaryDetail() *bool
	SetRegion(v string) *CreateOrUpdateSwimmingLaneGroupRequest
	GetRegion() *string
	SetRouteIds(v []*int64) *CreateOrUpdateSwimmingLaneGroupRequest
	GetRouteIds() []*int64
	SetStatus(v int32) *CreateOrUpdateSwimmingLaneGroupRequest
	GetStatus() *int32
	SetSwimVersion(v int32) *CreateOrUpdateSwimmingLaneGroupRequest
	GetSwimVersion() *int32
}

type CreateOrUpdateSwimmingLaneGroupRequest struct {
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
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
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
	RouteIds []*int64 `json:"RouteIds,omitempty" xml:"RouteIds,omitempty" type:"Repeated"`
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

func (s CreateOrUpdateSwimmingLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetDbGrayEnable() *bool {
	return s.DbGrayEnable
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetEntryApp() *string {
	return s.EntryApp
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetMessageQueueFilterSide() *string {
	return s.MessageQueueFilterSide
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetMessageQueueGrayEnable() *bool {
	return s.MessageQueueGrayEnable
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetPaths() []*string {
	return s.Paths
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetRecordCanaryDetail() *bool {
	return s.RecordCanaryDetail
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetRouteIds() []*int64 {
	return s.RouteIds
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) GetSwimVersion() *int32 {
	return s.SwimVersion
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetAcceptLanguage(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetAppIds(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.AppIds = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.CanaryModel = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetDbGrayEnable(v bool) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.DbGrayEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetEntryApp(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.EntryApp = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetId(v int64) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetMessageQueueFilterSide(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.MessageQueueFilterSide = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetMessageQueueGrayEnable(v bool) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.MessageQueueGrayEnable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetName(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetNamespace(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.Namespace = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetPaths(v []*string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.Paths = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetRecordCanaryDetail(v bool) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.RecordCanaryDetail = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetRegion(v string) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.Region = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetRouteIds(v []*int64) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.RouteIds = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetStatus(v int32) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.Status = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) SetSwimVersion(v int32) *CreateOrUpdateSwimmingLaneGroupRequest {
	s.SwimVersion = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
