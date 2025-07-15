// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficMirrorServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *GetTrafficMirrorServiceStatusResponseBody
	GetEnabled() *bool
	SetRequestId(v string) *GetTrafficMirrorServiceStatusResponseBody
	GetRequestId() *string
}

type GetTrafficMirrorServiceStatusResponseBody struct {
	// Indicates whether the traffic mirror feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 37626066-2C6C-4B62-ADD3-498920C409C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrafficMirrorServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficMirrorServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrafficMirrorServiceStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetTrafficMirrorServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrafficMirrorServiceStatusResponseBody) SetEnabled(v bool) *GetTrafficMirrorServiceStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetTrafficMirrorServiceStatusResponseBody) SetRequestId(v string) *GetTrafficMirrorServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrafficMirrorServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
