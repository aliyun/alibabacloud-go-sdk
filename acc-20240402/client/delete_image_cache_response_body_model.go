// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteImageCacheResponseBody
	GetRequestId() *string
}

type DeleteImageCacheResponseBody struct {
	// example:
	//
	// 0E234675-3465-4CC3-9D0F-9A864BC*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageCacheResponseBody) SetRequestId(v string) *DeleteImageCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
