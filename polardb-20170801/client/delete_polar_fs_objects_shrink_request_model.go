// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsObjectsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeletePolarFsObjectsShrinkRequest
	GetDBClusterId() *string
	SetObjectsToDeleteShrink(v string) *DeletePolarFsObjectsShrinkRequest
	GetObjectsToDeleteShrink() *string
	SetPolarFsInstanceId(v string) *DeletePolarFsObjectsShrinkRequest
	GetPolarFsInstanceId() *string
}

type DeletePolarFsObjectsShrinkRequest struct {
	// example:
	//
	// pc-***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	ObjectsToDeleteShrink *string `json:"ObjectsToDelete,omitempty" xml:"ObjectsToDelete,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i7*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s DeletePolarFsObjectsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsObjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarFsObjectsShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeletePolarFsObjectsShrinkRequest) GetObjectsToDeleteShrink() *string {
	return s.ObjectsToDeleteShrink
}

func (s *DeletePolarFsObjectsShrinkRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DeletePolarFsObjectsShrinkRequest) SetDBClusterId(v string) *DeletePolarFsObjectsShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeletePolarFsObjectsShrinkRequest) SetObjectsToDeleteShrink(v string) *DeletePolarFsObjectsShrinkRequest {
	s.ObjectsToDeleteShrink = &v
	return s
}

func (s *DeletePolarFsObjectsShrinkRequest) SetPolarFsInstanceId(v string) *DeletePolarFsObjectsShrinkRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DeletePolarFsObjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
