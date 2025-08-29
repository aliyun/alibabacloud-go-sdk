// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLaboratoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLaboratoryRequest
	GetInstanceId() *string
}

type GetLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetLaboratoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *GetLaboratoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLaboratoryRequest) SetInstanceId(v string) *GetLaboratoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLaboratoryRequest) Validate() error {
	return dara.Validate(s)
}
