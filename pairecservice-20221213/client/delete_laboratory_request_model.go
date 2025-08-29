// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLaboratoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteLaboratoryRequest
	GetInstanceId() *string
}

type DeleteLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteLaboratoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteLaboratoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteLaboratoryRequest) SetInstanceId(v string) *DeleteLaboratoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteLaboratoryRequest) Validate() error {
	return dara.Validate(s)
}
