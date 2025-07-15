// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLogConfigResponseBody
	GetRequestId() *string
}

type DeleteLogConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogConfigResponseBody) SetRequestId(v string) *DeleteLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
