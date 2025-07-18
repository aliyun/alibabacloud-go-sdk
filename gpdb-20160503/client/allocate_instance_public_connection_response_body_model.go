// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody
	GetRequestId() *string
}

type AllocateInstancePublicConnectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ADD6EA90-EECB-4C12-9F26-0B6DB587****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateInstancePublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateInstancePublicConnectionResponseBody) SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
