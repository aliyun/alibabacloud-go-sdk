// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsScenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneIds(v []*string) *DeletePtsScenesRequest
	GetSceneIds() []*string
}

type DeletePtsScenesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ["XVB4DF","AFG3CV"]
	SceneIds []*string `json:"SceneIds,omitempty" xml:"SceneIds,omitempty" type:"Repeated"`
}

func (s DeletePtsScenesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsScenesRequest) GoString() string {
	return s.String()
}

func (s *DeletePtsScenesRequest) GetSceneIds() []*string {
	return s.SceneIds
}

func (s *DeletePtsScenesRequest) SetSceneIds(v []*string) *DeletePtsScenesRequest {
	s.SceneIds = v
	return s
}

func (s *DeletePtsScenesRequest) Validate() error {
	return dara.Validate(s)
}
