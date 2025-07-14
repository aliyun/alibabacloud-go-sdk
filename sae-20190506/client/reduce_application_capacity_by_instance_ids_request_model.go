// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReduceApplicationCapacityByInstanceIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ReduceApplicationCapacityByInstanceIdsRequest
	GetAppId() *string
	SetInstanceIds(v string) *ReduceApplicationCapacityByInstanceIdsRequest
	GetInstanceIds() *string
}

type ReduceApplicationCapacityByInstanceIdsRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance. Separate multiple instances with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// b2a8a925-477a-4ed7-b825-d5e22500****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s ReduceApplicationCapacityByInstanceIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReduceApplicationCapacityByInstanceIdsRequest) GoString() string {
	return s.String()
}

func (s *ReduceApplicationCapacityByInstanceIdsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ReduceApplicationCapacityByInstanceIdsRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ReduceApplicationCapacityByInstanceIdsRequest) SetAppId(v string) *ReduceApplicationCapacityByInstanceIdsRequest {
	s.AppId = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsRequest) SetInstanceIds(v string) *ReduceApplicationCapacityByInstanceIdsRequest {
	s.InstanceIds = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsRequest) Validate() error {
	return dara.Validate(s)
}
