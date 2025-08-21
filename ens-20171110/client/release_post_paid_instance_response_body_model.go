// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePostPaidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleasePostPaidInstanceResponseBody
	GetRequestId() *string
}

type ReleasePostPaidInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePostPaidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleasePostPaidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePostPaidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleasePostPaidInstanceResponseBody) SetRequestId(v string) *ReleasePostPaidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleasePostPaidInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
