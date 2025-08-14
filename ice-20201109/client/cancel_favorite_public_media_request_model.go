// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFavoritePublicMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaIds(v string) *CancelFavoritePublicMediaRequest
	GetMediaIds() *string
}

type CancelFavoritePublicMediaRequest struct {
	// example:
	//
	// icepublic-****7213c6050cbc66750b469701****,icepublic-****0b4697017213c6050cbc6675****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s CancelFavoritePublicMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelFavoritePublicMediaRequest) GoString() string {
	return s.String()
}

func (s *CancelFavoritePublicMediaRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *CancelFavoritePublicMediaRequest) SetMediaIds(v string) *CancelFavoritePublicMediaRequest {
	s.MediaIds = &v
	return s
}

func (s *CancelFavoritePublicMediaRequest) Validate() error {
	return dara.Validate(s)
}
