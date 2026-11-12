// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSnapshotAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateSnapshotAttributeRequest
	GetClientToken() *string
	SetRegionId(v string) *UpdateSnapshotAttributeRequest
	GetRegionId() *string
	SetRemark(v string) *UpdateSnapshotAttributeRequest
	GetRemark() *string
	SetSnapshotId(v string) *UpdateSnapshotAttributeRequest
	GetSnapshotId() *string
}

type UpdateSnapshotAttributeRequest struct {
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
	// The remarks of the snapshot of the simple application server.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The snapshot ID. You can call the ListSnapshots operation to query the snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp16oazlsold4dks****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s UpdateSnapshotAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSnapshotAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSnapshotAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSnapshotAttributeRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateSnapshotAttributeRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *UpdateSnapshotAttributeRequest) SetClientToken(v string) *UpdateSnapshotAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSnapshotAttributeRequest) SetRegionId(v string) *UpdateSnapshotAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSnapshotAttributeRequest) SetRemark(v string) *UpdateSnapshotAttributeRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSnapshotAttributeRequest) SetSnapshotId(v string) *UpdateSnapshotAttributeRequest {
	s.SnapshotId = &v
	return s
}

func (s *UpdateSnapshotAttributeRequest) Validate() error {
	return dara.Validate(s)
}
