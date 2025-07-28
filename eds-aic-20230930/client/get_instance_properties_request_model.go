// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancePropertiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstancePropertiesRequest
	GetInstanceId() *string
}

type GetInstancePropertiesRequest struct {
	// example:
	//
	// cpn-2ofr9kf41apy3****-014
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstancePropertiesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstancePropertiesRequest) GoString() string {
	return s.String()
}

func (s *GetInstancePropertiesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstancePropertiesRequest) SetInstanceId(v string) *GetInstancePropertiesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstancePropertiesRequest) Validate() error {
	return dara.Validate(s)
}
