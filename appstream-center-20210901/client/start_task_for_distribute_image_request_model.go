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
	DestinationRegionList []*string `json:"DestinationRegionList,omitempty" xml:"DestinationRegionList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// imgc-07jyldnd9i*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// RETRY
	RetryType *string `json:"RetryType,omitempty" xml:"RetryType,omitempty"`
	// example:
	//
	// cn-shanghai
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
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
