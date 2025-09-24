// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConfigResponseBody
	GetRequestId() *string
}

type DeleteConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A******C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConfigResponseBody) SetRequestId(v string) *DeleteConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
