// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCasterSceneConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCasterSceneConfigResponseBody
	GetRequestId() *string
}

type SetCasterSceneConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122CC1A52FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCasterSceneConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCasterSceneConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetCasterSceneConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCasterSceneConfigResponseBody) SetRequestId(v string) *SetCasterSceneConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCasterSceneConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
