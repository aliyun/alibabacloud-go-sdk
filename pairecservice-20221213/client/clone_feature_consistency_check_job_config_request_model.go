// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneFeatureConsistencyCheckJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CloneFeatureConsistencyCheckJobConfigRequest
	GetInstanceId() *string
}

type CloneFeatureConsistencyCheckJobConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneFeatureConsistencyCheckJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneFeatureConsistencyCheckJobConfigRequest) GoString() string {
	return s.String()
}

func (s *CloneFeatureConsistencyCheckJobConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CloneFeatureConsistencyCheckJobConfigRequest) SetInstanceId(v string) *CloneFeatureConsistencyCheckJobConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneFeatureConsistencyCheckJobConfigRequest) Validate() error {
	return dara.Validate(s)
}
