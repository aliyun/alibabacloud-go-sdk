// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewRenderingInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenewRenderingInstanceResponseBody
	GetRequestId() *string
}

type RenewRenderingInstanceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewRenderingInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewRenderingInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewRenderingInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewRenderingInstanceResponseBody) SetRequestId(v string) *RenewRenderingInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewRenderingInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
