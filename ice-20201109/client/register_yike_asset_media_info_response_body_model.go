// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterYikeAssetMediaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *RegisterYikeAssetMediaInfoResponseBody
	GetMediaId() *string
	SetRequestId(v string) *RegisterYikeAssetMediaInfoResponseBody
	GetRequestId() *string
}

type RegisterYikeAssetMediaInfoResponseBody struct {
	// example:
	//
	// ****c469e944b5a856828dc2****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterYikeAssetMediaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterYikeAssetMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterYikeAssetMediaInfoResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *RegisterYikeAssetMediaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterYikeAssetMediaInfoResponseBody) SetMediaId(v string) *RegisterYikeAssetMediaInfoResponseBody {
	s.MediaId = &v
	return s
}

func (s *RegisterYikeAssetMediaInfoResponseBody) SetRequestId(v string) *RegisterYikeAssetMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterYikeAssetMediaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
