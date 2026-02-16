// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockEmbodiedAIPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LockEmbodiedAIPlatformResponseBody
	GetRequestId() *string
}

type LockEmbodiedAIPlatformResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 19E994DC-A816-56DB-9F90-5F8E403DD19D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockEmbodiedAIPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LockEmbodiedAIPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *LockEmbodiedAIPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LockEmbodiedAIPlatformResponseBody) SetRequestId(v string) *LockEmbodiedAIPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockEmbodiedAIPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
