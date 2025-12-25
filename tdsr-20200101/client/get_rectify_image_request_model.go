// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRectifyImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubSceneId(v string) *GetRectifyImageRequest
	GetSubSceneId() *string
}

type GetRectifyImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetRectifyImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyImageRequest) GoString() string {
	return s.String()
}

func (s *GetRectifyImageRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *GetRectifyImageRequest) SetSubSceneId(v string) *GetRectifyImageRequest {
	s.SubSceneId = &v
	return s
}

func (s *GetRectifyImageRequest) Validate() error {
	return dara.Validate(s)
}
