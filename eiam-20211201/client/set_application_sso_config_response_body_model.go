// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationSsoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetApplicationSsoConfigResponseBody
	GetRequestId() *string
}

type SetApplicationSsoConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationSsoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationSsoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetApplicationSsoConfigResponseBody) SetRequestId(v string) *SetApplicationSsoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetApplicationSsoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
