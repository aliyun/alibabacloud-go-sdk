// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCasterSceneConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CopyCasterSceneConfigResponseBody
	GetRequestId() *string
}

type CopyCasterSceneConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122CC1A5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyCasterSceneConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyCasterSceneConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CopyCasterSceneConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyCasterSceneConfigResponseBody) SetRequestId(v string) *CopyCasterSceneConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyCasterSceneConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
