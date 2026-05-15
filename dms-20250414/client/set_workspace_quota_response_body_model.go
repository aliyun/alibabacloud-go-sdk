// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkspaceQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SetWorkspaceQuotaResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *SetWorkspaceQuotaResponseBodyData) *SetWorkspaceQuotaResponseBody
	GetData() *SetWorkspaceQuotaResponseBodyData
	SetErrorCode(v string) *SetWorkspaceQuotaResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *SetWorkspaceQuotaResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *SetWorkspaceQuotaResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetWorkspaceQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetWorkspaceQuotaResponseBody
	GetSuccess() *bool
}

type SetWorkspaceQuotaResponseBody struct {
	// example:
	//
	// NOT_FOUND
	AccessDeniedDetail *string                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *SetWorkspaceQuotaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// This record is being collected, please wait for a moment.
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

func (s SetWorkspaceQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetWorkspaceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetWorkspaceQuotaResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SetWorkspaceQuotaResponseBody) GetData() *SetWorkspaceQuotaResponseBodyData {
	return s.Data
}

func (s *SetWorkspaceQuotaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SetWorkspaceQuotaResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *SetWorkspaceQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetWorkspaceQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetWorkspaceQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetWorkspaceQuotaResponseBody) SetAccessDeniedDetail(v string) *SetWorkspaceQuotaResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBody) SetData(v *SetWorkspaceQuotaResponseBodyData) *SetWorkspaceQuotaResponseBody {
	s.Data = v
	return s
}

func (s *SetWorkspaceQuotaResponseBody) SetErrorCode(v string) *SetWorkspaceQuotaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBody) SetHttpStatusCode(v int64) *SetWorkspaceQuotaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBody) SetMessage(v string) *SetWorkspaceQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBody) SetRequestId(v string) *SetWorkspaceQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBody) SetSuccess(v bool) *SetWorkspaceQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetWorkspaceQuotaResponseBodyData struct {
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
	// i-2zehld3y1tphzctyyq7o
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 263013787210103
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// Catched
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SetWorkspaceQuotaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetWorkspaceQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetWorkspaceQuotaResponseBodyData) GetCuQuota() *int64 {
	return s.CuQuota
}

func (s *SetWorkspaceQuotaResponseBodyData) GetCuQuotaUsage() *int64 {
	return s.CuQuotaUsage
}

func (s *SetWorkspaceQuotaResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetWorkspaceQuotaResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *SetWorkspaceQuotaResponseBodyData) GetState() *string {
	return s.State
}

func (s *SetWorkspaceQuotaResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SetWorkspaceQuotaResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SetWorkspaceQuotaResponseBodyData) SetCuQuota(v int64) *SetWorkspaceQuotaResponseBodyData {
	s.CuQuota = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBodyData) SetCuQuotaUsage(v int64) *SetWorkspaceQuotaResponseBodyData {
	s.CuQuotaUsage = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBodyData) SetInstanceId(v string) *SetWorkspaceQuotaResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBodyData) SetOrderId(v string) *SetWorkspaceQuotaResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBodyData) SetState(v string) *SetWorkspaceQuotaResponseBodyData {
	s.State = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBodyData) SetStatus(v string) *SetWorkspaceQuotaResponseBodyData {
	s.Status = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBodyData) SetWorkspaceId(v string) *SetWorkspaceQuotaResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *SetWorkspaceQuotaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
