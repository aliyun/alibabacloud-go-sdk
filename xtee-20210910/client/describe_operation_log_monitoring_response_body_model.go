// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperationLogMonitoringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOperationLogMonitoringResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeOperationLogMonitoringResponseBodyResultObject) *DescribeOperationLogMonitoringResponseBody
	GetResultObject() []*DescribeOperationLogMonitoringResponseBodyResultObject
}

type DescribeOperationLogMonitoringResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object.
	ResultObject []*DescribeOperationLogMonitoringResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeOperationLogMonitoringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperationLogMonitoringResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperationLogMonitoringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOperationLogMonitoringResponseBody) GetResultObject() []*DescribeOperationLogMonitoringResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeOperationLogMonitoringResponseBody) SetRequestId(v string) *DescribeOperationLogMonitoringResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperationLogMonitoringResponseBody) SetResultObject(v []*DescribeOperationLogMonitoringResponseBodyResultObject) *DescribeOperationLogMonitoringResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeOperationLogMonitoringResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOperationLogMonitoringResponseBodyResultObject struct {
	// Time axis slice.
	//
	// example:
	//
	// 2025-07-19
	DateGrouped *string `json:"dateGrouped,omitempty" xml:"dateGrouped,omitempty"`
	// Total count.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s DescribeOperationLogMonitoringResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperationLogMonitoringResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeOperationLogMonitoringResponseBodyResultObject) GetDateGrouped() *string {
	return s.DateGrouped
}

func (s *DescribeOperationLogMonitoringResponseBodyResultObject) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeOperationLogMonitoringResponseBodyResultObject) SetDateGrouped(v string) *DescribeOperationLogMonitoringResponseBodyResultObject {
	s.DateGrouped = &v
	return s
}

func (s *DescribeOperationLogMonitoringResponseBodyResultObject) SetTotalCount(v string) *DescribeOperationLogMonitoringResponseBodyResultObject {
	s.TotalCount = &v
	return s
}

func (s *DescribeOperationLogMonitoringResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
