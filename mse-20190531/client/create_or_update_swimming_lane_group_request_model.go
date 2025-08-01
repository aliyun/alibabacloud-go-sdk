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
	Namespace *string   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Paths     []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// Specifies whether to record request details.
	RecordCanaryDetail *bool `json:"RecordCanaryDetail,omitempty" xml:"RecordCanaryDetail,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region   *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	RouteIds []*int64 `json:"RouteIds,omitempty" xml:"RouteIds,omitempty" type:"Repeated"`
	// The status of the lane group. The value 0 specifies that the lane group is disabled. The value 1 specifies that the lane group is enabled.
	//
	// example:
	//
	// 0
	Status      *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
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
