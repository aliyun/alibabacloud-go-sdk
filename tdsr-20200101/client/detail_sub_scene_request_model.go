// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailSubSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DetailSubSceneRequest
	GetId() *string
}

type DetailSubSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// F79siXLsJsVVca8Yp4fgPA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DetailSubSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DetailSubSceneRequest) GoString() string {
	return s.String()
}

func (s *DetailSubSceneRequest) GetId() *string {
	return s.Id
}

func (s *DetailSubSceneRequest) SetId(v string) *DetailSubSceneRequest {
	s.Id = &v
	return s
}

func (s *DetailSubSceneRequest) Validate() error {
	return dara.Validate(s)
}
