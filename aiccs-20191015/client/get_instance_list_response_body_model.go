// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceListResponseBody
	GetCode() *string
	SetCommodityInstances(v []*GetInstanceListResponseBodyCommodityInstances) *GetInstanceListResponseBody
	GetCommodityInstances() []*GetInstanceListResponseBodyCommodityInstances
	SetHttpStatusCode(v int32) *GetInstanceListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceListResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetInstanceListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetInstanceListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetInstanceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetInstanceListResponseBody
	GetTotalCount() *int32
}

type GetInstanceListResponseBody struct {
	Code               *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	CommodityInstances []*GetInstanceListResponseBodyCommodityInstances `json:"CommodityInstances,omitempty" xml:"CommodityInstances,omitempty" type:"Repeated"`
	HttpStatusCode     *int32                                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber         *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount         *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceListResponseBody) GetCommodityInstances() []*GetInstanceListResponseBodyCommodityInstances {
	return s.CommodityInstances
}

func (s *GetInstanceListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetInstanceListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetInstanceListResponseBody) SetCode(v string) *GetInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceListResponseBody) SetCommodityInstances(v []*GetInstanceListResponseBodyCommodityInstances) *GetInstanceListResponseBody {
	s.CommodityInstances = v
	return s
}

func (s *GetInstanceListResponseBody) SetHttpStatusCode(v int32) *GetInstanceListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceListResponseBody) SetMessage(v string) *GetInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceListResponseBody) SetPageNumber(v int32) *GetInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceListResponseBody) SetPageSize(v int32) *GetInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetInstanceListResponseBody) SetRequestId(v string) *GetInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceListResponseBody) SetSuccess(v bool) *GetInstanceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceListResponseBody) SetTotalCount(v int32) *GetInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceListResponseBody) Validate() error {
	if s.CommodityInstances != nil {
		for _, item := range s.CommodityInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceListResponseBodyCommodityInstances struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetInstanceListResponseBodyCommodityInstances) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyCommodityInstances) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyCommodityInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceListResponseBodyCommodityInstances) GetName() *string {
	return s.Name
}

func (s *GetInstanceListResponseBodyCommodityInstances) SetInstanceId(v string) *GetInstanceListResponseBodyCommodityInstances {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceListResponseBodyCommodityInstances) SetName(v string) *GetInstanceListResponseBodyCommodityInstances {
	s.Name = &v
	return s
}

func (s *GetInstanceListResponseBodyCommodityInstances) Validate() error {
	return dara.Validate(s)
}
