// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWebCCResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableWebCCResponseBody
	GetRequestId() *string
}

type DisableWebCCResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableWebCCResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableWebCCResponseBody) GoString() string {
	return s.String()
}

func (s *DisableWebCCResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableWebCCResponseBody) SetRequestId(v string) *DisableWebCCResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableWebCCResponseBody) Validate() error {
	return dara.Validate(s)
}
