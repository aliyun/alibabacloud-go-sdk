// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsRdsInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeDrdsRdsInstancesRequest
	GetDrdsInstanceId() *string
}

type DescribeDrdsRdsInstancesRequest struct {
	// The ID of the PolarDB-X instance.
	//
	// > You can call the [DescribeDrdsInstances](https://help.aliyun.com/document_detail/139284.html) operation to query the information about instances in the specified account, such as the IDs of the instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsRdsInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsRdsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsRdsInstancesRequest) SetDrdsInstanceId(v string) *DescribeDrdsRdsInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesRequest) Validate() error {
	return dara.Validate(s)
}
