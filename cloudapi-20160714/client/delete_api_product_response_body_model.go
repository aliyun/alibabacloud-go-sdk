// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApiProductResponseBody
	GetRequestId() *string
}

type DeleteApiProductResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiProductResponseBody) SetRequestId(v string) *DeleteApiProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiProductResponseBody) Validate() error {
	return dara.Validate(s)
}
