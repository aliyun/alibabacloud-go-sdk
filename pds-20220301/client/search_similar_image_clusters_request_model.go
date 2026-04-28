// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSimilarImageClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *SearchSimilarImageClustersRequest
	GetDriveId() *string
	SetImageThumbnailProcess(v string) *SearchSimilarImageClustersRequest
	GetImageThumbnailProcess() *string
	SetLimit(v int64) *SearchSimilarImageClustersRequest
	GetLimit() *int64
	SetMarker(v string) *SearchSimilarImageClustersRequest
	GetMarker() *string
	SetOrder(v string) *SearchSimilarImageClustersRequest
	GetOrder() *string
}

type SearchSimilarImageClustersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// example:
	//
	// 50
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// YWRzX3VzZXJfcHJvZmlsZV9je1bnQh***
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
}

func (s SearchSimilarImageClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchSimilarImageClustersRequest) GoString() string {
	return s.String()
}

func (s *SearchSimilarImageClustersRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *SearchSimilarImageClustersRequest) GetImageThumbnailProcess() *string {
	return s.ImageThumbnailProcess
}

func (s *SearchSimilarImageClustersRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *SearchSimilarImageClustersRequest) GetMarker() *string {
	return s.Marker
}

func (s *SearchSimilarImageClustersRequest) GetOrder() *string {
	return s.Order
}

func (s *SearchSimilarImageClustersRequest) SetDriveId(v string) *SearchSimilarImageClustersRequest {
	s.DriveId = &v
	return s
}

func (s *SearchSimilarImageClustersRequest) SetImageThumbnailProcess(v string) *SearchSimilarImageClustersRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *SearchSimilarImageClustersRequest) SetLimit(v int64) *SearchSimilarImageClustersRequest {
	s.Limit = &v
	return s
}

func (s *SearchSimilarImageClustersRequest) SetMarker(v string) *SearchSimilarImageClustersRequest {
	s.Marker = &v
	return s
}

func (s *SearchSimilarImageClustersRequest) SetOrder(v string) *SearchSimilarImageClustersRequest {
	s.Order = &v
	return s
}

func (s *SearchSimilarImageClustersRequest) Validate() error {
	return dara.Validate(s)
}
