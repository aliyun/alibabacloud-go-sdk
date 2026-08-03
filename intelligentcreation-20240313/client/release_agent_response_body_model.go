// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseAgentResponseBody
	GetRequestId() *string
	SetStatus(v string) *ReleaseAgentResponseBody
	GetStatus() *string
}

type ReleaseAgentResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ReleaseAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseAgentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ReleaseAgentResponseBody) SetRequestId(v string) *ReleaseAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseAgentResponseBody) SetStatus(v string) *ReleaseAgentResponseBody {
	s.Status = &v
	return s
}

func (s *ReleaseAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
