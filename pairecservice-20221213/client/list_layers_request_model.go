// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListLayersRequest
	GetInstanceId() *string
	SetLaboratoryId(v string) *ListLayersRequest
	GetLaboratoryId() *string
}

type ListLayersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
}

func (s ListLayersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLayersRequest) GoString() string {
	return s.String()
}

func (s *ListLayersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLayersRequest) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *ListLayersRequest) SetInstanceId(v string) *ListLayersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLayersRequest) SetLaboratoryId(v string) *ListLayersRequest {
	s.LaboratoryId = &v
	return s
}

func (s *ListLayersRequest) Validate() error {
	return dara.Validate(s)
}
