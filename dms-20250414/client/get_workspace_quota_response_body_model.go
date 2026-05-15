// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetWorkspaceQuotaResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *GetWorkspaceQuotaResponseBodyData) *GetWorkspaceQuotaResponseBody
	GetData() *GetWorkspaceQuotaResponseBodyData
	SetErrorCode(v string) *GetWorkspaceQuotaResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *GetWorkspaceQuotaResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetWorkspaceQuotaResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkspaceQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkspaceQuotaResponseBody
	GetSuccess() *bool
}

type GetWorkspaceQuotaResponseBody struct {
	// example:
	//
	// NOT_FOUND
	AccessDeniedDetail *string                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *GetWorkspaceQuotaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkspaceQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceQuotaResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetWorkspaceQuotaResponseBody) GetData() *GetWorkspaceQuotaResponseBodyData {
	return s.Data
}

func (s *GetWorkspaceQuotaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkspaceQuotaResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetWorkspaceQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkspaceQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkspaceQuotaResponseBody) SetAccessDeniedDetail(v string) *GetWorkspaceQuotaResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBody) SetData(v *GetWorkspaceQuotaResponseBodyData) *GetWorkspaceQuotaResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkspaceQuotaResponseBody) SetErrorCode(v string) *GetWorkspaceQuotaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBody) SetHttpStatusCode(v int64) *GetWorkspaceQuotaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBody) SetMessage(v string) *GetWorkspaceQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBody) SetRequestId(v string) *GetWorkspaceQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBody) SetSuccess(v bool) *GetWorkspaceQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceQuotaResponseBodyData struct {
	// example:
	//
	// 200
	CuQuota *int64 `json:"CuQuota,omitempty" xml:"CuQuota,omitempty"`
	// example:
	//
	// 20
	CuQuotaUsage *int64 `json:"CuQuotaUsage,omitempty" xml:"CuQuotaUsage,omitempty"`
	// example:
	//
	// i-bp16pha0zu99yybi59qr
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0
	NotebookFreeQuotaAvailable *int64 `json:"NotebookFreeQuotaAvailable,omitempty" xml:"NotebookFreeQuotaAvailable,omitempty"`
	// example:
	//
	// 0
	NotebookFreeQuotaTotal *int64 `json:"NotebookFreeQuotaTotal,omitempty" xml:"NotebookFreeQuotaTotal,omitempty"`
	// example:
	//
	// RELEASED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorkspaceQuotaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkspaceQuotaResponseBodyData) GetCuQuota() *int64 {
	return s.CuQuota
}

func (s *GetWorkspaceQuotaResponseBodyData) GetCuQuotaUsage() *int64 {
	return s.CuQuotaUsage
}

func (s *GetWorkspaceQuotaResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkspaceQuotaResponseBodyData) GetNotebookFreeQuotaAvailable() *int64 {
	return s.NotebookFreeQuotaAvailable
}

func (s *GetWorkspaceQuotaResponseBodyData) GetNotebookFreeQuotaTotal() *int64 {
	return s.NotebookFreeQuotaTotal
}

func (s *GetWorkspaceQuotaResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetWorkspaceQuotaResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceQuotaResponseBodyData) SetCuQuota(v int64) *GetWorkspaceQuotaResponseBodyData {
	s.CuQuota = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBodyData) SetCuQuotaUsage(v int64) *GetWorkspaceQuotaResponseBodyData {
	s.CuQuotaUsage = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBodyData) SetInstanceId(v string) *GetWorkspaceQuotaResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBodyData) SetNotebookFreeQuotaAvailable(v int64) *GetWorkspaceQuotaResponseBodyData {
	s.NotebookFreeQuotaAvailable = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBodyData) SetNotebookFreeQuotaTotal(v int64) *GetWorkspaceQuotaResponseBodyData {
	s.NotebookFreeQuotaTotal = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBodyData) SetStatus(v string) *GetWorkspaceQuotaResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBodyData) SetWorkspaceId(v string) *GetWorkspaceQuotaResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceQuotaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
