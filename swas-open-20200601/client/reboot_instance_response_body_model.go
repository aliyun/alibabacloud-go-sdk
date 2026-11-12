// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootInstanceResponseBody
	GetRequestId() *string
}

type RebootInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootInstanceResponseBody) SetRequestId(v string) *RebootInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
