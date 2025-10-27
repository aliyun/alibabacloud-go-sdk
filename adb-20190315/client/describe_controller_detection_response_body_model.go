// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeControllerDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeControllerDetectionResponseBody
	GetDBClusterId() *string
	SetDetectionItems(v []*DescribeControllerDetectionResponseBodyDetectionItems) *DescribeControllerDetectionResponseBody
	GetDetectionItems() []*DescribeControllerDetectionResponseBodyDetectionItems
	SetRequestId(v string) *DescribeControllerDetectionResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeControllerDetectionResponseBody
	GetTotalCount() *string
}

type DescribeControllerDetectionResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// amv-xxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried detection items and detection results.
	DetectionItems []*DescribeControllerDetectionResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D9856CFD-10DC-50AF-AE29-07C30FC57B86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeControllerDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeControllerDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeControllerDetectionResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeControllerDetectionResponseBody) GetDetectionItems() []*DescribeControllerDetectionResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeControllerDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeControllerDetectionResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeControllerDetectionResponseBody) SetDBClusterId(v string) *DescribeControllerDetectionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeControllerDetectionResponseBody) SetDetectionItems(v []*DescribeControllerDetectionResponseBodyDetectionItems) *DescribeControllerDetectionResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeControllerDetectionResponseBody) SetRequestId(v string) *DescribeControllerDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeControllerDetectionResponseBody) SetTotalCount(v string) *DescribeControllerDetectionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeControllerDetectionResponseBody) Validate() error {
	if s.DetectionItems != nil {
		for _, item := range s.DetectionItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeControllerDetectionResponseBodyDetectionItems struct {
	// The information about the detection result.
	//
	// example:
	//
	// A CPU increase is detected on the access node.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the detection item.
	//
	// example:
	//
	// CPU increase detection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The severity level of the detection result. Valid values:
	//
	// 	- NORMAL
	//
	// 	- WARNING
	//
	// 	- CRITICAL
	//
	// example:
	//
	// NORMAL
	//
	// WARNING
	//
	// CRITICAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeControllerDetectionResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeControllerDetectionResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeControllerDetectionResponseBodyDetectionItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeControllerDetectionResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeControllerDetectionResponseBodyDetectionItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeControllerDetectionResponseBodyDetectionItems) SetMessage(v string) *DescribeControllerDetectionResponseBodyDetectionItems {
	s.Message = &v
	return s
}

func (s *DescribeControllerDetectionResponseBodyDetectionItems) SetName(v string) *DescribeControllerDetectionResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeControllerDetectionResponseBodyDetectionItems) SetStatus(v string) *DescribeControllerDetectionResponseBodyDetectionItems {
	s.Status = &v
	return s
}

func (s *DescribeControllerDetectionResponseBodyDetectionItems) Validate() error {
	return dara.Validate(s)
}
