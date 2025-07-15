// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApisAuthoritiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetApisAuthoritiesResponseBody
	GetRequestId() *string
}

type SetApisAuthoritiesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D6E46F10-F26C-4AA0-BB69-FE2743D9AE62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApisAuthoritiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApisAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *SetApisAuthoritiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetApisAuthoritiesResponseBody) SetRequestId(v string) *SetApisAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetApisAuthoritiesResponseBody) Validate() error {
	return dara.Validate(s)
}
