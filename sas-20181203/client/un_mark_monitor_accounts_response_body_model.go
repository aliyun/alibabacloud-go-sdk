// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnMarkMonitorAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnMarkMonitorAccountsResponseBody
	GetRequestId() *string
}

type UnMarkMonitorAccountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3B7EF1A4-E970-5A7F-91CA-416F9881333E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnMarkMonitorAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnMarkMonitorAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *UnMarkMonitorAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnMarkMonitorAccountsResponseBody) SetRequestId(v string) *UnMarkMonitorAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnMarkMonitorAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
