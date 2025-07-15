// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrafficControlResponseBody
	GetRequestId() *string
}

type DeleteTrafficControlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrafficControlResponseBody) SetRequestId(v string) *DeleteTrafficControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrafficControlResponseBody) Validate() error {
	return dara.Validate(s)
}
