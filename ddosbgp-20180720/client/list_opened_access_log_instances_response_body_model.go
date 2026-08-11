// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenedAccessLogInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListOpenedAccessLogInstancesResponseBody
	GetRequestId() *string
	SetSlsConfigStatus(v []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) *ListOpenedAccessLogInstancesResponseBody
	GetSlsConfigStatus() []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus
	SetTotalCount(v int32) *ListOpenedAccessLogInstancesResponseBody
	GetTotalCount() *int32
}

type ListOpenedAccessLogInstancesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4DB64811-70A1-41C9-A0CE-CD8B260ED551
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The log analysis configurations of Anti-DDoS Origin instances.
	SlsConfigStatus []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus `json:"SlsConfigStatus,omitempty" xml:"SlsConfigStatus,omitempty" type:"Repeated"`
	// The number of results returned, which is the number of Anti-DDoS Origin instances that have log analysis enabled.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOpenedAccessLogInstancesResponseBody) GetSlsConfigStatus() []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus {
	return s.SlsConfigStatus
}

func (s *ListOpenedAccessLogInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetRequestId(v string) *ListOpenedAccessLogInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetSlsConfigStatus(v []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) *ListOpenedAccessLogInstancesResponseBody {
	s.SlsConfigStatus = v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetTotalCount(v int32) *ListOpenedAccessLogInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBody) Validate() error {
	if s.SlsConfigStatus != nil {
		for _, item := range s.SlsConfigStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOpenedAccessLogInstancesResponseBodySlsConfigStatus struct {
	// Indicates whether log analysis is enabled for the instance. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The instance ID of Anti-DDoS Origin.
	//
	// example:
	//
	// ddosbgp-cn-m7r1zce2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) String() string {
	return dara.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) GetEnable() *bool {
	return s.Enable
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) SetEnable(v bool) *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus {
	s.Enable = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) SetInstanceId(v string) *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus {
	s.InstanceId = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) Validate() error {
	return dara.Validate(s)
}
