// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnData(v string) *UpdateConnDataRequest
	GetConnData() *string
	SetSceneId(v string) *UpdateConnDataRequest
	GetSceneId() *string
}

type UpdateConnDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {}
	ConnData *string `json:"ConnData,omitempty" xml:"ConnData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s UpdateConnDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnDataRequest) GetConnData() *string {
	return s.ConnData
}

func (s *UpdateConnDataRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateConnDataRequest) SetConnData(v string) *UpdateConnDataRequest {
	s.ConnData = &v
	return s
}

func (s *UpdateConnDataRequest) SetSceneId(v string) *UpdateConnDataRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateConnDataRequest) Validate() error {
	return dara.Validate(s)
}
