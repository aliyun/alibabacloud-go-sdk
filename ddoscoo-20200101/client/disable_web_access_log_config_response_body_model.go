// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWebAccessLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableWebAccessLogConfigResponseBody
	GetRequestId() *string
}

type DisableWebAccessLogConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableWebAccessLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableWebAccessLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DisableWebAccessLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableWebAccessLogConfigResponseBody) SetRequestId(v string) *DisableWebAccessLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableWebAccessLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
