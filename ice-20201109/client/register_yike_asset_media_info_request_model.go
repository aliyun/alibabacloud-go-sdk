// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterYikeAssetMediaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *RegisterYikeAssetMediaInfoRequest
	GetFolderId() *string
	SetInputURL(v string) *RegisterYikeAssetMediaInfoRequest
	GetInputURL() *string
	SetMediaType(v string) *RegisterYikeAssetMediaInfoRequest
	GetMediaType() *string
	SetProductionId(v string) *RegisterYikeAssetMediaInfoRequest
	GetProductionId() *string
}

type RegisterYikeAssetMediaInfoRequest struct {
	// The ID of the folder.
	//
	// example:
	//
	// fd-ABMFfAB2bA
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The URL of the media asset. You can use the FileURL value that is returned by the CreateYikeAssetUpload operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The type of the media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The ID of the production. If this parameter is omitted, the media asset is uploaded to the default production.
	//
	// example:
	//
	// ProductionId
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
}

func (s RegisterYikeAssetMediaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterYikeAssetMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *RegisterYikeAssetMediaInfoRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *RegisterYikeAssetMediaInfoRequest) GetInputURL() *string {
	return s.InputURL
}

func (s *RegisterYikeAssetMediaInfoRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *RegisterYikeAssetMediaInfoRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *RegisterYikeAssetMediaInfoRequest) SetFolderId(v string) *RegisterYikeAssetMediaInfoRequest {
	s.FolderId = &v
	return s
}

func (s *RegisterYikeAssetMediaInfoRequest) SetInputURL(v string) *RegisterYikeAssetMediaInfoRequest {
	s.InputURL = &v
	return s
}

func (s *RegisterYikeAssetMediaInfoRequest) SetMediaType(v string) *RegisterYikeAssetMediaInfoRequest {
	s.MediaType = &v
	return s
}

func (s *RegisterYikeAssetMediaInfoRequest) SetProductionId(v string) *RegisterYikeAssetMediaInfoRequest {
	s.ProductionId = &v
	return s
}

func (s *RegisterYikeAssetMediaInfoRequest) Validate() error {
	return dara.Validate(s)
}
