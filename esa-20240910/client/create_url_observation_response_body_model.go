// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUrlObservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateUrlObservationResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateUrlObservationResponseBody
	GetRequestId() *string
}

type CreateUrlObservationResponseBody struct {
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUrlObservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUrlObservationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUrlObservationResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateUrlObservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUrlObservationResponseBody) SetConfigId(v int64) *CreateUrlObservationResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateUrlObservationResponseBody) SetRequestId(v string) *CreateUrlObservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUrlObservationResponseBody) Validate() error {
	return dara.Validate(s)
}
