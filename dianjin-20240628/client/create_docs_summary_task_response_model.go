// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocsSummaryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDocsSummaryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDocsSummaryTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDocsSummaryTaskResponseBody) *CreateDocsSummaryTaskResponse
	GetBody() *CreateDocsSummaryTaskResponseBody
}

type CreateDocsSummaryTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocsSummaryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDocsSummaryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDocsSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDocsSummaryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDocsSummaryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDocsSummaryTaskResponse) GetBody() *CreateDocsSummaryTaskResponseBody {
	return s.Body
}

func (s *CreateDocsSummaryTaskResponse) SetHeaders(v map[string]*string) *CreateDocsSummaryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDocsSummaryTaskResponse) SetStatusCode(v int32) *CreateDocsSummaryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocsSummaryTaskResponse) SetBody(v *CreateDocsSummaryTaskResponseBody) *CreateDocsSummaryTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDocsSummaryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
