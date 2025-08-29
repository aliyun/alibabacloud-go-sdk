// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureConsistencyCheckJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetFeatureConsistencyCheckJobRequest
	GetInstanceId() *string
}

type GetFeatureConsistencyCheckJobRequest struct {
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetFeatureConsistencyCheckJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobRequest) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetFeatureConsistencyCheckJobRequest) SetInstanceId(v string) *GetFeatureConsistencyCheckJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobRequest) Validate() error {
	return dara.Validate(s)
}
