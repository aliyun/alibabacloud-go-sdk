// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllTrafficSpecialControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAllTrafficSpecialControlResponseBody
	GetRequestId() *string
}

type DeleteAllTrafficSpecialControlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAllTrafficSpecialControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllTrafficSpecialControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAllTrafficSpecialControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAllTrafficSpecialControlResponseBody) SetRequestId(v string) *DeleteAllTrafficSpecialControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAllTrafficSpecialControlResponseBody) Validate() error {
	return dara.Validate(s)
}
