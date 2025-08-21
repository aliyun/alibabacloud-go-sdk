// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetDiskResponseBody
	GetRequestId() *string
}

type ResetDiskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetDiskResponseBody) SetRequestId(v string) *ResetDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
