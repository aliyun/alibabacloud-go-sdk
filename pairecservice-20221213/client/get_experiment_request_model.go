// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetExperimentRequest
	GetInstanceId() *string
}

type GetExperimentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetExperimentRequest) SetInstanceId(v string) *GetExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *GetExperimentRequest) Validate() error {
	return dara.Validate(s)
}
