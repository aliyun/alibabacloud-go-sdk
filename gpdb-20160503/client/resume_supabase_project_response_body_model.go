// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeSupabaseProjectResponseBody
	GetRequestId() *string
}

type ResumeSupabaseProjectResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeSupabaseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeSupabaseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeSupabaseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeSupabaseProjectResponseBody) SetRequestId(v string) *ResumeSupabaseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeSupabaseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
