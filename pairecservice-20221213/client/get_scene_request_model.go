// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSceneRequest
	GetInstanceId() *string
}

type GetSceneRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSceneRequest) GoString() string {
	return s.String()
}

func (s *GetSceneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSceneRequest) SetInstanceId(v string) *GetSceneRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSceneRequest) Validate() error {
	return dara.Validate(s)
}
