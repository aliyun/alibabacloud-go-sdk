// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEmbodiedAIPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEmbodiedAIPlatformResponseBody
	GetRequestId() *string
}

type CreateEmbodiedAIPlatformResponseBody struct {
	// example:
	//
	// 19E994DC-A816-56DB-9F90-5F8E403DD19D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEmbodiedAIPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEmbodiedAIPlatformResponseBody) SetRequestId(v string) *CreateEmbodiedAIPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEmbodiedAIPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
