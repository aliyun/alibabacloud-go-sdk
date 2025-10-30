// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTaskForDistributeImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionList(v []*string) *StartTaskForDistributeImageRequest
	GetDestinationRegionList() []*string
	SetImageId(v string) *StartTaskForDistributeImageRequest
	GetImageId() *string
	SetProductType(v string) *StartTaskForDistributeImageRequest
	GetProductType() *string
	SetRetryType(v string) *StartTaskForDistributeImageRequest
	GetRetryType() *string
	SetSourceRegion(v string) *StartTaskForDistributeImageRequest
	GetSourceRegion() *string
	SetVersionId(v string) *StartTaskForDistributeImageRequest
	GetVersionId() *string
}

type StartTaskForDistributeImageRequest struct {
	// The regions to which you want to replicate the image.
	DestinationRegionList []*string `json:"DestinationRegionList,omitempty" xml:"DestinationRegionList,omitempty" type:"Repeated"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-07jyldnd9i*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The product type.
	//
	// Valid values:
	//
	// 	- CloudDesktop: Elastic Desktop Service
	//
	// 	- CloudApp: App Streaming
	//
	// 	- WuyingServer: Workstation
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// RETRY
	RetryType *string `json:"RetryType,omitempty" xml:"RetryType,omitempty"`
	// The region where the source image is located. If you leave this parameter empty, a random region is selected.
	//
	// example:
	//
	// cn-shanghai
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The ID of the image version. If you do not specify this parameter, the latest image version is used by default.
	//
	// example:
	//
	// iv-07jyldnd9i****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s StartTaskForDistributeImageRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTaskForDistributeImageRequest) GoString() string {
	return s.String()
}

func (s *StartTaskForDistributeImageRequest) GetDestinationRegionList() []*string {
	return s.DestinationRegionList
}

func (s *StartTaskForDistributeImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *StartTaskForDistributeImageRequest) GetProductType() *string {
	return s.ProductType
}

func (s *StartTaskForDistributeImageRequest) GetRetryType() *string {
	return s.RetryType
}

func (s *StartTaskForDistributeImageRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *StartTaskForDistributeImageRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *StartTaskForDistributeImageRequest) SetDestinationRegionList(v []*string) *StartTaskForDistributeImageRequest {
	s.DestinationRegionList = v
	return s
}

func (s *StartTaskForDistributeImageRequest) SetImageId(v string) *StartTaskForDistributeImageRequest {
	s.ImageId = &v
	return s
}

func (s *StartTaskForDistributeImageRequest) SetProductType(v string) *StartTaskForDistributeImageRequest {
	s.ProductType = &v
	return s
}

func (s *StartTaskForDistributeImageRequest) SetRetryType(v string) *StartTaskForDistributeImageRequest {
	s.RetryType = &v
	return s
}

func (s *StartTaskForDistributeImageRequest) SetSourceRegion(v string) *StartTaskForDistributeImageRequest {
	s.SourceRegion = &v
	return s
}

func (s *StartTaskForDistributeImageRequest) SetVersionId(v string) *StartTaskForDistributeImageRequest {
	s.VersionId = &v
	return s
}

func (s *StartTaskForDistributeImageRequest) Validate() error {
	return dara.Validate(s)
}
