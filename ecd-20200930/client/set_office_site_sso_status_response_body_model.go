// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOfficeSiteSsoStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetOfficeSiteSsoStatusResponseBody
	GetRequestId() *string
}

type SetOfficeSiteSsoStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetOfficeSiteSsoStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetOfficeSiteSsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetOfficeSiteSsoStatusResponseBody) SetRequestId(v string) *SetOfficeSiteSsoStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetOfficeSiteSsoStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
