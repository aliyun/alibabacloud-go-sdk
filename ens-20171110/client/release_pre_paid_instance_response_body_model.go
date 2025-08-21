// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePrePaidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleasePrePaidInstanceResponseBody
	GetRequestId() *string
}

type ReleasePrePaidInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePrePaidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleasePrePaidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePrePaidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleasePrePaidInstanceResponseBody) SetRequestId(v string) *ReleasePrePaidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleasePrePaidInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
