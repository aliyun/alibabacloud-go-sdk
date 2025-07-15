// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighDefinitionMonitorLogAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody
	GetInstanceType() *string
	SetLogProject(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody
	GetLogProject() *string
	SetLogStore(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody
	GetLogStore() *string
	SetRequestId(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody
	GetSuccess() *string
}

type DescribeHighDefinitionMonitorLogAttributeResponseBody struct {
	// The ID of the instance whose fine-grained monitoring configurations you want to query.
	//
	// example:
	//
	// eip-wz9fi6qboho9fwgx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of instance for which you want to query fine-grained monitoring. Only **EIP*	- may be returned.
	//
	// example:
	//
	// EIP
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// hdmonitor-cn-shenzhen-1658206966225390
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// hdmonitor
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHighDefinitionMonitorLogAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighDefinitionMonitorLogAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) GetLogProject() *string {
	return s.LogProject
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) SetInstanceId(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) SetInstanceType(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) SetLogProject(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody {
	s.LogProject = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) SetLogStore(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody {
	s.LogStore = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) SetRequestId(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) SetSuccess(v string) *DescribeHighDefinitionMonitorLogAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHighDefinitionMonitorLogAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
