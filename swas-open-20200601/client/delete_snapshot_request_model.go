// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSnapshotRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteSnapshotRequest
	GetRegionId() *string
	SetSnapshotId(v string) *DeleteSnapshotRequest
	GetSnapshotId() *string
}

type DeleteSnapshotRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp16oazlsold4dks****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DeleteSnapshotRequest) SetClientToken(v string) *DeleteSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *DeleteSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
