// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncCatalogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SyncCatalogsRequest
	GetId() *string
	SetOwnerId(v int64) *SyncCatalogsRequest
	GetOwnerId() *int64
}

type SyncCatalogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3238****739092996
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SyncCatalogsRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncCatalogsRequest) GoString() string {
	return s.String()
}

func (s *SyncCatalogsRequest) GetId() *string {
	return s.Id
}

func (s *SyncCatalogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SyncCatalogsRequest) SetId(v string) *SyncCatalogsRequest {
	s.Id = &v
	return s
}

func (s *SyncCatalogsRequest) SetOwnerId(v int64) *SyncCatalogsRequest {
	s.OwnerId = &v
	return s
}

func (s *SyncCatalogsRequest) Validate() error {
	return dara.Validate(s)
}
