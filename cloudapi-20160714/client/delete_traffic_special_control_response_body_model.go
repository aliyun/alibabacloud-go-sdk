// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficSpecialControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrafficSpecialControlResponseBody
	GetRequestId() *string
}

type DeleteTrafficSpecialControlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficSpecialControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficSpecialControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficSpecialControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrafficSpecialControlResponseBody) SetRequestId(v string) *DeleteTrafficSpecialControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrafficSpecialControlResponseBody) Validate() error {
	return dara.Validate(s)
}
