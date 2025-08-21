// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePortResponseBody
	GetRequestId() *string
}

type DeletePortResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 39499F01-19D9-4EA4-A0E9-C6014BA5CDBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePortResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePortResponseBody) SetRequestId(v string) *DeletePortResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePortResponseBody) Validate() error {
	return dara.Validate(s)
}
