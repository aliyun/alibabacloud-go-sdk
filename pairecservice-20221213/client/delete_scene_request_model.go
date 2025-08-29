// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteSceneRequest
	GetInstanceId() *string
}

type DeleteSceneRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneRequest) GoString() string {
	return s.String()
}

func (s *DeleteSceneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSceneRequest) SetInstanceId(v string) *DeleteSceneRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSceneRequest) Validate() error {
	return dara.Validate(s)
}
