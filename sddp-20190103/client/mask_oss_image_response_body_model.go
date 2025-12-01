// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMaskOssImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MaskOssImageResponseBody
	GetRequestId() *string
}

type MaskOssImageResponseBody struct {
	// example:
	//
	// 136082B3-B21F-5E9D-B68E-991FFD205D24
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MaskOssImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MaskOssImageResponseBody) GoString() string {
	return s.String()
}

func (s *MaskOssImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MaskOssImageResponseBody) SetRequestId(v string) *MaskOssImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *MaskOssImageResponseBody) Validate() error {
	return dara.Validate(s)
}
