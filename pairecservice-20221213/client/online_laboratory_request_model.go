// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineLaboratoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnlineLaboratoryRequest
	GetInstanceId() *string
}

type OnlineLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnlineLaboratoryRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *OnlineLaboratoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnlineLaboratoryRequest) SetInstanceId(v string) *OnlineLaboratoryRequest {
	s.InstanceId = &v
	return s
}

func (s *OnlineLaboratoryRequest) Validate() error {
	return dara.Validate(s)
}
