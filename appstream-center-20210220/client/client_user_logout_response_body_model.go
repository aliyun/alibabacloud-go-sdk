// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientUserLogoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClientUserLogoutResponseBody
	GetRequestId() *string
}

type ClientUserLogoutResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClientUserLogoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClientUserLogoutResponseBody) GoString() string {
	return s.String()
}

func (s *ClientUserLogoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClientUserLogoutResponseBody) SetRequestId(v string) *ClientUserLogoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClientUserLogoutResponseBody) Validate() error {
	return dara.Validate(s)
}
