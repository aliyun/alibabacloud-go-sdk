// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIdListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetInstanceIdListRequest
	GetAppType() *string
	SetApprovedResult(v string) *GetInstanceIdListRequest
	GetApprovedResult() *string
	SetCreateFromTimeGMT(v string) *GetInstanceIdListRequest
	GetCreateFromTimeGMT() *string
	SetCreateToTimeGMT(v string) *GetInstanceIdListRequest
	GetCreateToTimeGMT() *string
	SetFormUuid(v string) *GetInstanceIdListRequest
	GetFormUuid() *string
	SetInstanceStatus(v string) *GetInstanceIdListRequest
	GetInstanceStatus() *string
	SetLanguage(v string) *GetInstanceIdListRequest
	GetLanguage() *string
	SetModifiedFromTimeGMT(v string) *GetInstanceIdListRequest
	GetModifiedFromTimeGMT() *string
	SetModifiedToTimeGMT(v string) *GetInstanceIdListRequest
	GetModifiedToTimeGMT() *string
	SetOriginatorId(v string) *GetInstanceIdListRequest
	GetOriginatorId() *string
	SetPageNumber(v int32) *GetInstanceIdListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetInstanceIdListRequest
	GetPageSize() *int32
	SetSearchFieldJson(v string) *GetInstanceIdListRequest
	GetSearchFieldJson() *string
	SetSystemToken(v string) *GetInstanceIdListRequest
	GetSystemToken() *string
	SetTaskId(v string) *GetInstanceIdListRequest
	GetTaskId() *string
}

type GetInstanceIdListRequest struct {
	// example:
	//
	// APP_PBxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// agree
	ApprovedResult *string `json:"ApprovedResult,omitempty" xml:"ApprovedResult,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateFromTimeGMT *string `json:"CreateFromTimeGMT,omitempty" xml:"CreateFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateToTimeGMT *string `json:"CreateToTimeGMT,omitempty" xml:"CreateToTimeGMT,omitempty"`
	// example:
	//
	// FORM-EF6Yxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 2021-05-01
	ModifiedFromTimeGMT *string `json:"ModifiedFromTimeGMT,omitempty" xml:"ModifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-09-10
	ModifiedToTimeGMT *string `json:"ModifiedToTimeGMT,omitempty" xml:"ModifiedToTimeGMT,omitempty"`
	// example:
	//
	// 012345
	OriginatorId *string `json:"OriginatorId,omitempty" xml:"OriginatorId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {\\"textField\\":\\"123\\"}
	SearchFieldJson *string `json:"SearchFieldJson,omitempty" xml:"SearchFieldJson,omitempty"`
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	// example:
	//
	// 1045001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetInstanceIdListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIdListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetInstanceIdListRequest) GetApprovedResult() *string {
	return s.ApprovedResult
}

func (s *GetInstanceIdListRequest) GetCreateFromTimeGMT() *string {
	return s.CreateFromTimeGMT
}

func (s *GetInstanceIdListRequest) GetCreateToTimeGMT() *string {
	return s.CreateToTimeGMT
}

func (s *GetInstanceIdListRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetInstanceIdListRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstanceIdListRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetInstanceIdListRequest) GetModifiedFromTimeGMT() *string {
	return s.ModifiedFromTimeGMT
}

func (s *GetInstanceIdListRequest) GetModifiedToTimeGMT() *string {
	return s.ModifiedToTimeGMT
}

func (s *GetInstanceIdListRequest) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *GetInstanceIdListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetInstanceIdListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetInstanceIdListRequest) GetSearchFieldJson() *string {
	return s.SearchFieldJson
}

func (s *GetInstanceIdListRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetInstanceIdListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetInstanceIdListRequest) SetAppType(v string) *GetInstanceIdListRequest {
	s.AppType = &v
	return s
}

func (s *GetInstanceIdListRequest) SetApprovedResult(v string) *GetInstanceIdListRequest {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstanceIdListRequest) SetCreateFromTimeGMT(v string) *GetInstanceIdListRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetCreateToTimeGMT(v string) *GetInstanceIdListRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetFormUuid(v string) *GetInstanceIdListRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceIdListRequest) SetInstanceStatus(v string) *GetInstanceIdListRequest {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceIdListRequest) SetLanguage(v string) *GetInstanceIdListRequest {
	s.Language = &v
	return s
}

func (s *GetInstanceIdListRequest) SetModifiedFromTimeGMT(v string) *GetInstanceIdListRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetModifiedToTimeGMT(v string) *GetInstanceIdListRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetOriginatorId(v string) *GetInstanceIdListRequest {
	s.OriginatorId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetPageNumber(v int32) *GetInstanceIdListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceIdListRequest) SetPageSize(v int32) *GetInstanceIdListRequest {
	s.PageSize = &v
	return s
}

func (s *GetInstanceIdListRequest) SetSearchFieldJson(v string) *GetInstanceIdListRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *GetInstanceIdListRequest) SetSystemToken(v string) *GetInstanceIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstanceIdListRequest) SetTaskId(v string) *GetInstanceIdListRequest {
	s.TaskId = &v
	return s
}

func (s *GetInstanceIdListRequest) Validate() error {
	return dara.Validate(s)
}
