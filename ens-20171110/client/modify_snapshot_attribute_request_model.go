// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifySnapshotAttributeRequest
	GetDescription() *string
	SetSnapshotId(v string) *ModifySnapshotAttributeRequest
	GetSnapshotId() *string
	SetSnapshotName(v string) *ModifySnapshotAttributeRequest
	GetSnapshotName() *string
}

type ModifySnapshotAttributeRequest struct {
	// The description of the snapshot. The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-bp199lyny9bb47pa****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The name of the snapshot. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// The name cannot start with **auto*	- because snapshots whose names start with auto are recognized as automatic snapshots.
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s ModifySnapshotAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySnapshotAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySnapshotAttributeRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ModifySnapshotAttributeRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *ModifySnapshotAttributeRequest) SetDescription(v string) *ModifySnapshotAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetSnapshotId(v string) *ModifySnapshotAttributeRequest {
	s.SnapshotId = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetSnapshotName(v string) *ModifySnapshotAttributeRequest {
	s.SnapshotName = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) Validate() error {
	return dara.Validate(s)
}
