// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHotspotFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *AddHotspotFileRequest
	GetFileName() *string
	SetSceneId(v string) *AddHotspotFileRequest
	GetSceneId() *string
	SetType(v string) *AddHotspotFileRequest
	GetType() *string
}

type AddHotspotFileRequest struct {
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddHotspotFileRequest) String() string {
	return dara.Prettify(s)
}

func (s AddHotspotFileRequest) GoString() string {
	return s.String()
}

func (s *AddHotspotFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *AddHotspotFileRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *AddHotspotFileRequest) GetType() *string {
	return s.Type
}

func (s *AddHotspotFileRequest) SetFileName(v string) *AddHotspotFileRequest {
	s.FileName = &v
	return s
}

func (s *AddHotspotFileRequest) SetSceneId(v string) *AddHotspotFileRequest {
	s.SceneId = &v
	return s
}

func (s *AddHotspotFileRequest) SetType(v string) *AddHotspotFileRequest {
	s.Type = &v
	return s
}

func (s *AddHotspotFileRequest) Validate() error {
	return dara.Validate(s)
}
