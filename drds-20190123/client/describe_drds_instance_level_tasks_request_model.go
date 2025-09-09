// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceLevelTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeDrdsInstanceLevelTasksRequest
	GetDrdsInstanceId() *string
}

type DescribeDrdsInstanceLevelTasksRequest struct {
	// The ID of the PolarDB-X 1.0 instance of which the unfinished tasks you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdssen12****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsInstanceLevelTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsInstanceLevelTasksRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceLevelTasksRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksRequest) Validate() error {
	return dara.Validate(s)
}
