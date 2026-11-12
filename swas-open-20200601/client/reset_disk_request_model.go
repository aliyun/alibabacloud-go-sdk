// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ResetDiskRequest
	GetClientToken() *string
	SetDiskId(v string) *ResetDiskRequest
	GetDiskId() *string
	SetRegionId(v string) *ResetDiskRequest
	GetRegionId() *string
	SetSnapshotId(v string) *ResetDiskRequest
	GetSnapshotId() *string
}

type ResetDiskRequest struct {
	// Ensures the idempotence of the request. You can generate a parameter value on your client to ensure that the value is unique across different requests. **ClientToken*	- supports only ASCII characters and can be up to 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the disk to be rolled back.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp14wq0149cpp2xy****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the simple application server to which the snapshot belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp16oazlsold4dks****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDiskRequest) GoString() string {
	return s.String()
}

func (s *ResetDiskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResetDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ResetDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetDiskRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ResetDiskRequest) SetClientToken(v string) *ResetDiskRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetDiskRequest) SetDiskId(v string) *ResetDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResetDiskRequest) SetRegionId(v string) *ResetDiskRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDiskRequest) SetSnapshotId(v string) *ResetDiskRequest {
	s.SnapshotId = &v
	return s
}

func (s *ResetDiskRequest) Validate() error {
	return dara.Validate(s)
}
