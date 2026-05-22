// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUrlObservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUrlObservationResponseBody
	GetRequestId() *string
}

type UpdateUrlObservationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUrlObservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUrlObservationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUrlObservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUrlObservationResponseBody) SetRequestId(v string) *UpdateUrlObservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUrlObservationResponseBody) Validate() error {
	return dara.Validate(s)
}
