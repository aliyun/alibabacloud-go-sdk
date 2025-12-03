// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessLogsDownloadAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogsDownloadAttributes(v *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) *DescribeAccessLogsDownloadAttributeResponseBody
	GetLogsDownloadAttributes() *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes
	SetPageNumber(v int32) *DescribeAccessLogsDownloadAttributeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessLogsDownloadAttributeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccessLogsDownloadAttributeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAccessLogsDownloadAttributeResponseBody
	GetTotalCount() *int32
}

type DescribeAccessLogsDownloadAttributeResponseBody struct {
	// The configuration of the access log.
	LogsDownloadAttributes *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes `json:"LogsDownloadAttributes,omitempty" xml:"LogsDownloadAttributes,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8B9DB03B-ED39-5DB8-9C9F-1ED5F548D61E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessLogsDownloadAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) GetLogsDownloadAttributes() *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes {
	return s.LogsDownloadAttributes
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetLogsDownloadAttributes(v *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.LogsDownloadAttributes = v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetPageNumber(v int32) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetPageSize(v int32) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetRequestId(v string) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetTotalCount(v int32) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) Validate() error {
	if s.LogsDownloadAttributes != nil {
		if err := s.LogsDownloadAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes struct {
	LogsDownloadAttribute []*DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute `json:"LogsDownloadAttribute,omitempty" xml:"LogsDownloadAttribute,omitempty" type:"Repeated"`
}

func (s DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) GetLogsDownloadAttribute() []*DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	return s.LogsDownloadAttribute
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) SetLogsDownloadAttribute(v []*DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes {
	s.LogsDownloadAttribute = v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) Validate() error {
	if s.LogsDownloadAttribute != nil {
		for _, item := range s.LogsDownloadAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute struct {
	// The CLB instance ID.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the Log Service project.
	//
	// example:
	//
	// test-log-project
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// test-log-store
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The type of access log. Only **layer7*	- is returned, which indicates Layer 7 access logs.
	//
	// example:
	//
	// layer7
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The region ID of the CLB instance.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) GetLogProject() *string {
	return s.LogProject
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) GetLogType() *string {
	return s.LogType
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) GetRegion() *string {
	return s.Region
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetLoadBalancerId(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetLogProject(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.LogProject = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetLogStore(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.LogStore = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetLogType(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.LogType = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetRegion(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.Region = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) Validate() error {
	return dara.Validate(s)
}
