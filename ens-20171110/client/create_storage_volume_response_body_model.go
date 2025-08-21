// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateStorageVolumeResponseBody
	GetRequestId() *string
	SetVolumeId(v []*string) *CreateStorageVolumeResponseBody
	GetVolumeId() []*string
}

type CreateStorageVolumeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7030AB96-57CF-1C68-9FEE-D60E547FD79C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of volume IDs.
	VolumeId []*string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty" type:"Repeated"`
}

func (s CreateStorageVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStorageVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStorageVolumeResponseBody) GetVolumeId() []*string {
	return s.VolumeId
}

func (s *CreateStorageVolumeResponseBody) SetRequestId(v string) *CreateStorageVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStorageVolumeResponseBody) SetVolumeId(v []*string) *CreateStorageVolumeResponseBody {
	s.VolumeId = v
	return s
}

func (s *CreateStorageVolumeResponseBody) Validate() error {
	return dara.Validate(s)
}
