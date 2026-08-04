// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindCustomerSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInfoType(v string) *FindCustomerSnapshotRequest
	GetInfoType() *string
	SetPk(v int64) *FindCustomerSnapshotRequest
	GetPk() *int64
	SetVersionId(v string) *FindCustomerSnapshotRequest
	GetVersionId() *string
}

type FindCustomerSnapshotRequest struct {
	// This parameter is required.
	InfoType *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	// This parameter is required.
	Pk *int64 `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s FindCustomerSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerSnapshotRequest) GoString() string {
	return s.String()
}

func (s *FindCustomerSnapshotRequest) GetInfoType() *string {
	return s.InfoType
}

func (s *FindCustomerSnapshotRequest) GetPk() *int64 {
	return s.Pk
}

func (s *FindCustomerSnapshotRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *FindCustomerSnapshotRequest) SetInfoType(v string) *FindCustomerSnapshotRequest {
	s.InfoType = &v
	return s
}

func (s *FindCustomerSnapshotRequest) SetPk(v int64) *FindCustomerSnapshotRequest {
	s.Pk = &v
	return s
}

func (s *FindCustomerSnapshotRequest) SetVersionId(v string) *FindCustomerSnapshotRequest {
	s.VersionId = &v
	return s
}

func (s *FindCustomerSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
