// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchTracesByPageResponseBodyPageBean) *SearchTracesByPageResponseBody
	GetPageBean() *SearchTracesByPageResponseBodyPageBean
	SetRequestId(v string) *SearchTracesByPageResponseBody
	GetRequestId() *string
}

type SearchTracesByPageResponseBody struct {
	PageBean  *SearchTracesByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTracesByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBody) GetPageBean() *SearchTracesByPageResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchTracesByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTracesByPageResponseBody) SetPageBean(v *SearchTracesByPageResponseBodyPageBean) *SearchTracesByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTracesByPageResponseBody) SetRequestId(v string) *SearchTracesByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTracesByPageResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTracesByPageResponseBodyPageBean struct {
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	TraceInfos []*SearchTracesByPageResponseBodyPageBeanTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Repeated"`
}

func (s SearchTracesByPageResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTracesByPageResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTracesByPageResponseBodyPageBean) GetTotal() *int32 {
	return s.Total
}

func (s *SearchTracesByPageResponseBodyPageBean) GetTraceInfos() []*SearchTracesByPageResponseBodyPageBeanTraceInfos {
	return s.TraceInfos
}

func (s *SearchTracesByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetPageSize(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetTotal(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetTraceInfos(v []*SearchTracesByPageResponseBodyPageBeanTraceInfos) *SearchTracesByPageResponseBodyPageBean {
	s.TraceInfos = v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) Validate() error {
	if s.TraceInfos != nil {
		for _, item := range s.TraceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTracesByPageResponseBodyPageBeanTraceInfos struct {
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetDuration() *int64 {
	return s.Duration
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetOperationName() *string {
	return s.OperationName
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetServiceName() *string {
	return s.ServiceName
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetTraceID() *string {
	return s.TraceID
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetDuration(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Duration = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetOperationName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceIp(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTimestamp(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTraceID(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.TraceID = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) Validate() error {
	return dara.Validate(s)
}
