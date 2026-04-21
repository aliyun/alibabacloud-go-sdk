// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetYikeAssetMediaInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaIds(v string) *BatchGetYikeAssetMediaInfosRequest
	GetMediaIds() *string
}

type BatchGetYikeAssetMediaInfosRequest struct {
	// example:
	//
	// ******b48fb04483915d4f2cd8******,******c48fb37407365d4f2cd8******
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s BatchGetYikeAssetMediaInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAssetMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAssetMediaInfosRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *BatchGetYikeAssetMediaInfosRequest) SetMediaIds(v string) *BatchGetYikeAssetMediaInfosRequest {
	s.MediaIds = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosRequest) Validate() error {
	return dara.Validate(s)
}
