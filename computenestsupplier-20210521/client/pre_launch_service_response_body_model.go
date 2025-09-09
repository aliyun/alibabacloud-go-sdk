// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreLaunchServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PreLaunchServiceResponseBody
	GetRequestId() *string
}

type PreLaunchServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreLaunchServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreLaunchServiceResponseBody) GoString() string {
	return s.String()
}

func (s *PreLaunchServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreLaunchServiceResponseBody) SetRequestId(v string) *PreLaunchServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreLaunchServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
