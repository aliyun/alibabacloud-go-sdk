// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetSnapshotRequest
	GetId() *string
}

type GetSnapshotRequest struct {
	// The unique ID of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8b723a9e8bd443af920b77e39aeb4f63
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *GetSnapshotRequest) GetId() *string {
	return s.Id
}

func (s *GetSnapshotRequest) SetId(v string) *GetSnapshotRequest {
	s.Id = &v
	return s
}

func (s *GetSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
