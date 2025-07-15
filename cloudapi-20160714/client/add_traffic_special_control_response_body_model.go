// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTrafficSpecialControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTrafficSpecialControlResponseBody
	GetRequestId() *string
}

type AddTrafficSpecialControlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTrafficSpecialControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTrafficSpecialControlResponseBody) GoString() string {
	return s.String()
}

func (s *AddTrafficSpecialControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTrafficSpecialControlResponseBody) SetRequestId(v string) *AddTrafficSpecialControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTrafficSpecialControlResponseBody) Validate() error {
	return dara.Validate(s)
}
