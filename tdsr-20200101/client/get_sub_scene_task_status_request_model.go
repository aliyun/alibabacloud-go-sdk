// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubSceneTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubSceneId(v string) *GetSubSceneTaskStatusRequest
	GetSubSceneId() *string
}

type GetSubSceneTaskStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetSubSceneTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubSceneTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *GetSubSceneTaskStatusRequest) SetSubSceneId(v string) *GetSubSceneTaskStatusRequest {
	s.SubSceneId = &v
	return s
}

func (s *GetSubSceneTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
