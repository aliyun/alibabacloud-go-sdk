// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobOrdersWithSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *ListJobOrdersWithSizeRequest
	GetAppCode() *string
	SetAppName(v string) *ListJobOrdersWithSizeRequest
	GetAppName() *string
	SetChangingType(v string) *ListJobOrdersWithSizeRequest
	GetChangingType() *string
	SetClientSystem(v string) *ListJobOrdersWithSizeRequest
	GetClientSystem() *string
	SetClientUniqueId(v string) *ListJobOrdersWithSizeRequest
	GetClientUniqueId() *string
	SetCursor(v int64) *ListJobOrdersWithSizeRequest
	GetCursor() *int64
	SetEndTime(v string) *ListJobOrdersWithSizeRequest
	GetEndTime() *string
	SetHandler(v string) *ListJobOrdersWithSizeRequest
	GetHandler() *string
	SetId(v string) *ListJobOrdersWithSizeRequest
	GetId() *string
	SetOrderStatus(v string) *ListJobOrdersWithSizeRequest
	GetOrderStatus() *string
	SetPageSize(v int32) *ListJobOrdersWithSizeRequest
	GetPageSize() *int32
	SetRequestId(v string) *ListJobOrdersWithSizeRequest
	GetRequestId() *string
	SetStartTime(v string) *ListJobOrdersWithSizeRequest
	GetStartTime() *string
	SetStatus(v string) *ListJobOrdersWithSizeRequest
	GetStatus() *string
	SetTitle(v string) *ListJobOrdersWithSizeRequest
	GetTitle() *string
}

type ListJobOrdersWithSizeRequest struct {
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

func (s ListJobOrdersWithSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobOrdersWithSizeRequest) GoString() string {
	return s.String()
}

func (s *ListJobOrdersWithSizeRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *ListJobOrdersWithSizeRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListJobOrdersWithSizeRequest) GetChangingType() *string {
	return s.ChangingType
}

func (s *ListJobOrdersWithSizeRequest) GetClientSystem() *string {
	return s.ClientSystem
}

func (s *ListJobOrdersWithSizeRequest) GetClientUniqueId() *string {
	return s.ClientUniqueId
}

func (s *ListJobOrdersWithSizeRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *ListJobOrdersWithSizeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListJobOrdersWithSizeRequest) GetHandler() *string {
	return s.Handler
}

func (s *ListJobOrdersWithSizeRequest) GetId() *string {
	return s.Id
}

func (s *ListJobOrdersWithSizeRequest) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *ListJobOrdersWithSizeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobOrdersWithSizeRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobOrdersWithSizeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobOrdersWithSizeRequest) GetStatus() *string {
	return s.Status
}

func (s *ListJobOrdersWithSizeRequest) GetTitle() *string {
	return s.Title
}

func (s *ListJobOrdersWithSizeRequest) SetAppCode(v string) *ListJobOrdersWithSizeRequest {
	s.AppCode = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetAppName(v string) *ListJobOrdersWithSizeRequest {
	s.AppName = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetChangingType(v string) *ListJobOrdersWithSizeRequest {
	s.ChangingType = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetClientSystem(v string) *ListJobOrdersWithSizeRequest {
	s.ClientSystem = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetClientUniqueId(v string) *ListJobOrdersWithSizeRequest {
	s.ClientUniqueId = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetCursor(v int64) *ListJobOrdersWithSizeRequest {
	s.Cursor = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetEndTime(v string) *ListJobOrdersWithSizeRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetHandler(v string) *ListJobOrdersWithSizeRequest {
	s.Handler = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetId(v string) *ListJobOrdersWithSizeRequest {
	s.Id = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetOrderStatus(v string) *ListJobOrdersWithSizeRequest {
	s.OrderStatus = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetPageSize(v int32) *ListJobOrdersWithSizeRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetRequestId(v string) *ListJobOrdersWithSizeRequest {
	s.RequestId = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetStartTime(v string) *ListJobOrdersWithSizeRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetStatus(v string) *ListJobOrdersWithSizeRequest {
	s.Status = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) SetTitle(v string) *ListJobOrdersWithSizeRequest {
	s.Title = &v
	return s
}

func (s *ListJobOrdersWithSizeRequest) Validate() error {
	return dara.Validate(s)
}
