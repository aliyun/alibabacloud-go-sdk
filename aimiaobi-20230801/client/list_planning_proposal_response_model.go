// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlanningProposalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPlanningProposalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPlanningProposalResponse
	GetStatusCode() *int32
	SetBody(v *ListPlanningProposalResponseBody) *ListPlanningProposalResponse
	GetBody() *ListPlanningProposalResponseBody
}

type ListPlanningProposalResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPlanningProposalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPlanningProposalResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPlanningProposalResponse) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPlanningProposalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPlanningProposalResponse) GetBody() *ListPlanningProposalResponseBody {
	return s.Body
}

func (s *ListPlanningProposalResponse) SetHeaders(v map[string]*string) *ListPlanningProposalResponse {
	s.Headers = v
	return s
}

func (s *ListPlanningProposalResponse) SetStatusCode(v int32) *ListPlanningProposalResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlanningProposalResponse) SetBody(v *ListPlanningProposalResponseBody) *ListPlanningProposalResponse {
	s.Body = v
	return s
}

func (s *ListPlanningProposalResponse) Validate() error {
	return dara.Validate(s)
}
