// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeactivateZonesResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeactivateZonesResponseBody
	GetResult() *bool
}

type DeactivateZonesResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result:
	//
	// - true: Zone offline successful
	//
	// - false: Zone offline failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeactivateZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactivateZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactivateZonesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeactivateZonesResponseBody) SetRequestId(v string) *DeactivateZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactivateZonesResponseBody) SetResult(v bool) *DeactivateZonesResponseBody {
	s.Result = &v
	return s
}

func (s *DeactivateZonesResponseBody) Validate() error {
	return dara.Validate(s)
}
