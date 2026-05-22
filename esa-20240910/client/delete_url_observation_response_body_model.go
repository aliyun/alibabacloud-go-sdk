// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrlObservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUrlObservationResponseBody
	GetRequestId() *string
}

type DeleteUrlObservationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUrlObservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrlObservationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUrlObservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUrlObservationResponseBody) SetRequestId(v string) *DeleteUrlObservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUrlObservationResponseBody) Validate() error {
	return dara.Validate(s)
}
