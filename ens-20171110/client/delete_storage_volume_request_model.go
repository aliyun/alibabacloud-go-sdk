// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVolumeId(v string) *DeleteStorageVolumeRequest
	GetVolumeId() *string
}

type DeleteStorageVolumeRequest struct {
	// The ID of the volume.
	//
	// This parameter is required.
	//
	// example:
	//
	// sv-***
	VolumeId *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
}

func (s DeleteStorageVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageVolumeRequest) GoString() string {
	return s.String()
}

func (s *DeleteStorageVolumeRequest) GetVolumeId() *string {
	return s.VolumeId
}

func (s *DeleteStorageVolumeRequest) SetVolumeId(v string) *DeleteStorageVolumeRequest {
	s.VolumeId = &v
	return s
}

func (s *DeleteStorageVolumeRequest) Validate() error {
	return dara.Validate(s)
}
