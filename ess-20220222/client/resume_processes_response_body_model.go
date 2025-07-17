// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeProcessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeProcessesResponseBody
	GetRequestId() *string
}

type ResumeProcessesResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// E38EB733-D714-4658-8A5F-0688AB68****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeProcessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeProcessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeProcessesResponseBody) SetRequestId(v string) *ResumeProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeProcessesResponseBody) Validate() error {
	return dara.Validate(s)
}
