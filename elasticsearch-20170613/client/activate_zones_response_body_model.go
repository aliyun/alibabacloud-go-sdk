// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActivateZonesResponseBody
	GetRequestId() *string
	SetResult(v bool) *ActivateZonesResponseBody
	GetResult() *bool
}

type ActivateZonesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A5D8E74-565C-43DC-B031-29289FA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the nodes in disabled zones are restored. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ActivateZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateZonesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ActivateZonesResponseBody) SetRequestId(v string) *ActivateZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateZonesResponseBody) SetResult(v bool) *ActivateZonesResponseBody {
	s.Result = &v
	return s
}

func (s *ActivateZonesResponseBody) Validate() error {
	return dara.Validate(s)
}
