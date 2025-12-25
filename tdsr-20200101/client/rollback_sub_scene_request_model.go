// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackSubSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *RollbackSubSceneRequest
	GetId() *string
}

type RollbackSubSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// skjjskjk****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s RollbackSubSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackSubSceneRequest) GoString() string {
	return s.String()
}

func (s *RollbackSubSceneRequest) GetId() *string {
	return s.Id
}

func (s *RollbackSubSceneRequest) SetId(v string) *RollbackSubSceneRequest {
	s.Id = &v
	return s
}

func (s *RollbackSubSceneRequest) Validate() error {
	return dara.Validate(s)
}
