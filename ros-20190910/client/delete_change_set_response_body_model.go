// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChangeSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteChangeSetResponseBody
	GetRequestId() *string
}

type DeleteChangeSetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChangeSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChangeSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChangeSetResponseBody) SetRequestId(v string) *DeleteChangeSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChangeSetResponseBody) Validate() error {
	return dara.Validate(s)
}
