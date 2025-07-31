// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuckets(v string) *GetScanNumRequest
	GetBuckets() *string
	SetMediaType(v int32) *GetScanNumRequest
	GetMediaType() *int32
	SetRegionId(v string) *GetScanNumRequest
	GetRegionId() *string
}

type GetScanNumRequest struct {
	// example:
	//
	// tmpsample
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// image
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetScanNumRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScanNumRequest) GoString() string {
	return s.String()
}

func (s *GetScanNumRequest) GetBuckets() *string {
	return s.Buckets
}

func (s *GetScanNumRequest) GetMediaType() *int32 {
	return s.MediaType
}

func (s *GetScanNumRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetScanNumRequest) SetBuckets(v string) *GetScanNumRequest {
	s.Buckets = &v
	return s
}

func (s *GetScanNumRequest) SetMediaType(v int32) *GetScanNumRequest {
	s.MediaType = &v
	return s
}

func (s *GetScanNumRequest) SetRegionId(v string) *GetScanNumRequest {
	s.RegionId = &v
	return s
}

func (s *GetScanNumRequest) Validate() error {
	return dara.Validate(s)
}
