// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAIServiceRequest
	GetDBInstanceId() *string
	SetServiceId(v string) *DescribeAIServiceRequest
	GetServiceId() *string
	SetType(v string) *DescribeAIServiceRequest
	GetType() *string
}

type DescribeAIServiceRequest struct {
	// The ID of the instance.
	//
	// > To view details of all instances in a region, including their IDs, call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The type of service. Only drama is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// drama
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAIServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAIServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeAIServiceRequest) GetType() *string {
	return s.Type
}

func (s *DescribeAIServiceRequest) SetDBInstanceId(v string) *DescribeAIServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAIServiceRequest) SetServiceId(v string) *DescribeAIServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DescribeAIServiceRequest) SetType(v string) *DescribeAIServiceRequest {
	s.Type = &v
	return s
}

func (s *DescribeAIServiceRequest) Validate() error {
	return dara.Validate(s)
}
