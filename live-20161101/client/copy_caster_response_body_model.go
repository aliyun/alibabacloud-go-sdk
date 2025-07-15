// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCasterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *CopyCasterResponseBody
	GetCasterId() *string
	SetRequestId(v string) *CopyCasterResponseBody
	GetRequestId() *string
}

type CopyCasterResponseBody struct {
	// The ID of the new production studio.
	//
	// example:
	//
	// 1909f043-e3d3-49e9-82d6-4329ec4a****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyCasterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyCasterResponseBody) GoString() string {
	return s.String()
}

func (s *CopyCasterResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *CopyCasterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyCasterResponseBody) SetCasterId(v string) *CopyCasterResponseBody {
	s.CasterId = &v
	return s
}

func (s *CopyCasterResponseBody) SetRequestId(v string) *CopyCasterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyCasterResponseBody) Validate() error {
	return dara.Validate(s)
}
