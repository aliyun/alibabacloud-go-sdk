// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetOssCheckStatusRequest
	GetRegionId() *string
}

type GetOssCheckStatusRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetOssCheckStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOssCheckStatusRequest) SetRegionId(v string) *GetOssCheckStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetOssCheckStatusRequest) Validate() error {
	return dara.Validate(s)
}
