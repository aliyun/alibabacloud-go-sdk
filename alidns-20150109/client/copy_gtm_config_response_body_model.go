// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyGtmConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CopyGtmConfigResponseBody
	GetRequestId() *string
}

type CopyGtmConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 389DFFA3-77A5-4A9E-BF3D-147C6F98A5BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyGtmConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyGtmConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CopyGtmConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyGtmConfigResponseBody) SetRequestId(v string) *CopyGtmConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyGtmConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
