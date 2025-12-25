// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempPreviewStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *TempPreviewStatusRequest
	GetSceneId() *string
}

type TempPreviewStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// m+0cmndEGjg9pv/hy4jh****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s TempPreviewStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s TempPreviewStatusRequest) GoString() string {
	return s.String()
}

func (s *TempPreviewStatusRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *TempPreviewStatusRequest) SetSceneId(v string) *TempPreviewStatusRequest {
	s.SceneId = &v
	return s
}

func (s *TempPreviewStatusRequest) Validate() error {
	return dara.Validate(s)
}
