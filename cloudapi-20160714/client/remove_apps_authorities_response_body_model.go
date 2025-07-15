// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAppsAuthoritiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveAppsAuthoritiesResponseBody
	GetRequestId() *string
}

type RemoveAppsAuthoritiesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D6E46F10-F26C-4AA0-BB69-FE2743D9AE62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAppsAuthoritiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveAppsAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAppsAuthoritiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveAppsAuthoritiesResponseBody) SetRequestId(v string) *RemoveAppsAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAppsAuthoritiesResponseBody) Validate() error {
	return dara.Validate(s)
}
