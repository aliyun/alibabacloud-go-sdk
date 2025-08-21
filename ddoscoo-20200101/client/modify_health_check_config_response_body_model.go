// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHealthCheckConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHealthCheckConfigResponseBody
	GetRequestId() *string
}

type ModifyHealthCheckConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHealthCheckConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHealthCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHealthCheckConfigResponseBody) SetRequestId(v string) *ModifyHealthCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHealthCheckConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
