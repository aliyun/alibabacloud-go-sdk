// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTrialStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetInstanceTrialStatusResponseBody
	GetRequestId() *string
	SetTrialStatus(v bool) *GetInstanceTrialStatusResponseBody
	GetTrialStatus() *bool
}

type GetInstanceTrialStatusResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	TrialStatus *bool `json:"TrialStatus,omitempty" xml:"TrialStatus,omitempty"`
}

func (s GetInstanceTrialStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrialStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceTrialStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceTrialStatusResponseBody) GetTrialStatus() *bool {
	return s.TrialStatus
}

func (s *GetInstanceTrialStatusResponseBody) SetRequestId(v string) *GetInstanceTrialStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceTrialStatusResponseBody) SetTrialStatus(v bool) *GetInstanceTrialStatusResponseBody {
	s.TrialStatus = &v
	return s
}

func (s *GetInstanceTrialStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
