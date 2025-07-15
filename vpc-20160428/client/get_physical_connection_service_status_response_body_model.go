// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalConnectionServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *GetPhysicalConnectionServiceStatusResponseBody
	GetEnabled() *bool
	SetRequestId(v string) *GetPhysicalConnectionServiceStatusResponseBody
	GetRequestId() *string
}

type GetPhysicalConnectionServiceStatusResponseBody struct {
	// Indicates whether billing for outbound data transfer is enabled. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhysicalConnectionServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalConnectionServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalConnectionServiceStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetPhysicalConnectionServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhysicalConnectionServiceStatusResponseBody) SetEnabled(v bool) *GetPhysicalConnectionServiceStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetPhysicalConnectionServiceStatusResponseBody) SetRequestId(v string) *GetPhysicalConnectionServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalConnectionServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
