// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootDesktopsResponseBody
	GetRequestId() *string
}

type RebootDesktopsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootDesktopsResponseBody) SetRequestId(v string) *RebootDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootDesktopsResponseBody) Validate() error {
	return dara.Validate(s)
}
