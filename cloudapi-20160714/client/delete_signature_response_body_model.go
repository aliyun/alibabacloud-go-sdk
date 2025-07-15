// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSignatureResponseBody
	GetRequestId() *string
}

type DeleteSignatureResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSignatureResponseBody) SetRequestId(v string) *DeleteSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSignatureResponseBody) Validate() error {
	return dara.Validate(s)
}
