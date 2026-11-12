// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSnapshotsRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteSnapshotsRequest
	GetRegionId() *string
	SetSnapshotIds(v string) *DeleteSnapshotsRequest
	GetSnapshotIds() *string
}

type DeleteSnapshotsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot IDs. The value can be a JSON array that consists of up to 100 snapshot IDs. Separate multiple snapshot IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["s-bp16oazlsold4dks****", "s-bp16oazlsold4abc****"]
	SnapshotIds *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
}

func (s DeleteSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSnapshotsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnapshotsRequest) GetSnapshotIds() *string {
	return s.SnapshotIds
}

func (s *DeleteSnapshotsRequest) SetClientToken(v string) *DeleteSnapshotsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSnapshotsRequest) SetRegionId(v string) *DeleteSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotsRequest) SetSnapshotIds(v string) *DeleteSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

func (s *DeleteSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
