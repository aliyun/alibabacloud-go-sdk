// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeInstanceMonitorDataResponseBody
	GetCode() *int32
	SetMonitorData(v *DescribeInstanceMonitorDataResponseBodyMonitorData) *DescribeInstanceMonitorDataResponseBody
	GetMonitorData() *DescribeInstanceMonitorDataResponseBodyMonitorData
	SetRequestId(v string) *DescribeInstanceMonitorDataResponseBody
	GetRequestId() *string
}

type DescribeInstanceMonitorDataResponseBody struct {
	// The returned service code. A value of 0 indicates that the operation was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The set of InstanceMonitorDataType data.
	MonitorData *DescribeInstanceMonitorDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C8B26B44-0189-443E-9816-D951F59623A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstanceMonitorDataResponseBody) GetMonitorData() *DescribeInstanceMonitorDataResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeInstanceMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceMonitorDataResponseBody) SetCode(v int32) *DescribeInstanceMonitorDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBody) SetMonitorData(v *DescribeInstanceMonitorDataResponseBodyMonitorData) *DescribeInstanceMonitorDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBody) SetRequestId(v string) *DescribeInstanceMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBody) Validate() error {
	if s.MonitorData != nil {
		if err := s.MonitorData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceMonitorDataResponseBodyMonitorData struct {
	InstanceMonitorData []*DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData `json:"InstanceMonitorData,omitempty" xml:"InstanceMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorData) GetInstanceMonitorData() []*DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	return s.InstanceMonitorData
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorData) SetInstanceMonitorData(v []*DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) *DescribeInstanceMonitorDataResponseBodyMonitorData {
	s.InstanceMonitorData = v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorData) Validate() error {
	if s.InstanceMonitorData != nil {
		for _, item := range s.InstanceMonitorData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData struct {
	// The vCPU usage of the instance, which is raw data. For example, a value of 0.02 indicates that the usage is 2%.
	//
	// example:
	//
	// 0.02
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// yourInstance ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is not yet supported.
	//
	// example:
	//
	// Not currently supported
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetCPU() *string {
	return s.CPU
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetMemory() *string {
	return s.Memory
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetCPU(v string) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.CPU = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetInstanceId(v string) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetMemory(v string) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.Memory = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) Validate() error {
	return dara.Validate(s)
}
