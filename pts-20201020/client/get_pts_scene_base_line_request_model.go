// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneBaseLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *GetPtsSceneBaseLineRequest
	GetSceneId() *string
}

type GetPtsSceneBaseLineRequest struct {
	// The ID of the scene. For more information, see the [table](https://help.aliyun.com/document_detail/201321.html) provided in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// NB54CV
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsSceneBaseLineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneBaseLineRequest) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetPtsSceneBaseLineRequest) SetSceneId(v string) *GetPtsSceneBaseLineRequest {
	s.SceneId = &v
	return s
}

func (s *GetPtsSceneBaseLineRequest) Validate() error {
	return dara.Validate(s)
}
