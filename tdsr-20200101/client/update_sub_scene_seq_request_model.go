// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubSceneSeqRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *UpdateSubSceneSeqRequest
	GetSceneId() *string
	SetSortSubSceneIds(v []*string) *UpdateSubSceneSeqRequest
	GetSortSubSceneIds() []*string
}

type UpdateSubSceneSeqRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sgyuyewyew****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	SortSubSceneIds []*string `json:"SortSubSceneIds,omitempty" xml:"SortSubSceneIds,omitempty" type:"Repeated"`
}

func (s UpdateSubSceneSeqRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneSeqRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneSeqRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateSubSceneSeqRequest) GetSortSubSceneIds() []*string {
	return s.SortSubSceneIds
}

func (s *UpdateSubSceneSeqRequest) SetSceneId(v string) *UpdateSubSceneSeqRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateSubSceneSeqRequest) SetSortSubSceneIds(v []*string) *UpdateSubSceneSeqRequest {
	s.SortSubSceneIds = v
	return s
}

func (s *UpdateSubSceneSeqRequest) Validate() error {
	return dara.Validate(s)
}
