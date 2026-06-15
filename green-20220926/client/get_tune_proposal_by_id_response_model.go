// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTuneProposalByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTuneProposalByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTuneProposalByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetTuneProposalByIdResponseBody) *GetTuneProposalByIdResponse
	GetBody() *GetTuneProposalByIdResponseBody
}

type GetTuneProposalByIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTuneProposalByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTuneProposalByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTuneProposalByIdResponse) GoString() string {
	return s.String()
}

func (s *GetTuneProposalByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTuneProposalByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTuneProposalByIdResponse) GetBody() *GetTuneProposalByIdResponseBody {
	return s.Body
}

func (s *GetTuneProposalByIdResponse) SetHeaders(v map[string]*string) *GetTuneProposalByIdResponse {
	s.Headers = v
	return s
}

func (s *GetTuneProposalByIdResponse) SetStatusCode(v int32) *GetTuneProposalByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTuneProposalByIdResponse) SetBody(v *GetTuneProposalByIdResponseBody) *GetTuneProposalByIdResponse {
	s.Body = v
	return s
}

func (s *GetTuneProposalByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
