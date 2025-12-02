// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupBucketsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetBackupBucketsListRequest
	GetRegionId() *string
}

type GetBackupBucketsListRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBackupBucketsListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBackupBucketsListRequest) GoString() string {
	return s.String()
}

func (s *GetBackupBucketsListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBackupBucketsListRequest) SetRegionId(v string) *GetBackupBucketsListRequest {
	s.RegionId = &v
	return s
}

func (s *GetBackupBucketsListRequest) Validate() error {
	return dara.Validate(s)
}
