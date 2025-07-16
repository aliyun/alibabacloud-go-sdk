// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelActiveOperationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *CancelActiveOperationTasksRequest
	GetIds() *string
	SetRegionId(v string) *CancelActiveOperationTasksRequest
	GetRegionId() *string
}

type CancelActiveOperationTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CancelActiveOperationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksRequest) GetIds() *string {
	return s.Ids
}

func (s *CancelActiveOperationTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelActiveOperationTasksRequest) SetIds(v string) *CancelActiveOperationTasksRequest {
	s.Ids = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetRegionId(v string) *CancelActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) Validate() error {
	return dara.Validate(s)
}
