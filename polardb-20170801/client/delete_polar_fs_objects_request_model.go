// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeletePolarFsObjectsRequest
	GetDBClusterId() *string
	SetObjectsToDelete(v []*string) *DeletePolarFsObjectsRequest
	GetObjectsToDelete() []*string
	SetPolarFsInstanceId(v string) *DeletePolarFsObjectsRequest
	GetPolarFsInstanceId() *string
}

type DeletePolarFsObjectsRequest struct {
	// example:
	//
	// pc-***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	ObjectsToDelete []*string `json:"ObjectsToDelete,omitempty" xml:"ObjectsToDelete,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i7*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s DeletePolarFsObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsObjectsRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarFsObjectsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeletePolarFsObjectsRequest) GetObjectsToDelete() []*string {
	return s.ObjectsToDelete
}

func (s *DeletePolarFsObjectsRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DeletePolarFsObjectsRequest) SetDBClusterId(v string) *DeletePolarFsObjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeletePolarFsObjectsRequest) SetObjectsToDelete(v []*string) *DeletePolarFsObjectsRequest {
	s.ObjectsToDelete = v
	return s
}

func (s *DeletePolarFsObjectsRequest) SetPolarFsInstanceId(v string) *DeletePolarFsObjectsRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DeletePolarFsObjectsRequest) Validate() error {
	return dara.Validate(s)
}
