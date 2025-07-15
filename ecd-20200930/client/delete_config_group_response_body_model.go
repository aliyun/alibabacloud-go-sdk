// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConfigGroupResponseBody
	GetRequestId() *string
}

type DeleteConfigGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F7E4322D-D679-5ACB-A909-490D2F0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConfigGroupResponseBody) SetRequestId(v string) *DeleteConfigGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
