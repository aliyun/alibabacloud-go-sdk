// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCenVbrHealthCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableCenVbrHealthCheckResponseBody
	GetRequestId() *string
}

type DisableCenVbrHealthCheckResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A278B8A6-A5B8-4FDE-9F70-95F0F6A1D68A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCenVbrHealthCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCenVbrHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCenVbrHealthCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCenVbrHealthCheckResponseBody) SetRequestId(v string) *DisableCenVbrHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCenVbrHealthCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
