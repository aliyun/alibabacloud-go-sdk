// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemovalProAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ImageRemovalProAsyncRequest
	GetBizId() *string
	SetImageUrl(v string) *ImageRemovalProAsyncRequest
	GetImageUrl() *string
}

type ImageRemovalProAsyncRequest struct {
	// The business order ID of the caller. This ID is recorded when the request is accepted and returned as-is when you query the asynchronous task result.
	//
	// example:
	//
	// order-20260928-001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The URL of the original image. The image must be in the JPG, JPEG, PNG, BMP, or WEBP format.
	//
	// · Resolution: 256×256 to 3000×3000. File size: ≤ 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s ImageRemovalProAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRemovalProAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImageRemovalProAsyncRequest) GetBizId() *string {
	return s.BizId
}

func (s *ImageRemovalProAsyncRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRemovalProAsyncRequest) SetBizId(v string) *ImageRemovalProAsyncRequest {
	s.BizId = &v
	return s
}

func (s *ImageRemovalProAsyncRequest) SetImageUrl(v string) *ImageRemovalProAsyncRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRemovalProAsyncRequest) Validate() error {
	return dara.Validate(s)
}
