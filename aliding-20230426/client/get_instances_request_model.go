// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetInstancesRequest
	GetAppType() *string
	SetApprovedResult(v string) *GetInstancesRequest
	GetApprovedResult() *string
	SetCreateFromTimeGMT(v string) *GetInstancesRequest
	GetCreateFromTimeGMT() *string
	SetCreateToTimeGMT(v string) *GetInstancesRequest
	GetCreateToTimeGMT() *string
	SetFormUuid(v string) *GetInstancesRequest
	GetFormUuid() *string
	SetInstanceStatus(v string) *GetInstancesRequest
	GetInstanceStatus() *string
	SetLanguage(v string) *GetInstancesRequest
	GetLanguage() *string
	SetModifiedFromTimeGMT(v string) *GetInstancesRequest
	GetModifiedFromTimeGMT() *string
	SetModifiedToTimeGMT(v string) *GetInstancesRequest
	GetModifiedToTimeGMT() *string
	SetOrderConfigJson(v string) *GetInstancesRequest
	GetOrderConfigJson() *string
	SetOriginatorId(v string) *GetInstancesRequest
	GetOriginatorId() *string
	SetPageNumber(v int32) *GetInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetInstancesRequest
	GetPageSize() *int32
	SetSearchFieldJson(v string) *GetInstancesRequest
	GetSearchFieldJson() *string
	SetSystemToken(v string) *GetInstancesRequest
	GetSystemToken() *string
	SetTaskId(v string) *GetInstancesRequest
	GetTaskId() *string
}

type GetInstancesRequest struct {
	AppType             *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	ApprovedResult      *string `json:"ApprovedResult,omitempty" xml:"ApprovedResult,omitempty"`
	CreateFromTimeGMT   *string `json:"CreateFromTimeGMT,omitempty" xml:"CreateFromTimeGMT,omitempty"`
	CreateToTimeGMT     *string `json:"CreateToTimeGMT,omitempty" xml:"CreateToTimeGMT,omitempty"`
	FormUuid            *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	InstanceStatus      *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Language            *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ModifiedFromTimeGMT *string `json:"ModifiedFromTimeGMT,omitempty" xml:"ModifiedFromTimeGMT,omitempty"`
	ModifiedToTimeGMT   *string `json:"ModifiedToTimeGMT,omitempty" xml:"ModifiedToTimeGMT,omitempty"`
	OrderConfigJson     *string `json:"OrderConfigJson,omitempty" xml:"OrderConfigJson,omitempty"`
	OriginatorId        *string `json:"OriginatorId,omitempty" xml:"OriginatorId,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchFieldJson     *string `json:"SearchFieldJson,omitempty" xml:"SearchFieldJson,omitempty"`
	SystemToken         *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetInstancesRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetInstancesRequest) GetApprovedResult() *string {
	return s.ApprovedResult
}

func (s *GetInstancesRequest) GetCreateFromTimeGMT() *string {
	return s.CreateFromTimeGMT
}

func (s *GetInstancesRequest) GetCreateToTimeGMT() *string {
	return s.CreateToTimeGMT
}

func (s *GetInstancesRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetInstancesRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstancesRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetInstancesRequest) GetModifiedFromTimeGMT() *string {
	return s.ModifiedFromTimeGMT
}

func (s *GetInstancesRequest) GetModifiedToTimeGMT() *string {
	return s.ModifiedToTimeGMT
}

func (s *GetInstancesRequest) GetOrderConfigJson() *string {
	return s.OrderConfigJson
}

func (s *GetInstancesRequest) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *GetInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetInstancesRequest) GetSearchFieldJson() *string {
	return s.SearchFieldJson
}

func (s *GetInstancesRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetInstancesRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetInstancesRequest) SetAppType(v string) *GetInstancesRequest {
	s.AppType = &v
	return s
}

func (s *GetInstancesRequest) SetApprovedResult(v string) *GetInstancesRequest {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesRequest) SetCreateFromTimeGMT(v string) *GetInstancesRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetCreateToTimeGMT(v string) *GetInstancesRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetFormUuid(v string) *GetInstancesRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesRequest) SetInstanceStatus(v string) *GetInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesRequest) SetLanguage(v string) *GetInstancesRequest {
	s.Language = &v
	return s
}

func (s *GetInstancesRequest) SetModifiedFromTimeGMT(v string) *GetInstancesRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetModifiedToTimeGMT(v string) *GetInstancesRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetOrderConfigJson(v string) *GetInstancesRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *GetInstancesRequest) SetOriginatorId(v string) *GetInstancesRequest {
	s.OriginatorId = &v
	return s
}

func (s *GetInstancesRequest) SetPageNumber(v int32) *GetInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstancesRequest) SetPageSize(v int32) *GetInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *GetInstancesRequest) SetSearchFieldJson(v string) *GetInstancesRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *GetInstancesRequest) SetSystemToken(v string) *GetInstancesRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstancesRequest) SetTaskId(v string) *GetInstancesRequest {
	s.TaskId = &v
	return s
}

func (s *GetInstancesRequest) Validate() error {
	return dara.Validate(s)
}
