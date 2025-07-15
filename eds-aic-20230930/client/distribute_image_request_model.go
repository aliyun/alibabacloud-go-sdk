// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistributeImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDistributeRegionList(v []*string) *DistributeImageRequest
	GetDistributeRegionList() []*string
	SetImageId(v string) *DistributeImageRequest
	GetImageId() *string
}

type DistributeImageRequest struct {
	// The regions to which you want to distribute an image.
	//
	// This parameter is required.
	DistributeRegionList []*string `json:"DistributeRegionList,omitempty" xml:"DistributeRegionList,omitempty" type:"Repeated"`
	// The ID of the image that you want to distribute.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DistributeImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DistributeImageRequest) GoString() string {
	return s.String()
}

func (s *DistributeImageRequest) GetDistributeRegionList() []*string {
	return s.DistributeRegionList
}

func (s *DistributeImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DistributeImageRequest) SetDistributeRegionList(v []*string) *DistributeImageRequest {
	s.DistributeRegionList = v
	return s
}

func (s *DistributeImageRequest) SetImageId(v string) *DistributeImageRequest {
	s.ImageId = &v
	return s
}

func (s *DistributeImageRequest) Validate() error {
	return dara.Validate(s)
}
