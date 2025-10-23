// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReductionProposalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReductionProposalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReductionProposalResponse
	GetStatusCode() *int32
	SetBody(v *GetReductionProposalResponseBody) *GetReductionProposalResponse
	GetBody() *GetReductionProposalResponseBody
}

type GetReductionProposalResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReductionProposalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReductionProposalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReductionProposalResponse) GoString() string {
	return s.String()
}

func (s *GetReductionProposalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReductionProposalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReductionProposalResponse) GetBody() *GetReductionProposalResponseBody {
	return s.Body
}

func (s *GetReductionProposalResponse) SetHeaders(v map[string]*string) *GetReductionProposalResponse {
	s.Headers = v
	return s
}

func (s *GetReductionProposalResponse) SetStatusCode(v int32) *GetReductionProposalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReductionProposalResponse) SetBody(v *GetReductionProposalResponseBody) *GetReductionProposalResponse {
	s.Body = v
	return s
}

func (s *GetReductionProposalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
