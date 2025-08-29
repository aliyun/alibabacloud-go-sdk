// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineExperimentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OfflineExperimentGroupRequest
	GetInstanceId() *string
}

type OfflineExperimentGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OfflineExperimentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *OfflineExperimentGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OfflineExperimentGroupRequest) SetInstanceId(v string) *OfflineExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *OfflineExperimentGroupRequest) Validate() error {
	return dara.Validate(s)
}
