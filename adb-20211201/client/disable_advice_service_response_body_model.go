// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAdviceServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableAdviceServiceResponseBody
	GetRequestId() *string
}

type DisableAdviceServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 57EC6CCA-A582-572C-A33D-F61845CBC03C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAdviceServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAdviceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAdviceServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAdviceServiceResponseBody) SetRequestId(v string) *DisableAdviceServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAdviceServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
