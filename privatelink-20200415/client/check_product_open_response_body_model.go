// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckProductOpenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckProductOpenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckProductOpenResponseBody
	GetSuccess() *bool
}

type CheckProductOpenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether PrivateLink is activated.
	//
	// Only **true*	- is returned. The value indicates that PrivateLink is activated.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckProductOpenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckProductOpenResponseBody) GoString() string {
	return s.String()
}

func (s *CheckProductOpenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckProductOpenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckProductOpenResponseBody) SetRequestId(v string) *CheckProductOpenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckProductOpenResponseBody) SetSuccess(v bool) *CheckProductOpenResponseBody {
	s.Success = &v
	return s
}

func (s *CheckProductOpenResponseBody) Validate() error {
	return dara.Validate(s)
}
