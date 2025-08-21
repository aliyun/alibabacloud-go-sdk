// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateStreamSnapshotRequest
	GetId() *string
	SetLocation(v string) *CreateStreamSnapshotRequest
	GetLocation() *string
	SetOwnerId(v int64) *CreateStreamSnapshotRequest
	GetOwnerId() *int64
}

type CreateStreamSnapshotRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323*****997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// device
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateStreamSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotRequest) GetId() *string {
	return s.Id
}

func (s *CreateStreamSnapshotRequest) GetLocation() *string {
	return s.Location
}

func (s *CreateStreamSnapshotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateStreamSnapshotRequest) SetId(v string) *CreateStreamSnapshotRequest {
	s.Id = &v
	return s
}

func (s *CreateStreamSnapshotRequest) SetLocation(v string) *CreateStreamSnapshotRequest {
	s.Location = &v
	return s
}

func (s *CreateStreamSnapshotRequest) SetOwnerId(v int64) *CreateStreamSnapshotRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateStreamSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
