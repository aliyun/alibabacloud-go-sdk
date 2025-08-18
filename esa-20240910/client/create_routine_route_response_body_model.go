// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateRoutineRouteResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateRoutineRouteResponseBody
	GetRequestId() *string
}

type CreateRoutineRouteResponseBody struct {
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRoutineRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoutineRouteResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateRoutineRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoutineRouteResponseBody) SetConfigId(v int64) *CreateRoutineRouteResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateRoutineRouteResponseBody) SetRequestId(v string) *CreateRoutineRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoutineRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
