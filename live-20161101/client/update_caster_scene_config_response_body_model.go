// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCasterSceneConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCasterSceneConfigResponseBody
	GetRequestId() *string
}

type UpdateCasterSceneConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122CC1A52FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCasterSceneConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterSceneConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCasterSceneConfigResponseBody) SetRequestId(v string) *UpdateCasterSceneConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCasterSceneConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
