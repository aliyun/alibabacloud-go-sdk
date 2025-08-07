// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBLinkResponseBody
	GetRequestId() *string
}

type DeleteDBLinkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F9F1CB1A-B1D5-4EF5-A53A-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBLinkResponseBody) SetRequestId(v string) *DeleteDBLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
