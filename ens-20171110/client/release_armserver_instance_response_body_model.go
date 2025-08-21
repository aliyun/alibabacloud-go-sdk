// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseARMServerInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseARMServerInstanceResponseBody
	GetRequestId() *string
}

type ReleaseARMServerInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DCAE84DF-4187-5CC5-B819-37BCD2B83BD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseARMServerInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseARMServerInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseARMServerInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseARMServerInstanceResponseBody) SetRequestId(v string) *ReleaseARMServerInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseARMServerInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
