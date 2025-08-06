// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody
	GetLogstore() *string
	SetProject(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody
	GetProject() *string
	SetRegion(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody
	GetRegion() *string
	SetRequestId(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody
	GetStatus() *string
}

type DescribeVodDomainRealtimeLogDeliveryResponseBody struct {
	Logstore  *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVodDomainRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) GetLogstore() *string {
	return s.Logstore
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) GetProject() *string {
	return s.Project
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetLogstore(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.Logstore = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetProject(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.Project = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetRegion(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetStatus(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
