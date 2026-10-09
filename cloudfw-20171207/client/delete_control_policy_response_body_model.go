// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteControlPolicyResponseBody
	GetDryRun() *bool
	SetRequestId(v string) *DeleteControlPolicyResponseBody
	GetRequestId() *string
}

type DeleteControlPolicyResponseBody struct {
	// Indicates whether the response is for a successful dry run. A value of true indicates that only the precheck is completed and no actual changes are made. This field is not returned or is set to false for actual calls.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponseBody) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteControlPolicyResponseBody) SetDryRun(v bool) *DeleteControlPolicyResponseBody {
	s.DryRun = &v
	return s
}

func (s *DeleteControlPolicyResponseBody) SetRequestId(v string) *DeleteControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
