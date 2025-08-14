// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFavoritePublicMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaIds(v string) *AddFavoritePublicMediaRequest
	GetMediaIds() *string
}

type AddFavoritePublicMediaRequest struct {
	// example:
	//
	// icepublic-****7213c6050cbc66750b469701****,icepublic-****0b4697017213c6050cbc6675****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s AddFavoritePublicMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFavoritePublicMediaRequest) GoString() string {
	return s.String()
}

func (s *AddFavoritePublicMediaRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *AddFavoritePublicMediaRequest) SetMediaIds(v string) *AddFavoritePublicMediaRequest {
	s.MediaIds = &v
	return s
}

func (s *AddFavoritePublicMediaRequest) Validate() error {
	return dara.Validate(s)
}
