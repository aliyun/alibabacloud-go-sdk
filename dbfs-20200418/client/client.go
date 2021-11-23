// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddTagsBatchRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbfsList    *string `json:"DbfsList,omitempty" xml:"DbfsList,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddTagsBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTagsBatchRequest) GoString() string {
	return s.String()
}

func (s *AddTagsBatchRequest) SetClientToken(v string) *AddTagsBatchRequest {
	s.ClientToken = &v
	return s
}

func (s *AddTagsBatchRequest) SetDbfsList(v string) *AddTagsBatchRequest {
	s.DbfsList = &v
	return s
}

func (s *AddTagsBatchRequest) SetRegionId(v string) *AddTagsBatchRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagsBatchRequest) SetTags(v string) *AddTagsBatchRequest {
	s.Tags = &v
	return s
}

type AddTagsBatchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTagsBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTagsBatchResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsBatchResponseBody) SetRequestId(v string) *AddTagsBatchResponseBody {
	s.RequestId = &v
	return s
}

type AddTagsBatchResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddTagsBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTagsBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTagsBatchResponse) GoString() string {
	return s.String()
}

func (s *AddTagsBatchResponse) SetHeaders(v map[string]*string) *AddTagsBatchResponse {
	s.Headers = v
	return s
}

func (s *AddTagsBatchResponse) SetBody(v *AddTagsBatchResponseBody) *AddTagsBatchResponse {
	s.Body = v
	return s
}

type AttachDbfsRequest struct {
	AttachMode    *string `json:"AttachMode,omitempty" xml:"AttachMode,omitempty"`
	AttachPoint   *string `json:"AttachPoint,omitempty" xml:"AttachPoint,omitempty"`
	ECSInstanceId *string `json:"ECSInstanceId,omitempty" xml:"ECSInstanceId,omitempty"`
	FsId          *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerUrl     *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
}

func (s AttachDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachDbfsRequest) GoString() string {
	return s.String()
}

func (s *AttachDbfsRequest) SetAttachMode(v string) *AttachDbfsRequest {
	s.AttachMode = &v
	return s
}

func (s *AttachDbfsRequest) SetAttachPoint(v string) *AttachDbfsRequest {
	s.AttachPoint = &v
	return s
}

func (s *AttachDbfsRequest) SetECSInstanceId(v string) *AttachDbfsRequest {
	s.ECSInstanceId = &v
	return s
}

func (s *AttachDbfsRequest) SetFsId(v string) *AttachDbfsRequest {
	s.FsId = &v
	return s
}

func (s *AttachDbfsRequest) SetRegionId(v string) *AttachDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *AttachDbfsRequest) SetServerUrl(v string) *AttachDbfsRequest {
	s.ServerUrl = &v
	return s
}

type AttachDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDbfsResponseBody) SetRequestId(v string) *AttachDbfsResponseBody {
	s.RequestId = &v
	return s
}

type AttachDbfsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachDbfsResponse) GoString() string {
	return s.String()
}

func (s *AttachDbfsResponse) SetHeaders(v map[string]*string) *AttachDbfsResponse {
	s.Headers = v
	return s
}

func (s *AttachDbfsResponse) SetBody(v *AttachDbfsResponseBody) *AttachDbfsResponse {
	s.Body = v
	return s
}

type CreateConstantsRequest struct {
	ConstantsData *string `json:"ConstantsData,omitempty" xml:"ConstantsData,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateConstantsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConstantsRequest) GoString() string {
	return s.String()
}

func (s *CreateConstantsRequest) SetConstantsData(v string) *CreateConstantsRequest {
	s.ConstantsData = &v
	return s
}

func (s *CreateConstantsRequest) SetPageNumber(v int32) *CreateConstantsRequest {
	s.PageNumber = &v
	return s
}

func (s *CreateConstantsRequest) SetPageSize(v int32) *CreateConstantsRequest {
	s.PageSize = &v
	return s
}

func (s *CreateConstantsRequest) SetRegionId(v string) *CreateConstantsRequest {
	s.RegionId = &v
	return s
}

type CreateConstantsResponseBody struct {
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateConstantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConstantsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConstantsResponseBody) SetData(v string) *CreateConstantsResponseBody {
	s.Data = &v
	return s
}

func (s *CreateConstantsResponseBody) SetPageNumber(v int64) *CreateConstantsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CreateConstantsResponseBody) SetPageSize(v int64) *CreateConstantsResponseBody {
	s.PageSize = &v
	return s
}

func (s *CreateConstantsResponseBody) SetRequestId(v string) *CreateConstantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConstantsResponseBody) SetTotalCount(v int64) *CreateConstantsResponseBody {
	s.TotalCount = &v
	return s
}

type CreateConstantsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConstantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConstantsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConstantsResponse) GoString() string {
	return s.String()
}

func (s *CreateConstantsResponse) SetHeaders(v map[string]*string) *CreateConstantsResponse {
	s.Headers = v
	return s
}

func (s *CreateConstantsResponse) SetBody(v *CreateConstantsResponseBody) *CreateConstantsResponse {
	s.Body = v
	return s
}

type CreateDbfsRequest struct {
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeleteSnapshot       *bool   `json:"DeleteSnapshot,omitempty" xml:"DeleteSnapshot,omitempty"`
	EnableRaid           *bool   `json:"EnableRaid,omitempty" xml:"EnableRaid,omitempty"`
	Encryption           *bool   `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	FsName               *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KMSKeyId             *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	RaidStripeUnitNumber *int32  `json:"RaidStripeUnitNumber,omitempty" xml:"RaidStripeUnitNumber,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SizeG                *int32  `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
	SnapshotId           *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	UsedScene            *string `json:"UsedScene,omitempty" xml:"UsedScene,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDbfsRequest) GoString() string {
	return s.String()
}

func (s *CreateDbfsRequest) SetCategory(v string) *CreateDbfsRequest {
	s.Category = &v
	return s
}

func (s *CreateDbfsRequest) SetClientToken(v string) *CreateDbfsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDbfsRequest) SetDeleteSnapshot(v bool) *CreateDbfsRequest {
	s.DeleteSnapshot = &v
	return s
}

func (s *CreateDbfsRequest) SetEnableRaid(v bool) *CreateDbfsRequest {
	s.EnableRaid = &v
	return s
}

func (s *CreateDbfsRequest) SetEncryption(v bool) *CreateDbfsRequest {
	s.Encryption = &v
	return s
}

func (s *CreateDbfsRequest) SetFsName(v string) *CreateDbfsRequest {
	s.FsName = &v
	return s
}

func (s *CreateDbfsRequest) SetInstanceType(v string) *CreateDbfsRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateDbfsRequest) SetKMSKeyId(v string) *CreateDbfsRequest {
	s.KMSKeyId = &v
	return s
}

func (s *CreateDbfsRequest) SetPerformanceLevel(v string) *CreateDbfsRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateDbfsRequest) SetRaidStripeUnitNumber(v int32) *CreateDbfsRequest {
	s.RaidStripeUnitNumber = &v
	return s
}

func (s *CreateDbfsRequest) SetRegionId(v string) *CreateDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDbfsRequest) SetSizeG(v int32) *CreateDbfsRequest {
	s.SizeG = &v
	return s
}

func (s *CreateDbfsRequest) SetSnapshotId(v string) *CreateDbfsRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateDbfsRequest) SetUsedScene(v string) *CreateDbfsRequest {
	s.UsedScene = &v
	return s
}

func (s *CreateDbfsRequest) SetZoneId(v string) *CreateDbfsRequest {
	s.ZoneId = &v
	return s
}

type CreateDbfsResponseBody struct {
	FsId      *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDbfsResponseBody) SetFsId(v string) *CreateDbfsResponseBody {
	s.FsId = &v
	return s
}

func (s *CreateDbfsResponseBody) SetRequestId(v string) *CreateDbfsResponseBody {
	s.RequestId = &v
	return s
}

type CreateDbfsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDbfsResponse) GoString() string {
	return s.String()
}

func (s *CreateDbfsResponse) SetHeaders(v map[string]*string) *CreateDbfsResponse {
	s.Headers = v
	return s
}

func (s *CreateDbfsResponse) SetBody(v *CreateDbfsResponseBody) *CreateDbfsResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetClientToken(v string) *CreateServiceLinkedRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CreateServiceLinkedRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FsId          *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RetentionDays *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SnapshotName  *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetClientToken(v string) *CreateSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetFsId(v string) *CreateSnapshotRequest {
	s.FsId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRetentionDays(v int32) *CreateSnapshotRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

type CreateSnapshotResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type DeleteConstantsRequest struct {
	ConstantsData *string `json:"ConstantsData,omitempty" xml:"ConstantsData,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteConstantsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConstantsRequest) GoString() string {
	return s.String()
}

