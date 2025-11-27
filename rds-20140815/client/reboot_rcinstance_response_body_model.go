// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootRCInstanceResponseBody
	GetRequestId() *string
}

type RebootRCInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3E36DB6E-AE3B-53B6-A703-85F883FD1B2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootRCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootRCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebootRCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootRCInstanceResponseBody) SetRequestId(v string) *RebootRCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootRCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
