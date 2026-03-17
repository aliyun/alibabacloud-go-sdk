// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHealthCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHealthCheckResponseBody
	GetRequestId() *string
}

type ModifyHealthCheckResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F0290F85-8087-4FB7-81F8-84226A4DAAB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHealthCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHealthCheckResponseBody) SetRequestId(v string) *ModifyHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHealthCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
