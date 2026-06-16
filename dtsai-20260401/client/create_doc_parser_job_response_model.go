// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocParserJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDocParserJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDocParserJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateDocParserJobResponseBody) *CreateDocParserJobResponse
	GetBody() *CreateDocParserJobResponseBody
}

type CreateDocParserJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocParserJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDocParserJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDocParserJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDocParserJobResponse) GetBody() *CreateDocParserJobResponseBody {
	return s.Body
}

func (s *CreateDocParserJobResponse) SetHeaders(v map[string]*string) *CreateDocParserJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDocParserJobResponse) SetStatusCode(v int32) *CreateDocParserJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocParserJobResponse) SetBody(v *CreateDocParserJobResponseBody) *CreateDocParserJobResponse {
	s.Body = v
	return s
}

func (s *CreateDocParserJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
