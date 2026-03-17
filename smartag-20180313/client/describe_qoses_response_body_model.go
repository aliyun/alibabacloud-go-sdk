// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeQosesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeQosesResponseBody
	GetPageSize() *int32
	SetQoses(v *DescribeQosesResponseBodyQoses) *DescribeQosesResponseBody
	GetQoses() *DescribeQosesResponseBodyQoses
	SetRequestId(v string) *DescribeQosesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeQosesResponseBody
	GetTotalCount() *int32
}

type DescribeQosesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Qoses    *DescribeQosesResponseBodyQoses `json:"Qoses,omitempty" xml:"Qoses,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2B5F35DD-0D66-41FC-AA99-BAE473E1A7A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of QoS polices.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeQosesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQosesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeQosesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeQosesResponseBody) GetQoses() *DescribeQosesResponseBodyQoses {
	return s.Qoses
}

func (s *DescribeQosesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQosesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeQosesResponseBody) SetPageNumber(v int32) *DescribeQosesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeQosesResponseBody) SetPageSize(v int32) *DescribeQosesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeQosesResponseBody) SetQoses(v *DescribeQosesResponseBodyQoses) *DescribeQosesResponseBody {
	s.Qoses = v
	return s
}

func (s *DescribeQosesResponseBody) SetRequestId(v string) *DescribeQosesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQosesResponseBody) SetTotalCount(v int32) *DescribeQosesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeQosesResponseBody) Validate() error {
	if s.Qoses != nil {
		if err := s.Qoses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeQosesResponseBodyQoses struct {
	Qos []*DescribeQosesResponseBodyQosesQos `json:"Qos,omitempty" xml:"Qos,omitempty" type:"Repeated"`
}

func (s DescribeQosesResponseBodyQoses) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosesResponseBodyQoses) GoString() string {
	return s.String()
}

func (s *DescribeQosesResponseBodyQoses) GetQos() []*DescribeQosesResponseBodyQosesQos {
	return s.Qos
}

func (s *DescribeQosesResponseBodyQoses) SetQos(v []*DescribeQosesResponseBodyQosesQos) *DescribeQosesResponseBodyQoses {
	s.Qos = v
	return s
}

func (s *DescribeQosesResponseBodyQoses) Validate() error {
	if s.Qos != nil {
		for _, item := range s.Qos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeQosesResponseBodyQosesQos struct {
	QosDescription  *string `json:"QosDescription,omitempty" xml:"QosDescription,omitempty"`
	QosId           *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	QosName         *string `json:"QosName,omitempty" xml:"QosName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SagCount        *string `json:"SagCount,omitempty" xml:"SagCount,omitempty"`
	SmartAGIds      *string `json:"SmartAGIds,omitempty" xml:"SmartAGIds,omitempty"`
}

func (s DescribeQosesResponseBodyQosesQos) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosesResponseBodyQosesQos) GoString() string {
	return s.String()
}

func (s *DescribeQosesResponseBodyQosesQos) GetQosDescription() *string {
	return s.QosDescription
}

func (s *DescribeQosesResponseBodyQosesQos) GetQosId() *string {
	return s.QosId
}

func (s *DescribeQosesResponseBodyQosesQos) GetQosName() *string {
	return s.QosName
}

func (s *DescribeQosesResponseBodyQosesQos) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeQosesResponseBodyQosesQos) GetSagCount() *string {
	return s.SagCount
}

func (s *DescribeQosesResponseBodyQosesQos) GetSmartAGIds() *string {
	return s.SmartAGIds
}

func (s *DescribeQosesResponseBodyQosesQos) SetQosDescription(v string) *DescribeQosesResponseBodyQosesQos {
	s.QosDescription = &v
	return s
}

func (s *DescribeQosesResponseBodyQosesQos) SetQosId(v string) *DescribeQosesResponseBodyQosesQos {
	s.QosId = &v
	return s
}

func (s *DescribeQosesResponseBodyQosesQos) SetQosName(v string) *DescribeQosesResponseBodyQosesQos {
	s.QosName = &v
	return s
}

func (s *DescribeQosesResponseBodyQosesQos) SetResourceGroupId(v string) *DescribeQosesResponseBodyQosesQos {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeQosesResponseBodyQosesQos) SetSagCount(v string) *DescribeQosesResponseBodyQosesQos {
	s.SagCount = &v
	return s
}

func (s *DescribeQosesResponseBodyQosesQos) SetSmartAGIds(v string) *DescribeQosesResponseBodyQosesQos {
	s.SmartAGIds = &v
	return s
}

func (s *DescribeQosesResponseBodyQosesQos) Validate() error {
	return dara.Validate(s)
}
