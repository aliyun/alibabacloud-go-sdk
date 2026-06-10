// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackAppCodeSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v string) *RollbackAppCodeSnapshotRequest
	GetSiteId() *string
	SetTargetLogicalNumber(v int32) *RollbackAppCodeSnapshotRequest
	GetTargetLogicalNumber() *int32
}

type RollbackAppCodeSnapshotRequest struct {
	// site ID
	//
	// example:
	//
	// 1067072706415168
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Target snapshot version number
	//
	// example:
	//
	// 1231
	TargetLogicalNumber *int32 `json:"TargetLogicalNumber,omitempty" xml:"TargetLogicalNumber,omitempty"`
}

func (s RollbackAppCodeSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppCodeSnapshotRequest) GoString() string {
	return s.String()
}

func (s *RollbackAppCodeSnapshotRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *RollbackAppCodeSnapshotRequest) GetTargetLogicalNumber() *int32 {
	return s.TargetLogicalNumber
}

func (s *RollbackAppCodeSnapshotRequest) SetSiteId(v string) *RollbackAppCodeSnapshotRequest {
	s.SiteId = &v
	return s
}

func (s *RollbackAppCodeSnapshotRequest) SetTargetLogicalNumber(v int32) *RollbackAppCodeSnapshotRequest {
	s.TargetLogicalNumber = &v
	return s
}

func (s *RollbackAppCodeSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
