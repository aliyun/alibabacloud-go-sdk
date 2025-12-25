// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempPreviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *TempPreviewRequest
	GetSceneId() *string
}

type TempPreviewRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// m+0cmndEGjg9pv/hy4jh****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s TempPreviewRequest) String() string {
	return dara.Prettify(s)
}

func (s TempPreviewRequest) GoString() string {
	return s.String()
}

func (s *TempPreviewRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *TempPreviewRequest) SetSceneId(v string) *TempPreviewRequest {
	s.SceneId = &v
	return s
}

func (s *TempPreviewRequest) Validate() error {
	return dara.Validate(s)
}
