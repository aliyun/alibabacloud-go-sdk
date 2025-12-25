// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSubSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DropSubSceneRequest
	GetId() *string
}

type DropSubSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DropSubSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DropSubSceneRequest) GoString() string {
	return s.String()
}

func (s *DropSubSceneRequest) GetId() *string {
	return s.Id
}

func (s *DropSubSceneRequest) SetId(v string) *DropSubSceneRequest {
	s.Id = &v
	return s
}

func (s *DropSubSceneRequest) Validate() error {
	return dara.Validate(s)
}
