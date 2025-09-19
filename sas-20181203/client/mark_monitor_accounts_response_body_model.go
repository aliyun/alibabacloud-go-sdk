// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkMonitorAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MarkMonitorAccountsResponseBody
	GetRequestId() *string
}

type MarkMonitorAccountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MarkMonitorAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MarkMonitorAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *MarkMonitorAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MarkMonitorAccountsResponseBody) SetRequestId(v string) *MarkMonitorAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *MarkMonitorAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
