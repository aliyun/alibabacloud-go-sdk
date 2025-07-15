// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApisAuthoritiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveApisAuthoritiesResponseBody
	GetRequestId() *string
}

type RemoveApisAuthoritiesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D6E46F10-F26C-4AA0-BB69-FE2743D9AE62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveApisAuthoritiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveApisAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApisAuthoritiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveApisAuthoritiesResponseBody) SetRequestId(v string) *RemoveApisAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveApisAuthoritiesResponseBody) Validate() error {
	return dara.Validate(s)
}
