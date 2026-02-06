// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmbodiedAIPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEmbodiedAIPlatformResponseBody
	GetRequestId() *string
}

type DeleteEmbodiedAIPlatformResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEmbodiedAIPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmbodiedAIPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEmbodiedAIPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEmbodiedAIPlatformResponseBody) SetRequestId(v string) *DeleteEmbodiedAIPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEmbodiedAIPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
