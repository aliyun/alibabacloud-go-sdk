// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartHomeSceneListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFamilySceneList(v []*GetSmartHomeSceneListResponseBodyFamilySceneList) *GetSmartHomeSceneListResponseBody
	GetFamilySceneList() []*GetSmartHomeSceneListResponseBodyFamilySceneList
	SetRequestId(v string) *GetSmartHomeSceneListResponseBody
	GetRequestId() *string
}

type GetSmartHomeSceneListResponseBody struct {
	FamilySceneList []*GetSmartHomeSceneListResponseBodyFamilySceneList `json:"FamilySceneList,omitempty" xml:"FamilySceneList,omitempty" type:"Repeated"`
	// example:
	//
	// 435CF567-12DC-5761-AFA8-650774502E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSmartHomeSceneListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHomeSceneListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmartHomeSceneListResponseBody) GetFamilySceneList() []*GetSmartHomeSceneListResponseBodyFamilySceneList {
	return s.FamilySceneList
}

func (s *GetSmartHomeSceneListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmartHomeSceneListResponseBody) SetFamilySceneList(v []*GetSmartHomeSceneListResponseBodyFamilySceneList) *GetSmartHomeSceneListResponseBody {
	s.FamilySceneList = v
	return s
}

func (s *GetSmartHomeSceneListResponseBody) SetRequestId(v string) *GetSmartHomeSceneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmartHomeSceneListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSmartHomeSceneListResponseBodyFamilySceneList struct {
	// example:
	//
	// 2iS1AH5eo8qrw1PYBL/Ulq==
	FamilyId   *string                                                      `json:"familyId,omitempty" xml:"familyId,omitempty"`
	FamilyName *string                                                      `json:"familyName,omitempty" xml:"familyName,omitempty"`
	SceneList  []*GetSmartHomeSceneListResponseBodyFamilySceneListSceneList `json:"sceneList,omitempty" xml:"sceneList,omitempty" type:"Repeated"`
}

func (s GetSmartHomeSceneListResponseBodyFamilySceneList) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHomeSceneListResponseBodyFamilySceneList) GoString() string {
	return s.String()
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneList) GetFamilyId() *string {
	return s.FamilyId
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneList) GetFamilyName() *string {
	return s.FamilyName
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneList) GetSceneList() []*GetSmartHomeSceneListResponseBodyFamilySceneListSceneList {
	return s.SceneList
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneList) SetFamilyId(v string) *GetSmartHomeSceneListResponseBodyFamilySceneList {
	s.FamilyId = &v
	return s
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneList) SetFamilyName(v string) *GetSmartHomeSceneListResponseBodyFamilySceneList {
	s.FamilyName = &v
	return s
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneList) SetSceneList(v []*GetSmartHomeSceneListResponseBodyFamilySceneListSceneList) *GetSmartHomeSceneListResponseBodyFamilySceneList {
	s.SceneList = v
	return s
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneList) Validate() error {
	return dara.Validate(s)
}

type GetSmartHomeSceneListResponseBodyFamilySceneListSceneList struct {
	// example:
	//
	// 6813AH5586qrw1P5ln/123==
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s GetSmartHomeSceneListResponseBodyFamilySceneListSceneList) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHomeSceneListResponseBodyFamilySceneListSceneList) GoString() string {
	return s.String()
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneListSceneList) GetSceneId() *string {
	return s.SceneId
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneListSceneList) GetSceneName() *string {
	return s.SceneName
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneListSceneList) SetSceneId(v string) *GetSmartHomeSceneListResponseBodyFamilySceneListSceneList {
	s.SceneId = &v
	return s
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneListSceneList) SetSceneName(v string) *GetSmartHomeSceneListResponseBodyFamilySceneListSceneList {
	s.SceneName = &v
	return s
}

func (s *GetSmartHomeSceneListResponseBodyFamilySceneListSceneList) Validate() error {
	return dara.Validate(s)
}
