// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureConsistencyCheckJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetFeatureConsistencyCheckJobConfigRequest
	GetInstanceId() *string
}

type GetFeatureConsistencyCheckJobConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetFeatureConsistencyCheckJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobConfigRequest) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetFeatureConsistencyCheckJobConfigRequest) SetInstanceId(v string) *GetFeatureConsistencyCheckJobConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigRequest) Validate() error {
	return dara.Validate(s)
}
