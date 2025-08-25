// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageCacheId(v string) *CreateImageCacheResponseBody
	GetImageCacheId() *string
	SetRequestId(v string) *CreateImageCacheResponseBody
	GetRequestId() *string
}

type CreateImageCacheResponseBody struct {
	// example:
	//
	// imc-bp1dj*****
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// example:
	//
	// 0E234675-3465-4CC3-9D0F-9A864BC*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponseBody) GetImageCacheId() *string {
	return s.ImageCacheId
}

func (s *CreateImageCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageCacheResponseBody) SetImageCacheId(v string) *CreateImageCacheResponseBody {
	s.ImageCacheId = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetRequestId(v string) *CreateImageCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
