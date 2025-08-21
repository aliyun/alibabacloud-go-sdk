// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RescaleApplicationResponseBody
	GetRequestId() *string
}

type RescaleApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RescaleApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RescaleApplicationResponseBody) SetRequestId(v string) *RescaleApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RescaleApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
