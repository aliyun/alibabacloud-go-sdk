// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpfilterListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetIpfilterListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetIpfilterListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetIpfilterListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetIpfilterListResponseBody
	GetTotalCount() *int32
	SetData(v *GetIpfilterListResponseBodyData) *GetIpfilterListResponseBody
	GetData() *GetIpfilterListResponseBodyData
}

type GetIpfilterListResponseBody struct {
	// Current page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 84DD77C7-A091-5139-9530-2D1F7CCE59E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Data records
	Data *GetIpfilterListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetIpfilterListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIpfilterListResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetIpfilterListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetIpfilterListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIpfilterListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetIpfilterListResponseBody) GetData() *GetIpfilterListResponseBodyData {
	return s.Data
}

func (s *GetIpfilterListResponseBody) SetPageNumber(v int32) *GetIpfilterListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetIpfilterListResponseBody) SetPageSize(v int32) *GetIpfilterListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetIpfilterListResponseBody) SetRequestId(v string) *GetIpfilterListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpfilterListResponseBody) SetTotalCount(v int32) *GetIpfilterListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetIpfilterListResponseBody) SetData(v *GetIpfilterListResponseBodyData) *GetIpfilterListResponseBody {
	s.Data = v
	return s
}

func (s *GetIpfilterListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIpfilterListResponseBodyData struct {
	Ipfilters []*GetIpfilterListResponseBodyDataIpfilters `json:"ipfilters,omitempty" xml:"ipfilters,omitempty" type:"Repeated"`
}

func (s GetIpfilterListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIpfilterListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponseBodyData) GetIpfilters() []*GetIpfilterListResponseBodyDataIpfilters {
	return s.Ipfilters
}

func (s *GetIpfilterListResponseBodyData) SetIpfilters(v []*GetIpfilterListResponseBodyDataIpfilters) *GetIpfilterListResponseBodyData {
	s.Ipfilters = v
	return s
}

func (s *GetIpfilterListResponseBodyData) Validate() error {
	if s.Ipfilters != nil {
		for _, item := range s.Ipfilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetIpfilterListResponseBodyDataIpfilters struct {
	// timestamp
	//
	// example:
	//
	// 1653547140
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Record ID
	//
	// example:
	//
	// 10083
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// IP address/IP range/IP segment
	//
	// example:
	//
	// xxx.xxx.xxx.xxx
	//
	// xxx.xxx.xxx.xxx-xxx.xxx.xxx.xxx
	//
	// xxx.xxx.xxx.xxx/xxx
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
}

func (s GetIpfilterListResponseBodyDataIpfilters) String() string {
	return dara.Prettify(s)
}

func (s GetIpfilterListResponseBodyDataIpfilters) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponseBodyDataIpfilters) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetIpfilterListResponseBodyDataIpfilters) GetId() *string {
	return s.Id
}

func (s *GetIpfilterListResponseBodyDataIpfilters) GetIpAddress() *string {
	return s.IpAddress
}

func (s *GetIpfilterListResponseBodyDataIpfilters) SetCreateTime(v string) *GetIpfilterListResponseBodyDataIpfilters {
	s.CreateTime = &v
	return s
}

func (s *GetIpfilterListResponseBodyDataIpfilters) SetId(v string) *GetIpfilterListResponseBodyDataIpfilters {
	s.Id = &v
	return s
}

func (s *GetIpfilterListResponseBodyDataIpfilters) SetIpAddress(v string) *GetIpfilterListResponseBodyDataIpfilters {
	s.IpAddress = &v
	return s
}

func (s *GetIpfilterListResponseBodyDataIpfilters) Validate() error {
	return dara.Validate(s)
}
