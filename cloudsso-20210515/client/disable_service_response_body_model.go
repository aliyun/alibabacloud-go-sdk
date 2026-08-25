// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableServiceResponseBody
	GetRequestId() *string
}

type DisableServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3257EAD2-8723-1F26-B69C-F8707D8B565D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableServiceResponseBody) SetRequestId(v string) *DisableServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
