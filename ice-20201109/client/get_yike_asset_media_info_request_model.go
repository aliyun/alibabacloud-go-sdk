// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAssetMediaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetYikeAssetMediaInfoRequest
	GetMediaId() *string
}

type GetYikeAssetMediaInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetYikeAssetMediaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetYikeAssetMediaInfoRequest) SetMediaId(v string) *GetYikeAssetMediaInfoRequest {
	s.MediaId = &v
	return s
}

func (s *GetYikeAssetMediaInfoRequest) Validate() error {
	return dara.Validate(s)
}
