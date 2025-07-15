// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppsAuthoritiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAppsAuthoritiesResponseBody
	GetRequestId() *string
}

type SetAppsAuthoritiesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D6E46F10-F26C-4AA0-BB69-FE2743D9AE62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAppsAuthoritiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAppsAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppsAuthoritiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAppsAuthoritiesResponseBody) SetRequestId(v string) *SetAppsAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppsAuthoritiesResponseBody) Validate() error {
	return dara.Validate(s)
}
