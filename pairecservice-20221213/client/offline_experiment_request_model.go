// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OfflineExperimentRequest
	GetInstanceId() *string
}

type OfflineExperimentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OfflineExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineExperimentRequest) GoString() string {
	return s.String()
}

func (s *OfflineExperimentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OfflineExperimentRequest) SetInstanceId(v string) *OfflineExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *OfflineExperimentRequest) Validate() error {
	return dara.Validate(s)
}
