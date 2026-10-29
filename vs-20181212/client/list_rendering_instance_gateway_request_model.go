// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingInstanceGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListRenderingInstanceGatewayRequest
	GetEndTime() *string
	SetGatewayInstanceId(v string) *ListRenderingInstanceGatewayRequest
	GetGatewayInstanceId() *string
	SetPageNumber(v int64) *ListRenderingInstanceGatewayRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRenderingInstanceGatewayRequest
	GetPageSize() *int64
	SetRenderingInstanceId(v string) *ListRenderingInstanceGatewayRequest
	GetRenderingInstanceId() *string
	SetStartTime(v string) *ListRenderingInstanceGatewayRequest
	GetStartTime() *string
}

type ListRenderingInstanceGatewayRequest struct {
	// The time range filter parameter. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-05-20T12:23:02Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID of the gateway.
	//
	// example:
	//
	// render-xxx
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" xml:"GatewayInstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in the paging Settings for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The instance ID of the cloud application service.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The time range filter parameter. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-05-18T02:20:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRenderingInstanceGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstanceGatewayRequest) GoString() string {
	return s.String()
}

func (s *ListRenderingInstanceGatewayRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRenderingInstanceGatewayRequest) GetGatewayInstanceId() *string {
	return s.GatewayInstanceId
}

func (s *ListRenderingInstanceGatewayRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRenderingInstanceGatewayRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRenderingInstanceGatewayRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingInstanceGatewayRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRenderingInstanceGatewayRequest) SetEndTime(v string) *ListRenderingInstanceGatewayRequest {
	s.EndTime = &v
	return s
}

func (s *ListRenderingInstanceGatewayRequest) SetGatewayInstanceId(v string) *ListRenderingInstanceGatewayRequest {
	s.GatewayInstanceId = &v
	return s
}

func (s *ListRenderingInstanceGatewayRequest) SetPageNumber(v int64) *ListRenderingInstanceGatewayRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRenderingInstanceGatewayRequest) SetPageSize(v int64) *ListRenderingInstanceGatewayRequest {
	s.PageSize = &v
	return s
}

func (s *ListRenderingInstanceGatewayRequest) SetRenderingInstanceId(v string) *ListRenderingInstanceGatewayRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingInstanceGatewayRequest) SetStartTime(v string) *ListRenderingInstanceGatewayRequest {
	s.StartTime = &v
	return s
}

func (s *ListRenderingInstanceGatewayRequest) Validate() error {
	return dara.Validate(s)
}
