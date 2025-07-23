// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUid(v string) *CopyImageRequest
	GetImageUid() *string
	SetTargetRegionId(v string) *CopyImageRequest
	GetTargetRegionId() *string
}

type CopyImageRequest struct {
	// The ID of the image.
	//
	// example:
	//
	// image-hafiudfahdd****
	ImageUid *string `json:"ImageUid,omitempty" xml:"ImageUid,omitempty"`
	// The ID of the destination region.
	//
	// example:
	//
	// cn-beijing
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
}

func (s CopyImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyImageRequest) GoString() string {
	return s.String()
}

func (s *CopyImageRequest) GetImageUid() *string {
	return s.ImageUid
}

func (s *CopyImageRequest) GetTargetRegionId() *string {
	return s.TargetRegionId
}

func (s *CopyImageRequest) SetImageUid(v string) *CopyImageRequest {
	s.ImageUid = &v
	return s
}

func (s *CopyImageRequest) SetTargetRegionId(v string) *CopyImageRequest {
	s.TargetRegionId = &v
	return s
}

func (s *CopyImageRequest) Validate() error {
	return dara.Validate(s)
}
