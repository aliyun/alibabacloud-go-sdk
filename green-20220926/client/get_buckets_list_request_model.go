// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetBucketsListRequest
	GetRegionId() *string
}

type GetBucketsListRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBucketsListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketsListRequest) GoString() string {
	return s.String()
}

func (s *GetBucketsListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBucketsListRequest) SetRegionId(v string) *GetBucketsListRequest {
	s.RegionId = &v
	return s
}

func (s *GetBucketsListRequest) Validate() error {
	return dara.Validate(s)
}
