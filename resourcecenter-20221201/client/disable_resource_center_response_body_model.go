// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableResourceCenterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableResourceCenterResponseBody
	GetRequestId() *string
}

type DisableResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D170D58E-6256-5344-8F5E-922EC9ECB7EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableResourceCenterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DisableResourceCenterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableResourceCenterResponseBody) SetRequestId(v string) *DisableResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableResourceCenterResponseBody) Validate() error {
	return dara.Validate(s)
}
