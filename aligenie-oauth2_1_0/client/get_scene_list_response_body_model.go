// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSceneListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSceneListResponseBody
	GetRequestId() *string
	SetSceneList(v []*GetSceneListResponseBodySceneList) *GetSceneListResponseBody
	GetSceneList() []*GetSceneListResponseBodySceneList
}

type GetSceneListResponseBody struct {
	// example:
	//
	// 435CF567-58DC-5761-AFA8-650772602E2D
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneList []*GetSceneListResponseBodySceneList `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Repeated"`
}

func (s GetSceneListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSceneListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSceneListResponseBody) GetSceneList() []*GetSceneListResponseBodySceneList {
	return s.SceneList
}

func (s *GetSceneListResponseBody) SetRequestId(v string) *GetSceneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSceneListResponseBody) SetSceneList(v []*GetSceneListResponseBodySceneList) *GetSceneListResponseBody {
	s.SceneList = v
	return s
}

func (s *GetSceneListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSceneListResponseBodySceneList struct {
	// example:
	//
	// 840960b85c3c48e0bd7260c1718295fd
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s GetSceneListResponseBodySceneList) String() string {
	return dara.Prettify(s)
}

func (s GetSceneListResponseBodySceneList) GoString() string {
	return s.String()
}

func (s *GetSceneListResponseBodySceneList) GetSceneId() *string {
	return s.SceneId
}

func (s *GetSceneListResponseBodySceneList) GetSceneName() *string {
	return s.SceneName
}

func (s *GetSceneListResponseBodySceneList) SetSceneId(v string) *GetSceneListResponseBodySceneList {
	s.SceneId = &v
	return s
}

func (s *GetSceneListResponseBodySceneList) SetSceneName(v string) *GetSceneListResponseBodySceneList {
	s.SceneName = &v
	return s
}

func (s *GetSceneListResponseBodySceneList) Validate() error {
	return dara.Validate(s)
}
