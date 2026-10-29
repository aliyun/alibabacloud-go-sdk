// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingInstanceGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayConfigurationInfos(v []*ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) *ListRenderingInstanceGatewayResponseBody
	GetGatewayConfigurationInfos() []*ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos
	SetPageNumber(v string) *ListRenderingInstanceGatewayResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListRenderingInstanceGatewayResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListRenderingInstanceGatewayResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRenderingInstanceGatewayResponseBody
	GetTotalCount() *string
}

type ListRenderingInstanceGatewayResponseBody struct {
	// The list of custom gateway information.
	GatewayConfigurationInfos []*ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos `json:"GatewayConfigurationInfos,omitempty" xml:"GatewayConfigurationInfos,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in the paging Settings for a paged query.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of installation records.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRenderingInstanceGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstanceGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ListRenderingInstanceGatewayResponseBody) GetGatewayConfigurationInfos() []*ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos {
	return s.GatewayConfigurationInfos
}

func (s *ListRenderingInstanceGatewayResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListRenderingInstanceGatewayResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListRenderingInstanceGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRenderingInstanceGatewayResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRenderingInstanceGatewayResponseBody) SetGatewayConfigurationInfos(v []*ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) *ListRenderingInstanceGatewayResponseBody {
	s.GatewayConfigurationInfos = v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBody) SetPageNumber(v string) *ListRenderingInstanceGatewayResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBody) SetPageSize(v string) *ListRenderingInstanceGatewayResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBody) SetRequestId(v string) *ListRenderingInstanceGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBody) SetTotalCount(v string) *ListRenderingInstanceGatewayResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBody) Validate() error {
	if s.GatewayConfigurationInfos != nil {
		for _, item := range s.GatewayConfigurationInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos struct {
	// The creation time of the cloud application service instance. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2024-10-15 10:19:13
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The instance ID of the custom gateway.
	//
	// example:
	//
	// render-xxx
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" xml:"GatewayInstanceId,omitempty"`
	// The instance ID of the cloud application service.
	//
	// example:
	//
	// render-342012a227dc4ddf91f024639e43051a
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The status of the custom gateway. Valid values:
	//
	// - Intermediate states: creating and deleting.
	//
	// - Final states: available and failed.
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2024-11-02 12:08:26
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) GoString() string {
	return s.String()
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) GetGatewayInstanceId() *string {
	return s.GatewayInstanceId
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) GetStatus() *string {
	return s.Status
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) SetCreationTime(v string) *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos {
	s.CreationTime = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) SetGatewayInstanceId(v string) *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos {
	s.GatewayInstanceId = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) SetRenderingInstanceId(v string) *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) SetStatus(v string) *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos {
	s.Status = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) SetUpdateTime(v string) *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos {
	s.UpdateTime = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponseBodyGatewayConfigurationInfos) Validate() error {
	return dara.Validate(s)
}
