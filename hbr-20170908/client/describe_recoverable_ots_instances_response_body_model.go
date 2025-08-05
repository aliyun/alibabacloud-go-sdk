// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecoverableOtsInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRecoverableOtsInstancesResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeRecoverableOtsInstancesResponseBody
	GetMessage() *string
	SetOtsInstances(v []*DescribeRecoverableOtsInstancesResponseBodyOtsInstances) *DescribeRecoverableOtsInstancesResponseBody
	GetOtsInstances() []*DescribeRecoverableOtsInstancesResponseBodyOtsInstances
	SetRequestId(v string) *DescribeRecoverableOtsInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRecoverableOtsInstancesResponseBody
	GetSuccess() *bool
}

type DescribeRecoverableOtsInstancesResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of Tablestore instances that can be restored and the tables in the instances.
	OtsInstances []*DescribeRecoverableOtsInstancesResponseBodyOtsInstances `json:"OtsInstances,omitempty" xml:"OtsInstances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 14DC089E-5DD3-5028-AEDB-93D78E11DB2A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRecoverableOtsInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecoverableOtsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableOtsInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRecoverableOtsInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRecoverableOtsInstancesResponseBody) GetOtsInstances() []*DescribeRecoverableOtsInstancesResponseBodyOtsInstances {
	return s.OtsInstances
}

func (s *DescribeRecoverableOtsInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecoverableOtsInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetCode(v string) *DescribeRecoverableOtsInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetMessage(v string) *DescribeRecoverableOtsInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetOtsInstances(v []*DescribeRecoverableOtsInstancesResponseBodyOtsInstances) *DescribeRecoverableOtsInstancesResponseBody {
	s.OtsInstances = v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetRequestId(v string) *DescribeRecoverableOtsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetSuccess(v bool) *DescribeRecoverableOtsInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRecoverableOtsInstancesResponseBodyOtsInstances struct {
	// The name of the Tablestore instance that can be restored.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The names of the tables in the Tablestore instance.
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s DescribeRecoverableOtsInstancesResponseBodyOtsInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecoverableOtsInstancesResponseBodyOtsInstances) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableOtsInstancesResponseBodyOtsInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRecoverableOtsInstancesResponseBodyOtsInstances) GetTableNames() []*string {
	return s.TableNames
}

func (s *DescribeRecoverableOtsInstancesResponseBodyOtsInstances) SetInstanceName(v string) *DescribeRecoverableOtsInstancesResponseBodyOtsInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBodyOtsInstances) SetTableNames(v []*string) *DescribeRecoverableOtsInstancesResponseBodyOtsInstances {
	s.TableNames = v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBodyOtsInstances) Validate() error {
	return dara.Validate(s)
}
