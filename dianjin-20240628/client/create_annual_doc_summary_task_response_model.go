// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnualDocSummaryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAnnualDocSummaryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAnnualDocSummaryTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAnnualDocSummaryTaskResponseBody) *CreateAnnualDocSummaryTaskResponse
	GetBody() *CreateAnnualDocSummaryTaskResponseBody
}

type CreateAnnualDocSummaryTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnnualDocSummaryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnnualDocSummaryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnualDocSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAnnualDocSummaryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAnnualDocSummaryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAnnualDocSummaryTaskResponse) GetBody() *CreateAnnualDocSummaryTaskResponseBody {
	return s.Body
}

func (s *CreateAnnualDocSummaryTaskResponse) SetHeaders(v map[string]*string) *CreateAnnualDocSummaryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponse) SetStatusCode(v int32) *CreateAnnualDocSummaryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponse) SetBody(v *CreateAnnualDocSummaryTaskResponseBody) *CreateAnnualDocSummaryTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
