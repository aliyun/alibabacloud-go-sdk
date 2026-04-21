// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteYikeAssetMediaInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogicDelete(v bool) *DeleteYikeAssetMediaInfosRequest
	GetLogicDelete() *bool
	SetMediaIds(v string) *DeleteYikeAssetMediaInfosRequest
	GetMediaIds() *string
}

type DeleteYikeAssetMediaInfosRequest struct {
	// example:
	//
	// true
	LogicDelete *bool `json:"LogicDelete,omitempty" xml:"LogicDelete,omitempty"`
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****,****15d4a4b0448391508f2cb486****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s DeleteYikeAssetMediaInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteYikeAssetMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *DeleteYikeAssetMediaInfosRequest) GetLogicDelete() *bool {
	return s.LogicDelete
}

func (s *DeleteYikeAssetMediaInfosRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *DeleteYikeAssetMediaInfosRequest) SetLogicDelete(v bool) *DeleteYikeAssetMediaInfosRequest {
	s.LogicDelete = &v
	return s
}

func (s *DeleteYikeAssetMediaInfosRequest) SetMediaIds(v string) *DeleteYikeAssetMediaInfosRequest {
	s.MediaIds = &v
	return s
}

func (s *DeleteYikeAssetMediaInfosRequest) Validate() error {
	return dara.Validate(s)
}
