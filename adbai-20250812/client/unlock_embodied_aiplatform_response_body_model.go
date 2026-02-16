// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockEmbodiedAIPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnlockEmbodiedAIPlatformResponseBody
	GetRequestId() *string
}

type UnlockEmbodiedAIPlatformResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockEmbodiedAIPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockEmbodiedAIPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockEmbodiedAIPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockEmbodiedAIPlatformResponseBody) SetRequestId(v string) *UnlockEmbodiedAIPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockEmbodiedAIPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
