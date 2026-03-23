// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *ListJobOrdersRequest
	GetAppCode() *string
	SetAppName(v string) *ListJobOrdersRequest
	GetAppName() *string
	SetChangingType(v string) *ListJobOrdersRequest
	GetChangingType() *string
	SetClientSystem(v string) *ListJobOrdersRequest
	GetClientSystem() *string
	SetClientUniqueId(v string) *ListJobOrdersRequest
	GetClientUniqueId() *string
	SetCursor(v int64) *ListJobOrdersRequest
	GetCursor() *int64
	SetEndTime(v string) *ListJobOrdersRequest
	GetEndTime() *string
	SetHandler(v string) *ListJobOrdersRequest
	GetHandler() *string
	SetId(v string) *ListJobOrdersRequest
	GetId() *string
	SetOrderStatus(v string) *ListJobOrdersRequest
	GetOrderStatus() *string
	SetPageSize(v int32) *ListJobOrdersRequest
	GetPageSize() *int32
	SetRequestId(v string) *ListJobOrdersRequest
	GetRequestId() *string
	SetStartTime(v string) *ListJobOrdersRequest
	GetStartTime() *string
	SetStatus(v string) *ListJobOrdersRequest
	GetStatus() *string
	SetTitle(v string) *ListJobOrdersRequest
	GetTitle() *string
}

type ListJobOrdersRequest struct {
	AppCode        *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ChangingType   *string `json:"ChangingType,omitempty" xml:"ChangingType,omitempty"`
	ClientSystem   *string `json:"ClientSystem,omitempty" xml:"ClientSystem,omitempty"`
	ClientUniqueId *string `json:"ClientUniqueId,omitempty" xml:"ClientUniqueId,omitempty"`
	// This parameter is required.
	Cursor      *int64  `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Handler     *string `json:"Handler,omitempty" xml:"Handler,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListJobOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListJobOrdersRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *ListJobOrdersRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListJobOrdersRequest) GetChangingType() *string {
	return s.ChangingType
}

func (s *ListJobOrdersRequest) GetClientSystem() *string {
	return s.ClientSystem
}

func (s *ListJobOrdersRequest) GetClientUniqueId() *string {
	return s.ClientUniqueId
}

func (s *ListJobOrdersRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *ListJobOrdersRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListJobOrdersRequest) GetHandler() *string {
	return s.Handler
}

func (s *ListJobOrdersRequest) GetId() *string {
	return s.Id
}

func (s *ListJobOrdersRequest) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *ListJobOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobOrdersRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobOrdersRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobOrdersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListJobOrdersRequest) GetTitle() *string {
	return s.Title
}

func (s *ListJobOrdersRequest) SetAppCode(v string) *ListJobOrdersRequest {
	s.AppCode = &v
	return s
}

func (s *ListJobOrdersRequest) SetAppName(v string) *ListJobOrdersRequest {
	s.AppName = &v
	return s
}

func (s *ListJobOrdersRequest) SetChangingType(v string) *ListJobOrdersRequest {
	s.ChangingType = &v
	return s
}

func (s *ListJobOrdersRequest) SetClientSystem(v string) *ListJobOrdersRequest {
	s.ClientSystem = &v
	return s
}

func (s *ListJobOrdersRequest) SetClientUniqueId(v string) *ListJobOrdersRequest {
	s.ClientUniqueId = &v
	return s
}

func (s *ListJobOrdersRequest) SetCursor(v int64) *ListJobOrdersRequest {
	s.Cursor = &v
	return s
}

func (s *ListJobOrdersRequest) SetEndTime(v string) *ListJobOrdersRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobOrdersRequest) SetHandler(v string) *ListJobOrdersRequest {
	s.Handler = &v
	return s
}

func (s *ListJobOrdersRequest) SetId(v string) *ListJobOrdersRequest {
	s.Id = &v
	return s
}

func (s *ListJobOrdersRequest) SetOrderStatus(v string) *ListJobOrdersRequest {
	s.OrderStatus = &v
	return s
}

func (s *ListJobOrdersRequest) SetPageSize(v int32) *ListJobOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobOrdersRequest) SetRequestId(v string) *ListJobOrdersRequest {
	s.RequestId = &v
	return s
}

func (s *ListJobOrdersRequest) SetStartTime(v string) *ListJobOrdersRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobOrdersRequest) SetStatus(v string) *ListJobOrdersRequest {
	s.Status = &v
	return s
}

func (s *ListJobOrdersRequest) SetTitle(v string) *ListJobOrdersRequest {
	s.Title = &v
	return s
}

func (s *ListJobOrdersRequest) Validate() error {
	return dara.Validate(s)
}
