// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecursionZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddRecursionZoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRecursionZoneResponseBody
	GetSuccess() *bool
	SetZoneId(v string) *AddRecursionZoneResponseBody
	GetZoneId() *string
}

type AddRecursionZoneResponseBody struct {
	// example:
	//
	// 51899B6F-04A5-5B96-977D-340673091ACA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Zone ID。
	//
	// example:
	//
	// 173671468000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddRecursionZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRecursionZoneResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecursionZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRecursionZoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRecursionZoneResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddRecursionZoneResponseBody) SetRequestId(v string) *AddRecursionZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRecursionZoneResponseBody) SetSuccess(v bool) *AddRecursionZoneResponseBody {
	s.Success = &v
	return s
}

func (s *AddRecursionZoneResponseBody) SetZoneId(v string) *AddRecursionZoneResponseBody {
	s.ZoneId = &v
	return s
}

func (s *AddRecursionZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