func (s *DeleteConstantsRequest) SetConstantsData(v string) *DeleteConstantsRequest {
	s.ConstantsData = &v
	return s
}

func (s *DeleteConstantsRequest) SetPageNumber(v int32) *DeleteConstantsRequest {
	s.PageNumber = &v
	return s
}

func (s *DeleteConstantsRequest) SetPageSize(v int32) *DeleteConstantsRequest {
	s.PageSize = &v
	return s
}

func (s *DeleteConstantsRequest) SetRegionId(v string) *DeleteConstantsRequest {
	s.RegionId = &v
	return s
}

type DeleteConstantsResponseBody struct {
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DeleteConstantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConstantsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConstantsResponseBody) SetData(v string) *DeleteConstantsResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteConstantsResponseBody) SetPageNumber(v int64) *DeleteConstantsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DeleteConstantsResponseBody) SetPageSize(v int64) *DeleteConstantsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DeleteConstantsResponseBody) SetRequestId(v string) *DeleteConstantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConstantsResponseBody) SetTotalCount(v int64) *DeleteConstantsResponseBody {
	s.TotalCount = &v
	return s
}

type DeleteConstantsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConstantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConstantsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConstantsResponse) GoString() string {
	return s.String()
}

func (s *DeleteConstantsResponse) SetHeaders(v map[string]*string) *DeleteConstantsResponse {
	s.Headers = v
	return s
}

func (s *DeleteConstantsResponse) SetBody(v *DeleteConstantsResponseBody) *DeleteConstantsResponse {
	s.Body = v
	return s
}

type DeleteDbfsRequest struct {
	FsId     *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDbfsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDbfsRequest) SetFsId(v string) *DeleteDbfsRequest {
	s.FsId = &v
	return s
}

func (s *DeleteDbfsRequest) SetRegionId(v string) *DeleteDbfsRequest {
	s.RegionId = &v
	return s
}

type DeleteDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDbfsResponseBody) SetRequestId(v string) *DeleteDbfsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDbfsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDbfsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDbfsResponse) SetHeaders(v map[string]*string) *DeleteDbfsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDbfsResponse) SetBody(v *DeleteDbfsResponseBody) *DeleteDbfsResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRequest struct {
	Force      *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetForce(v bool) *DeleteSnapshotRequest {
	s.Force = &v
	return s
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type DeleteSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type DeleteTagsBatchRequest struct {
	DbfsList *string `json:"DbfsList,omitempty" xml:"DbfsList,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DeleteTagsBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsBatchRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagsBatchRequest) SetDbfsList(v string) *DeleteTagsBatchRequest {
	s.DbfsList = &v
	return s
}

func (s *DeleteTagsBatchRequest) SetRegionId(v string) *DeleteTagsBatchRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTagsBatchRequest) SetTags(v string) *DeleteTagsBatchRequest {
	s.Tags = &v
	return s
}

type DeleteTagsBatchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagsBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagsBatchResponseBody) SetRequestId(v string) *DeleteTagsBatchResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTagsBatchResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTagsBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTagsBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsBatchResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagsBatchResponse) SetHeaders(v map[string]*string) *DeleteTagsBatchResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagsBatchResponse) SetBody(v *DeleteTagsBatchResponseBody) *DeleteTagsBatchResponse {
	s.Body = v
	return s
}

type DescribeDbfsSpecificationsRequest struct {
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDbfsSpecificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbfsSpecificationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbfsSpecificationsRequest) SetCategory(v string) *DescribeDbfsSpecificationsRequest {
	s.Category = &v
	return s
}

func (s *DescribeDbfsSpecificationsRequest) SetEcsInstanceType(v string) *DescribeDbfsSpecificationsRequest {
	s.EcsInstanceType = &v
	return s
}

func (s *DescribeDbfsSpecificationsRequest) SetRegionId(v string) *DescribeDbfsSpecificationsRequest {
	s.RegionId = &v
	return s
}

type DescribeDbfsSpecificationsResponseBody struct {
	MaxDbfsNumberPerEcs            map[string]interface{} `json:"MaxDbfsNumberPerEcs,omitempty" xml:"MaxDbfsNumberPerEcs,omitempty"`
	RequestId                      *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportedEcsInstanceTypeFamily []*string              `json:"SupportedEcsInstanceTypeFamily,omitempty" xml:"SupportedEcsInstanceTypeFamily,omitempty" type:"Repeated"`
}

func (s DescribeDbfsSpecificationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbfsSpecificationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbfsSpecificationsResponseBody) SetMaxDbfsNumberPerEcs(v map[string]interface{}) *DescribeDbfsSpecificationsResponseBody {
	s.MaxDbfsNumberPerEcs = v
	return s
}

func (s *DescribeDbfsSpecificationsResponseBody) SetRequestId(v string) *DescribeDbfsSpecificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbfsSpecificationsResponseBody) SetSupportedEcsInstanceTypeFamily(v []*string) *DescribeDbfsSpecificationsResponseBody {
	s.SupportedEcsInstanceTypeFamily = v
	return s
}

type DescribeDbfsSpecificationsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDbfsSpecificationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDbfsSpecificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbfsSpecificationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbfsSpecificationsResponse) SetHeaders(v map[string]*string) *DescribeDbfsSpecificationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbfsSpecificationsResponse) SetBody(v *DescribeDbfsSpecificationsResponseBody) *DescribeDbfsSpecificationsResponse {
	s.Body = v
	return s
}

type DescribeInstanceTypesRequest struct {
	// 地域ID
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesRequest) SetRegionId(v string) *DescribeInstanceTypesRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceTypesResponseBody struct {
	// Array of InstanceType
	InstanceTypes []*DescribeInstanceTypesResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBody) SetInstanceTypes(v []*DescribeInstanceTypesResponseBodyInstanceTypes) *DescribeInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *DescribeInstanceTypesResponseBody) SetRequestId(v string) *DescribeInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceTypesResponseBodyInstanceTypes struct {
	CpuCoreCount            *float32 `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	InstanceTypeDescription *string  `json:"InstanceTypeDescription,omitempty" xml:"InstanceTypeDescription,omitempty"`
	InstanceTypeId          *string  `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	MemorySize              *float32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetCpuCoreCount(v float32) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.CpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetInstanceTypeDescription(v string) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.InstanceTypeDescription = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetInstanceTypeId(v string) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.InstanceTypeId = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetMemorySize(v float32) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.MemorySize = &v
	return s
}

type DescribeInstanceTypesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponse) SetHeaders(v map[string]*string) *DescribeInstanceTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTypesResponse) SetBody(v *DescribeInstanceTypesResponseBody) *DescribeInstanceTypesResponse {
	s.Body = v
	return s
}

type DetachDbfsRequest struct {
	ECSInstanceId *string `json:"ECSInstanceId,omitempty" xml:"ECSInstanceId,omitempty"`
	FsId          *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachDbfsRequest) GoString() string {
	return s.String()
}

func (s *DetachDbfsRequest) SetECSInstanceId(v string) *DetachDbfsRequest {
	s.ECSInstanceId = &v
	return s
}

func (s *DetachDbfsRequest) SetFsId(v string) *DetachDbfsRequest {
	s.FsId = &v
	return s
}

func (s *DetachDbfsRequest) SetRegionId(v string) *DetachDbfsRequest {
	s.RegionId = &v
	return s
}

type DetachDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDbfsResponseBody) SetRequestId(v string) *DetachDbfsResponseBody {
	s.RequestId = &v
	return s
}

type DetachDbfsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachDbfsResponse) GoString() string {
	return s.String()
}

func (s *DetachDbfsResponse) SetHeaders(v map[string]*string) *DetachDbfsResponse {
	s.Headers = v
	return s
}

func (s *DetachDbfsResponse) SetBody(v *DetachDbfsResponseBody) *DetachDbfsResponse {
	s.Body = v
	return s
}

type GetDbfsRequest struct {
	FsId     *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsRequest) GoString() string {
	return s.String()
}

func (s *GetDbfsRequest) SetFsId(v string) *GetDbfsRequest {
	s.FsId = &v
	return s
}

func (s *GetDbfsRequest) SetRegionId(v string) *GetDbfsRequest {
	s.RegionId = &v
	return s
}

type GetDbfsResponseBody struct {
	DBFSInfo  *GetDbfsResponseBodyDBFSInfo `json:"DBFSInfo,omitempty" xml:"DBFSInfo,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBody) SetDBFSInfo(v *GetDbfsResponseBodyDBFSInfo) *GetDbfsResponseBody {
	s.DBFSInfo = v
	return s
}

func (s *GetDbfsResponseBody) SetRequestId(v string) *GetDbfsResponseBody {
	s.RequestId = &v
	return s
}

type GetDbfsResponseBodyDBFSInfo struct {
	AttachNodeNumber *int32                                `json:"AttachNodeNumber,omitempty" xml:"AttachNodeNumber,omitempty"`
	Category         *string                               `json:"Category,omitempty" xml:"Category,omitempty"`
	CreatedTime      *string                               `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DBFSClusterId    *string                               `json:"DBFSClusterId,omitempty" xml:"DBFSClusterId,omitempty"`
	Description      *string                               `json:"Description,omitempty" xml:"Description,omitempty"`
	EbsList          []*GetDbfsResponseBodyDBFSInfoEbsList `json:"EbsList,omitempty" xml:"EbsList,omitempty" type:"Repeated"`
	EcsList          []*GetDbfsResponseBodyDBFSInfoEcsList `json:"EcsList,omitempty" xml:"EcsList,omitempty" type:"Repeated"`
	EnableRaid       *bool                                 `json:"EnableRaid,omitempty" xml:"EnableRaid,omitempty"`
	Encryption       *bool                                 `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	FsId             *string                               `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName           *string                               `json:"FsName,omitempty" xml:"FsName,omitempty"`
	InstanceType     *string                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KMSKeyId         *string                               `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	LastMountTime    *string                               `json:"LastMountTime,omitempty" xml:"LastMountTime,omitempty"`
	LastUmountTime   *string                               `json:"LastUmountTime,omitempty" xml:"LastUmountTime,omitempty"`
	PayType          *string                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PerformanceLevel *string                               `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	RaidStrip        *int32                                `json:"RaidStrip,omitempty" xml:"RaidStrip,omitempty"`
	RegionId         *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SizeG            *int32                                `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
	Status           *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags             []*GetDbfsResponseBodyDBFSInfoTags    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UsedScene        *string                               `json:"UsedScene,omitempty" xml:"UsedScene,omitempty"`
	ZoneId           *string                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetDbfsResponseBodyDBFSInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBodyDBFSInfo) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBodyDBFSInfo) SetAttachNodeNumber(v int32) *GetDbfsResponseBodyDBFSInfo {
	s.AttachNodeNumber = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetCategory(v string) *GetDbfsResponseBodyDBFSInfo {
	s.Category = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetCreatedTime(v string) *GetDbfsResponseBodyDBFSInfo {
	s.CreatedTime = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetDBFSClusterId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.DBFSClusterId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetDescription(v string) *GetDbfsResponseBodyDBFSInfo {
	s.Description = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetEbsList(v []*GetDbfsResponseBodyDBFSInfoEbsList) *GetDbfsResponseBodyDBFSInfo {
	s.EbsList = v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetEcsList(v []*GetDbfsResponseBodyDBFSInfoEcsList) *GetDbfsResponseBodyDBFSInfo {
	s.EcsList = v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetEnableRaid(v bool) *GetDbfsResponseBodyDBFSInfo {
	s.EnableRaid = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetEncryption(v bool) *GetDbfsResponseBodyDBFSInfo {
	s.Encryption = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetFsId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.FsId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetFsName(v string) *GetDbfsResponseBodyDBFSInfo {
	s.FsName = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetInstanceType(v string) *GetDbfsResponseBodyDBFSInfo {
	s.InstanceType = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetKMSKeyId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.KMSKeyId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetLastMountTime(v string) *GetDbfsResponseBodyDBFSInfo {
	s.LastMountTime = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetLastUmountTime(v string) *GetDbfsResponseBodyDBFSInfo {
	s.LastUmountTime = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetPayType(v string) *GetDbfsResponseBodyDBFSInfo {
	s.PayType = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetPerformanceLevel(v string) *GetDbfsResponseBodyDBFSInfo {
	s.PerformanceLevel = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetRaidStrip(v int32) *GetDbfsResponseBodyDBFSInfo {
	s.RaidStrip = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetRegionId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.RegionId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetSizeG(v int32) *GetDbfsResponseBodyDBFSInfo {
	s.SizeG = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetStatus(v string) *GetDbfsResponseBodyDBFSInfo {
	s.Status = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetTags(v []*GetDbfsResponseBodyDBFSInfoTags) *GetDbfsResponseBodyDBFSInfo {
	s.Tags = v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetUsedScene(v string) *GetDbfsResponseBodyDBFSInfo {
	s.UsedScene = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetZoneId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.ZoneId = &v
	return s
}

type GetDbfsResponseBodyDBFSInfoEbsList struct {
	EbsId *string `json:"EbsId,omitempty" xml:"EbsId,omitempty"`
	SizeG *int32  `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
}

func (s GetDbfsResponseBodyDBFSInfoEbsList) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBodyDBFSInfoEbsList) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBodyDBFSInfoEbsList) SetEbsId(v string) *GetDbfsResponseBodyDBFSInfoEbsList {
	s.EbsId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfoEbsList) SetSizeG(v int32) *GetDbfsResponseBodyDBFSInfoEbsList {
	s.SizeG = &v
	return s
}

type GetDbfsResponseBodyDBFSInfoEcsList struct {
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
}

func (s GetDbfsResponseBodyDBFSInfoEcsList) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBodyDBFSInfoEcsList) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBodyDBFSInfoEcsList) SetEcsId(v string) *GetDbfsResponseBodyDBFSInfoEcsList {
	s.EcsId = &v
	return s
}

type GetDbfsResponseBodyDBFSInfoTags struct {
	Id       *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetDbfsResponseBodyDBFSInfoTags) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBodyDBFSInfoTags) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBodyDBFSInfoTags) SetId(v int32) *GetDbfsResponseBodyDBFSInfoTags {
	s.Id = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfoTags) SetTagKey(v string) *GetDbfsResponseBodyDBFSInfoTags {
	s.TagKey = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfoTags) SetTagValue(v string) *GetDbfsResponseBodyDBFSInfoTags {
	s.TagValue = &v
	return s
}

type GetDbfsResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponse) GoString() string {
	return s.String()
}

func (s *GetDbfsResponse) SetHeaders(v map[string]*string) *GetDbfsResponse {
	s.Headers = v
	return s
}

func (s *GetDbfsResponse) SetBody(v *GetDbfsResponseBody) *GetDbfsResponse {
	s.Body = v
	return s
}

type GetServiceLinkedRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleRequest) SetRegionId(v string) *GetServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type GetServiceLinkedRoleResponseBody struct {
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DbfsLinkedRole *bool   `json:"DbfsLinkedRole,omitempty" xml:"DbfsLinkedRole,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleResponseBody) SetAccountId(v string) *GetServiceLinkedRoleResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetServiceLinkedRoleResponseBody) SetDbfsLinkedRole(v bool) *GetServiceLinkedRoleResponseBody {
	s.DbfsLinkedRole = &v
	return s
}

func (s *GetServiceLinkedRoleResponseBody) SetRegionId(v string) *GetServiceLinkedRoleResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetServiceLinkedRoleResponseBody) SetRequestId(v string) *GetServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceLinkedRoleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *GetServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLinkedRoleResponse) SetBody(v *GetServiceLinkedRoleResponseBody) *GetServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type ListConstantsRequest struct {
	ConstantsData *string `json:"ConstantsData,omitempty" xml:"ConstantsData,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListConstantsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConstantsRequest) GoString() string {
	return s.String()
}

func (s *ListConstantsRequest) SetConstantsData(v string) *ListConstantsRequest {
	s.ConstantsData = &v
	return s
}

func (s *ListConstantsRequest) SetPageNumber(v int32) *ListConstantsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConstantsRequest) SetPageSize(v int32) *ListConstantsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConstantsRequest) SetRegionId(v string) *ListConstantsRequest {
	s.RegionId = &v
	return s
}

type ListConstantsResponseBody struct {
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConstantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConstantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConstantsResponseBody) SetData(v string) *ListConstantsResponseBody {
	s.Data = &v
	return s
}

func (s *ListConstantsResponseBody) SetPageNumber(v int64) *ListConstantsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListConstantsResponseBody) SetPageSize(v int64) *ListConstantsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListConstantsResponseBody) SetRequestId(v string) *ListConstantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConstantsResponseBody) SetTotalCount(v int64) *ListConstantsResponseBody {
	s.TotalCount = &v
	return s
}

type ListConstantsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConstantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConstantsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConstantsResponse) GoString() string {
	return s.String()
}

func (s *ListConstantsResponse) SetHeaders(v map[string]*string) *ListConstantsResponse {
	s.Headers = v
	return s
}

func (s *ListConstantsResponse) SetBody(v *ListConstantsResponseBody) *ListConstantsResponse {
	s.Body = v
	return s
}

type ListDbfsRequest struct {
	FilterKey   *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortKey     *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	SortType    *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsRequest) GoString() string {
	return s.String()
}

func (s *ListDbfsRequest) SetFilterKey(v string) *ListDbfsRequest {
	s.FilterKey = &v
	return s
}

func (s *ListDbfsRequest) SetFilterValue(v string) *ListDbfsRequest {
	s.FilterValue = &v
	return s
}

func (s *ListDbfsRequest) SetPageNumber(v int32) *ListDbfsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDbfsRequest) SetPageSize(v int32) *ListDbfsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDbfsRequest) SetRegionId(v string) *ListDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDbfsRequest) SetSortKey(v string) *ListDbfsRequest {
	s.SortKey = &v
	return s
}

func (s *ListDbfsRequest) SetSortType(v string) *ListDbfsRequest {
	s.SortType = &v
	return s
}

func (s *ListDbfsRequest) SetTags(v string) *ListDbfsRequest {
	s.Tags = &v
	return s
}

type ListDbfsResponseBody struct {
	DBFSInfo   []*ListDbfsResponseBodyDBFSInfo `json:"DBFSInfo,omitempty" xml:"DBFSInfo,omitempty" type:"Repeated"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBody) SetDBFSInfo(v []*ListDbfsResponseBodyDBFSInfo) *ListDbfsResponseBody {
	s.DBFSInfo = v
	return s
}

func (s *ListDbfsResponseBody) SetPageNumber(v int32) *ListDbfsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDbfsResponseBody) SetPageSize(v int32) *ListDbfsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDbfsResponseBody) SetRequestId(v string) *ListDbfsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDbfsResponseBody) SetTotalCount(v int32) *ListDbfsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDbfsResponseBodyDBFSInfo struct {
	AttachNodeNumber *int32                                 `json:"AttachNodeNumber,omitempty" xml:"AttachNodeNumber,omitempty"`
	Category         *string                                `json:"Category,omitempty" xml:"Category,omitempty"`
	CreatedTime      *string                                `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DBFSClusterId    *string                                `json:"DBFSClusterId,omitempty" xml:"DBFSClusterId,omitempty"`
	EbsList          []*ListDbfsResponseBodyDBFSInfoEbsList `json:"EbsList,omitempty" xml:"EbsList,omitempty" type:"Repeated"`
	EcsList          []*ListDbfsResponseBodyDBFSInfoEcsList `json:"EcsList,omitempty" xml:"EcsList,omitempty" type:"Repeated"`
	EnableRaid       *bool                                  `json:"EnableRaid,omitempty" xml:"EnableRaid,omitempty"`
	Encryption       *bool                                  `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	FsId             *string                                `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName           *string                                `json:"FsName,omitempty" xml:"FsName,omitempty"`
	InstanceType     *string                                `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KMSKeyId         *string                                `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	LastMountTime    *string                                `json:"LastMountTime,omitempty" xml:"LastMountTime,omitempty"`
	LastUmountTime   *string                                `json:"LastUmountTime,omitempty" xml:"LastUmountTime,omitempty"`
	PayType          *string                                `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PerformanceLevel *string                                `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	RaidStrip        *int32                                 `json:"RaidStrip,omitempty" xml:"RaidStrip,omitempty"`
	RegionId         *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SizeG            *int32                                 `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
	Status           *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags             []*ListDbfsResponseBodyDBFSInfoTags    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UsedScene        *string                                `json:"UsedScene,omitempty" xml:"UsedScene,omitempty"`
	ZoneId           *string                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListDbfsResponseBodyDBFSInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBodyDBFSInfo) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBodyDBFSInfo) SetAttachNodeNumber(v int32) *ListDbfsResponseBodyDBFSInfo {
	s.AttachNodeNumber = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetCategory(v string) *ListDbfsResponseBodyDBFSInfo {
	s.Category = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetCreatedTime(v string) *ListDbfsResponseBodyDBFSInfo {
	s.CreatedTime = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetDBFSClusterId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.DBFSClusterId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetEbsList(v []*ListDbfsResponseBodyDBFSInfoEbsList) *ListDbfsResponseBodyDBFSInfo {
	s.EbsList = v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetEcsList(v []*ListDbfsResponseBodyDBFSInfoEcsList) *ListDbfsResponseBodyDBFSInfo {
	s.EcsList = v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetEnableRaid(v bool) *ListDbfsResponseBodyDBFSInfo {
	s.EnableRaid = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetEncryption(v bool) *ListDbfsResponseBodyDBFSInfo {
	s.Encryption = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetFsId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.FsId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetFsName(v string) *ListDbfsResponseBodyDBFSInfo {
	s.FsName = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetInstanceType(v string) *ListDbfsResponseBodyDBFSInfo {
	s.InstanceType = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetKMSKeyId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.KMSKeyId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetLastMountTime(v string) *ListDbfsResponseBodyDBFSInfo {
	s.LastMountTime = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetLastUmountTime(v string) *ListDbfsResponseBodyDBFSInfo {
	s.LastUmountTime = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetPayType(v string) *ListDbfsResponseBodyDBFSInfo {
	s.PayType = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetPerformanceLevel(v string) *ListDbfsResponseBodyDBFSInfo {
	s.PerformanceLevel = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetRaidStrip(v int32) *ListDbfsResponseBodyDBFSInfo {
	s.RaidStrip = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetRegionId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.RegionId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetSizeG(v int32) *ListDbfsResponseBodyDBFSInfo {
	s.SizeG = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetStatus(v string) *ListDbfsResponseBodyDBFSInfo {
	s.Status = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetTags(v []*ListDbfsResponseBodyDBFSInfoTags) *ListDbfsResponseBodyDBFSInfo {
	s.Tags = v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetUsedScene(v string) *ListDbfsResponseBodyDBFSInfo {
	s.UsedScene = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetZoneId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.ZoneId = &v
	return s
}

type ListDbfsResponseBodyDBFSInfoEbsList struct {
	EbsId *string `json:"EbsId,omitempty" xml:"EbsId,omitempty"`
	SizeG *int32  `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
}

func (s ListDbfsResponseBodyDBFSInfoEbsList) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBodyDBFSInfoEbsList) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBodyDBFSInfoEbsList) SetEbsId(v string) *ListDbfsResponseBodyDBFSInfoEbsList {
	s.EbsId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfoEbsList) SetSizeG(v int32) *ListDbfsResponseBodyDBFSInfoEbsList {
	s.SizeG = &v
	return s
}

type ListDbfsResponseBodyDBFSInfoEcsList struct {
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
}

func (s ListDbfsResponseBodyDBFSInfoEcsList) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBodyDBFSInfoEcsList) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBodyDBFSInfoEcsList) SetEcsId(v string) *ListDbfsResponseBodyDBFSInfoEcsList {
	s.EcsId = &v
	return s
}

type ListDbfsResponseBodyDBFSInfoTags struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListDbfsResponseBodyDBFSInfoTags) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBodyDBFSInfoTags) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBodyDBFSInfoTags) SetId(v int64) *ListDbfsResponseBodyDBFSInfoTags {
	s.Id = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfoTags) SetTagKey(v string) *ListDbfsResponseBodyDBFSInfoTags {
	s.TagKey = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfoTags) SetTagValue(v string) *ListDbfsResponseBodyDBFSInfoTags {
	s.TagValue = &v
	return s
}

type ListDbfsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponse) GoString() string {
	return s.String()
}

func (s *ListDbfsResponse) SetHeaders(v map[string]*string) *ListDbfsResponse {
	s.Headers = v
	return s
}

func (s *ListDbfsResponse) SetBody(v *ListDbfsResponseBody) *ListDbfsResponse {
	s.Body = v
	return s
}

type ListSnapshotRequest struct {
	FilterKey    *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue  *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	FsId         *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotIds  *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	SortKey      *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	SortType     *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotRequest) SetFilterKey(v string) *ListSnapshotRequest {
	s.FilterKey = &v
	return s
}

func (s *ListSnapshotRequest) SetFilterValue(v string) *ListSnapshotRequest {
	s.FilterValue = &v
	return s
}

func (s *ListSnapshotRequest) SetFsId(v string) *ListSnapshotRequest {
	s.FsId = &v
	return s
}

func (s *ListSnapshotRequest) SetPageNumber(v int32) *ListSnapshotRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotRequest) SetPageSize(v int32) *ListSnapshotRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotRequest) SetRegionId(v string) *ListSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ListSnapshotRequest) SetSnapshotIds(v string) *ListSnapshotRequest {
	s.SnapshotIds = &v
	return s
}

func (s *ListSnapshotRequest) SetSnapshotName(v string) *ListSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *ListSnapshotRequest) SetSnapshotType(v string) *ListSnapshotRequest {
	s.SnapshotType = &v
	return s
}

func (s *ListSnapshotRequest) SetSortKey(v string) *ListSnapshotRequest {
	s.SortKey = &v
	return s
}

func (s *ListSnapshotRequest) SetSortType(v string) *ListSnapshotRequest {
	s.SortType = &v
	return s
}

func (s *ListSnapshotRequest) SetStatus(v string) *ListSnapshotRequest {
	s.Status = &v
	return s
}

type ListSnapshotResponseBody struct {
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots  []*ListSnapshotResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotResponseBody) SetPageNumber(v int32) *ListSnapshotResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotResponseBody) SetPageSize(v int32) *ListSnapshotResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotResponseBody) SetRequestId(v string) *ListSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotResponseBody) SetSnapshots(v []*ListSnapshotResponseBodySnapshots) *ListSnapshotResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotResponseBody) SetTotalCount(v int32) *ListSnapshotResponseBody {
	s.TotalCount = &v
	return s
}

type ListSnapshotResponseBodySnapshots struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Progress         *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RemainTime       *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	RetentionDays    *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SnapshotId       *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName     *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotType     *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	SourceFsId       *string `json:"SourceFsId,omitempty" xml:"SourceFsId,omitempty"`
	SourceFsSize     *int32  `json:"SourceFsSize,omitempty" xml:"SourceFsSize,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSnapshotResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *ListSnapshotResponseBodySnapshots) SetCategory(v string) *ListSnapshotResponseBodySnapshots {
	s.Category = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetCreationTime(v string) *ListSnapshotResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetDescription(v string) *ListSnapshotResponseBodySnapshots {
	s.Description = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetLastModifiedTime(v string) *ListSnapshotResponseBodySnapshots {
	s.LastModifiedTime = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetProgress(v string) *ListSnapshotResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetRemainTime(v int32) *ListSnapshotResponseBodySnapshots {
	s.RemainTime = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetRetentionDays(v int32) *ListSnapshotResponseBodySnapshots {
	s.RetentionDays = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSnapshotId(v string) *ListSnapshotResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSnapshotName(v string) *ListSnapshotResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSnapshotType(v string) *ListSnapshotResponseBodySnapshots {
	s.SnapshotType = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSourceFsId(v string) *ListSnapshotResponseBodySnapshots {
	s.SourceFsId = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSourceFsSize(v int32) *ListSnapshotResponseBodySnapshots {
	s.SourceFsSize = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetStatus(v string) *ListSnapshotResponseBodySnapshots {
	s.Status = &v
	return s
}

type ListSnapshotResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotResponse) SetHeaders(v map[string]*string) *ListSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotResponse) SetBody(v *ListSnapshotResponseBody) *ListSnapshotResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

type ListTagKeysResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys   []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*string) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

type ListTagValuesResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTagValues(v []*string) *ListTagValuesResponseBody {
	s.TagValues = v
	return s
}

type ListTagValuesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListTaskRequest struct {
	FilterKey   *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortKey     *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	SortType    *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskRequest) GoString() string {
	return s.String()
}

func (s *ListTaskRequest) SetFilterKey(v string) *ListTaskRequest {
	s.FilterKey = &v
	return s
}

func (s *ListTaskRequest) SetFilterValue(v string) *ListTaskRequest {
	s.FilterValue = &v
	return s
}

func (s *ListTaskRequest) SetPageNumber(v int32) *ListTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTaskRequest) SetPageSize(v int32) *ListTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskRequest) SetRegionId(v string) *ListTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ListTaskRequest) SetSortKey(v string) *ListTaskRequest {
	s.SortKey = &v
	return s
}

func (s *ListTaskRequest) SetSortType(v string) *ListTaskRequest {
	s.SortType = &v
	return s
}

type ListTaskResponseBody struct {
	PageNumber *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      []*ListTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	TotalCount *int32                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskResponseBody) SetPageNumber(v int32) *ListTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTaskResponseBody) SetPageSize(v int32) *ListTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTaskResponseBody) SetRequestId(v string) *ListTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskResponseBody) SetTasks(v []*ListTaskResponseBodyTasks) *ListTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTaskResponseBody) SetTotalCount(v int32) *ListTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListTaskResponseBodyTasks struct {
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompletionTime          *string `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	CreatedTime             *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Id                      *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxRetry                *int32  `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	NextExecutionTime       *string `json:"NextExecutionTime,omitempty" xml:"NextExecutionTime,omitempty"`
	Priority                *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	TaskAdder               *string `json:"TaskAdder,omitempty" xml:"TaskAdder,omitempty"`
	TaskErrorReason         *string `json:"TaskErrorReason,omitempty" xml:"TaskErrorReason,omitempty"`
	TaskExecutionCounts     *int32  `json:"TaskExecutionCounts,omitempty" xml:"TaskExecutionCounts,omitempty"`
	TaskName                *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskOwner               *string `json:"TaskOwner,omitempty" xml:"TaskOwner,omitempty"`
	TaskProgress            *int32  `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
	TaskProgressDescription *string `json:"TaskProgressDescription,omitempty" xml:"TaskProgressDescription,omitempty"`
	TaskRunner              *string `json:"TaskRunner,omitempty" xml:"TaskRunner,omitempty"`
	TaskStatus              *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode          *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType                *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListTaskResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTaskResponseBodyTasks) SetClientToken(v string) *ListTaskResponseBodyTasks {
	s.ClientToken = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetCompletionTime(v string) *ListTaskResponseBodyTasks {
	s.CompletionTime = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetCreatedTime(v string) *ListTaskResponseBodyTasks {
	s.CreatedTime = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetId(v int32) *ListTaskResponseBodyTasks {
	s.Id = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetMaxRetry(v int32) *ListTaskResponseBodyTasks {
	s.MaxRetry = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetNextExecutionTime(v string) *ListTaskResponseBodyTasks {
	s.NextExecutionTime = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetPriority(v string) *ListTaskResponseBodyTasks {
	s.Priority = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskAdder(v string) *ListTaskResponseBodyTasks {
	s.TaskAdder = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskErrorReason(v string) *ListTaskResponseBodyTasks {
	s.TaskErrorReason = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskExecutionCounts(v int32) *ListTaskResponseBodyTasks {
	s.TaskExecutionCounts = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskName(v string) *ListTaskResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskOwner(v string) *ListTaskResponseBodyTasks {
	s.TaskOwner = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskProgress(v int32) *ListTaskResponseBodyTasks {
	s.TaskProgress = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskProgressDescription(v string) *ListTaskResponseBodyTasks {
	s.TaskProgressDescription = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskRunner(v string) *ListTaskResponseBodyTasks {
	s.TaskRunner = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskStatus(v string) *ListTaskResponseBodyTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskStatusCode(v int32) *ListTaskResponseBodyTasks {
	s.TaskStatusCode = &v
	return s
}

func (s *ListTaskResponseBodyTasks) SetTaskType(v string) *ListTaskResponseBodyTasks {
	s.TaskType = &v
	return s
}

type ListTaskResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskResponse) GoString() string {
	return s.String()
}

func (s *ListTaskResponse) SetHeaders(v map[string]*string) *ListTaskResponse {
	s.Headers = v
	return s
}

func (s *ListTaskResponse) SetBody(v *ListTaskResponseBody) *ListTaskResponse {
	s.Body = v
	return s
}

type OpreateConstantsRequest struct {
	ConstantsData *string `json:"ConstantsData,omitempty" xml:"ConstantsData,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpreateConstantsRequest) String() string {
	return tea.Prettify(s)
}

func (s OpreateConstantsRequest) GoString() string {
	return s.String()
}

func (s *OpreateConstantsRequest) SetConstantsData(v string) *OpreateConstantsRequest {
	s.ConstantsData = &v
	return s
}

func (s *OpreateConstantsRequest) SetPageNumber(v int32) *OpreateConstantsRequest {
	s.PageNumber = &v
	return s
}

func (s *OpreateConstantsRequest) SetPageSize(v int32) *OpreateConstantsRequest {
	s.PageSize = &v
	return s
}

func (s *OpreateConstantsRequest) SetRegionId(v string) *OpreateConstantsRequest {
	s.RegionId = &v
	return s
}

type OpreateConstantsResponseBody struct {
	AccessData      *string `json:"AccessData,omitempty" xml:"AccessData,omitempty"`
	Data            *string `json:"Data,omitempty" xml:"Data,omitempty"`
	EndpointData    *string `json:"EndpointData,omitempty" xml:"EndpointData,omitempty"`
	MasterData      *string `json:"MasterData,omitempty" xml:"MasterData,omitempty"`
	OsversionData   *string `json:"OsversionData,omitempty" xml:"OsversionData,omitempty"`
	PageNumber      *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductCodeData *string `json:"ProductCodeData,omitempty" xml:"ProductCodeData,omitempty"`
	RegionData      *string `json:"RegionData,omitempty" xml:"RegionData,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ZoneData   *string `json:"ZoneData,omitempty" xml:"ZoneData,omitempty"`
}

func (s OpreateConstantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpreateConstantsResponseBody) GoString() string {
	return s.String()
}

func (s *OpreateConstantsResponseBody) SetAccessData(v string) *OpreateConstantsResponseBody {
	s.AccessData = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetData(v string) *OpreateConstantsResponseBody {
	s.Data = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetEndpointData(v string) *OpreateConstantsResponseBody {
	s.EndpointData = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetMasterData(v string) *OpreateConstantsResponseBody {
	s.MasterData = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetOsversionData(v string) *OpreateConstantsResponseBody {
	s.OsversionData = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetPageNumber(v int64) *OpreateConstantsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetPageSize(v int64) *OpreateConstantsResponseBody {
	s.PageSize = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetProductCodeData(v string) *OpreateConstantsResponseBody {
	s.ProductCodeData = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetRegionData(v string) *OpreateConstantsResponseBody {
	s.RegionData = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetRequestId(v string) *OpreateConstantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetTotalCount(v int64) *OpreateConstantsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *OpreateConstantsResponseBody) SetZoneData(v string) *OpreateConstantsResponseBody {
	s.ZoneData = &v
	return s
}

type OpreateConstantsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpreateConstantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpreateConstantsResponse) String() string {
	return tea.Prettify(s)
}

func (s OpreateConstantsResponse) GoString() string {
	return s.String()
}

func (s *OpreateConstantsResponse) SetHeaders(v map[string]*string) *OpreateConstantsResponse {
	s.Headers = v
	return s
}

func (s *OpreateConstantsResponse) SetBody(v *OpreateConstantsResponseBody) *OpreateConstantsResponse {
	s.Body = v
	return s
}

type RenameDbfsRequest struct {
	FsId     *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName   *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RenameDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameDbfsRequest) GoString() string {
	return s.String()
}

func (s *RenameDbfsRequest) SetFsId(v string) *RenameDbfsRequest {
	s.FsId = &v
	return s
}

func (s *RenameDbfsRequest) SetFsName(v string) *RenameDbfsRequest {
	s.FsName = &v
	return s
}

func (s *RenameDbfsRequest) SetRegionId(v string) *RenameDbfsRequest {
	s.RegionId = &v
	return s
}

type RenameDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *RenameDbfsResponseBody) SetRequestId(v string) *RenameDbfsResponseBody {
	s.RequestId = &v
	return s
}

type RenameDbfsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenameDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameDbfsResponse) GoString() string {
	return s.String()
}

func (s *RenameDbfsResponse) SetHeaders(v map[string]*string) *RenameDbfsResponse {
	s.Headers = v
	return s
}

func (s *RenameDbfsResponse) SetBody(v *RenameDbfsResponseBody) *RenameDbfsResponse {
	s.Body = v
	return s
}

type ResetDbfsRequest struct {
	FsId       *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetDbfsRequest) GoString() string {
	return s.String()
}

func (s *ResetDbfsRequest) SetFsId(v string) *ResetDbfsRequest {
	s.FsId = &v
	return s
}

func (s *ResetDbfsRequest) SetRegionId(v string) *ResetDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDbfsRequest) SetSnapshotId(v string) *ResetDbfsRequest {
	s.SnapshotId = &v
	return s
}

type ResetDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDbfsResponseBody) SetRequestId(v string) *ResetDbfsResponseBody {
	s.RequestId = &v
	return s
}

type ResetDbfsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetDbfsResponse) GoString() string {
	return s.String()
}

func (s *ResetDbfsResponse) SetHeaders(v map[string]*string) *ResetDbfsResponse {
	s.Headers = v
	return s
}

func (s *ResetDbfsResponse) SetBody(v *ResetDbfsResponseBody) *ResetDbfsResponse {
	s.Body = v
	return s
}

type ResizeDbfsRequest struct {
	FsId     *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	NewSizeG *int32  `json:"NewSizeG,omitempty" xml:"NewSizeG,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResizeDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeDbfsRequest) GoString() string {
	return s.String()
}

func (s *ResizeDbfsRequest) SetFsId(v string) *ResizeDbfsRequest {
	s.FsId = &v
	return s
}

func (s *ResizeDbfsRequest) SetNewSizeG(v int32) *ResizeDbfsRequest {
	s.NewSizeG = &v
	return s
}

func (s *ResizeDbfsRequest) SetRegionId(v string) *ResizeDbfsRequest {
	s.RegionId = &v
	return s
}

type ResizeDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeDbfsResponseBody) SetRequestId(v string) *ResizeDbfsResponseBody {
	s.RequestId = &v
	return s
}

type ResizeDbfsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResizeDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeDbfsResponse) GoString() string {
	return s.String()
}

func (s *ResizeDbfsResponse) SetHeaders(v map[string]*string) *ResizeDbfsResponse {
	s.Headers = v
	return s
}

func (s *ResizeDbfsResponse) SetBody(v *ResizeDbfsResponseBody) *ResizeDbfsResponse {
	s.Body = v
	return s
}

type TagDbfsRequest struct {
	DbfsId   *string `json:"DbfsId,omitempty" xml:"DbfsId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s TagDbfsRequest) GoString() string {
	return s.String()
}

func (s *TagDbfsRequest) SetDbfsId(v string) *TagDbfsRequest {
	s.DbfsId = &v
	return s
}

func (s *TagDbfsRequest) SetRegionId(v string) *TagDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *TagDbfsRequest) SetTags(v string) *TagDbfsRequest {
	s.Tags = &v
	return s
}

type TagDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *TagDbfsResponseBody) SetRequestId(v string) *TagDbfsResponseBody {
	s.RequestId = &v
	return s
}

type TagDbfsResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s TagDbfsResponse) GoString() string {
	return s.String()
}

func (s *TagDbfsResponse) SetHeaders(v map[string]*string) *TagDbfsResponse {
	s.Headers = v
	return s
}

func (s *TagDbfsResponse) SetBody(v *TagDbfsResponseBody) *TagDbfsResponse {
	s.Body = v
	return s
}

type UpdateConstantsRequest struct {
	ConstantsData *string `json:"ConstantsData,omitempty" xml:"ConstantsData,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateConstantsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConstantsRequest) GoString() string {
	return s.String()
}

func (s *UpdateConstantsRequest) SetConstantsData(v string) *UpdateConstantsRequest {
	s.ConstantsData = &v
	return s
}

func (s *UpdateConstantsRequest) SetPageNumber(v int32) *UpdateConstantsRequest {
	s.PageNumber = &v
	return s
}

func (s *UpdateConstantsRequest) SetPageSize(v int32) *UpdateConstantsRequest {
	s.PageSize = &v
	return s
}

func (s *UpdateConstantsRequest) SetRegionId(v string) *UpdateConstantsRequest {
	s.RegionId = &v
	return s
}

type UpdateConstantsResponseBody struct {
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s UpdateConstantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConstantsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConstantsResponseBody) SetData(v string) *UpdateConstantsResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateConstantsResponseBody) SetPageNumber(v int64) *UpdateConstantsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *UpdateConstantsResponseBody) SetPageSize(v int64) *UpdateConstantsResponseBody {
	s.PageSize = &v
	return s
}

func (s *UpdateConstantsResponseBody) SetRequestId(v string) *UpdateConstantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConstantsResponseBody) SetTotalCount(v int64) *UpdateConstantsResponseBody {
	s.TotalCount = &v
	return s
}

type UpdateConstantsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConstantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConstantsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConstantsResponse) GoString() string {
	return s.String()
}

func (s *UpdateConstantsResponse) SetHeaders(v map[string]*string) *UpdateConstantsResponse {
	s.Headers = v
	return s
}

func (s *UpdateConstantsResponse) SetBody(v *UpdateConstantsResponseBody) *UpdateConstantsResponse {
	s.Body = v
	return s
}

type UpdateTaskRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskIds      *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	TaskProgress *int32  `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
}

func (s UpdateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequest) SetRegionId(v string) *UpdateTaskRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTaskRequest) SetTaskIds(v string) *UpdateTaskRequest {
	s.TaskIds = &v
	return s
}

func (s *UpdateTaskRequest) SetTaskProgress(v int32) *UpdateTaskRequest {
	s.TaskProgress = &v
	return s
}

type UpdateTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskResponseBody) SetRequestId(v string) *UpdateTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTaskResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskResponse) SetHeaders(v map[string]*string) *UpdateTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskResponse) SetBody(v *UpdateTaskResponseBody) *UpdateTaskResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("dbfs.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("dbfs.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("dbfs.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("dbfs.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("dbfs.aliyuncs.com"),
		"cn-edge-1":                   tea.String("dbfs.aliyuncs.com"),
		"cn-fujian":                   tea.String("dbfs.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("dbfs.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("dbfs.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("dbfs.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("dbfs.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("dbfs.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("dbfs.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("dbfs.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("dbfs.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("dbfs.aliyuncs.com"),
		"cn-wuhan":                    tea.String("dbfs.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("dbfs.aliyuncs.com"),
		"cn-yushanfang":               tea.String("dbfs.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("dbfs.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("dbfs.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("dbfs.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("dbfs.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("dbfs.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("dbfs.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dbfs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTagsBatchWithOptions(request *AddTagsBatchRequest, runtime *util.RuntimeOptions) (_result *AddTagsBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddTagsBatchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddTagsBatch"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTagsBatch(request *AddTagsBatchRequest) (_result *AddTagsBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTagsBatchResponse{}
	_body, _err := client.AddTagsBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachDbfsWithOptions(request *AttachDbfsRequest, runtime *util.RuntimeOptions) (_result *AttachDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachDbfs(request *AttachDbfsRequest) (_result *AttachDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachDbfsResponse{}
	_body, _err := client.AttachDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConstantsWithOptions(request *CreateConstantsRequest, runtime *util.RuntimeOptions) (_result *CreateConstantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateConstantsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateConstants"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConstants(request *CreateConstantsRequest) (_result *CreateConstantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConstantsResponse{}
	_body, _err := client.CreateConstantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDbfsWithOptions(request *CreateDbfsRequest, runtime *util.RuntimeOptions) (_result *CreateDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDbfs(request *CreateDbfsRequest) (_result *CreateDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDbfsResponse{}
	_body, _err := client.CreateDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServiceLinkedRole"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSnapshot"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConstantsWithOptions(request *DeleteConstantsRequest, runtime *util.RuntimeOptions) (_result *DeleteConstantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteConstantsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteConstants"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConstants(request *DeleteConstantsRequest) (_result *DeleteConstantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConstantsResponse{}
	_body, _err := client.DeleteConstantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDbfsWithOptions(request *DeleteDbfsRequest, runtime *util.RuntimeOptions) (_result *DeleteDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDbfs(request *DeleteDbfsRequest) (_result *DeleteDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDbfsResponse{}
	_body, _err := client.DeleteDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSnapshot"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTagsBatchWithOptions(request *DeleteTagsBatchRequest, runtime *util.RuntimeOptions) (_result *DeleteTagsBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTagsBatchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTagsBatch"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTagsBatch(request *DeleteTagsBatchRequest) (_result *DeleteTagsBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTagsBatchResponse{}
	_body, _err := client.DeleteTagsBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDbfsSpecificationsWithOptions(request *DescribeDbfsSpecificationsRequest, runtime *util.RuntimeOptions) (_result *DescribeDbfsSpecificationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDbfsSpecificationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDbfsSpecifications"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDbfsSpecifications(request *DescribeDbfsSpecificationsRequest) (_result *DescribeDbfsSpecificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDbfsSpecificationsResponse{}
	_body, _err := client.DescribeDbfsSpecificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTypesWithOptions(request *DescribeInstanceTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceTypes"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (_result *DescribeInstanceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.DescribeInstanceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachDbfsWithOptions(request *DetachDbfsRequest, runtime *util.RuntimeOptions) (_result *DetachDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachDbfs(request *DetachDbfsRequest) (_result *DetachDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachDbfsResponse{}
	_body, _err := client.DetachDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDbfsWithOptions(request *GetDbfsRequest, runtime *util.RuntimeOptions) (_result *GetDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDbfs(request *GetDbfsRequest) (_result *GetDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDbfsResponse{}
	_body, _err := client.GetDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceLinkedRoleWithOptions(request *GetServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *GetServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetServiceLinkedRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetServiceLinkedRole"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceLinkedRole(request *GetServiceLinkedRoleRequest) (_result *GetServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceLinkedRoleResponse{}
	_body, _err := client.GetServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConstantsWithOptions(request *ListConstantsRequest, runtime *util.RuntimeOptions) (_result *ListConstantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConstantsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConstants"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConstants(request *ListConstantsRequest) (_result *ListConstantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConstantsResponse{}
	_body, _err := client.ListConstantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDbfsWithOptions(request *ListDbfsRequest, runtime *util.RuntimeOptions) (_result *ListDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDbfs(request *ListDbfsRequest) (_result *ListDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDbfsResponse{}
	_body, _err := client.ListDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSnapshotWithOptions(request *ListSnapshotRequest, runtime *util.RuntimeOptions) (_result *ListSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSnapshot"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSnapshot(request *ListSnapshotRequest) (_result *ListSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSnapshotResponse{}
	_body, _err := client.ListSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagValues"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaskWithOptions(request *ListTaskRequest, runtime *util.RuntimeOptions) (_result *ListTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTask"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTask(request *ListTaskRequest) (_result *ListTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskResponse{}
	_body, _err := client.ListTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpreateConstantsWithOptions(request *OpreateConstantsRequest, runtime *util.RuntimeOptions) (_result *OpreateConstantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpreateConstantsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpreateConstants"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpreateConstants(request *OpreateConstantsRequest) (_result *OpreateConstantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpreateConstantsResponse{}
	_body, _err := client.OpreateConstantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameDbfsWithOptions(request *RenameDbfsRequest, runtime *util.RuntimeOptions) (_result *RenameDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenameDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenameDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameDbfs(request *RenameDbfsRequest) (_result *RenameDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameDbfsResponse{}
	_body, _err := client.RenameDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetDbfsWithOptions(request *ResetDbfsRequest, runtime *util.RuntimeOptions) (_result *ResetDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetDbfs(request *ResetDbfsRequest) (_result *ResetDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetDbfsResponse{}
	_body, _err := client.ResetDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeDbfsWithOptions(request *ResizeDbfsRequest, runtime *util.RuntimeOptions) (_result *ResizeDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResizeDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResizeDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeDbfs(request *ResizeDbfsRequest) (_result *ResizeDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeDbfsResponse{}
	_body, _err := client.ResizeDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagDbfsWithOptions(request *TagDbfsRequest, runtime *util.RuntimeOptions) (_result *TagDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagDbfsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagDbfs"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagDbfs(request *TagDbfsRequest) (_result *TagDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagDbfsResponse{}
	_body, _err := client.TagDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConstantsWithOptions(request *UpdateConstantsRequest, runtime *util.RuntimeOptions) (_result *UpdateConstantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateConstantsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateConstants"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConstants(request *UpdateConstantsRequest) (_result *UpdateConstantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConstantsResponse{}
	_body, _err := client.UpdateConstantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskWithOptions(request *UpdateTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTask"), tea.String("2020-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTask(request *UpdateTaskRequest) (_result *UpdateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTaskResponse{}
	_body, _err := client.UpdateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
