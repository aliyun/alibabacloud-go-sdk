// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DetailSceneRequest
	GetId() *string
}

type DetailSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DetailSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DetailSceneRequest) GoString() string {
	return s.String()
}

func (s *DetailSceneRequest) GetId() *string {
	return s.Id
}

func (s *DetailSceneRequest) SetId(v string) *DetailSceneRequest {
	s.Id = &v
	return s
}

func (s *DetailSceneRequest) Validate() error {
	return dara.Validate(s)
}
