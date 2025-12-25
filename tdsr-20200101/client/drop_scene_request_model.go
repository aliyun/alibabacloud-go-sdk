// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DropSceneRequest
	GetId() *string
}

type DropSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DropSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DropSceneRequest) GoString() string {
	return s.String()
}

func (s *DropSceneRequest) GetId() *string {
	return s.Id
}

func (s *DropSceneRequest) SetId(v string) *DropSceneRequest {
	s.Id = &v
	return s
}

func (s *DropSceneRequest) Validate() error {
	return dara.Validate(s)
}
