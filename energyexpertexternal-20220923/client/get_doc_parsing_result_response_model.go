// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocParsingResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocParsingResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocParsingResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDocParsingResultResponseBody) *GetDocParsingResultResponse
	GetBody() *GetDocParsingResultResponseBody
}

type GetDocParsingResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocParsingResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocParsingResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocParsingResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocParsingResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocParsingResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocParsingResultResponse) GetBody() *GetDocParsingResultResponseBody {
	return s.Body
}

func (s *GetDocParsingResultResponse) SetHeaders(v map[string]*string) *GetDocParsingResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocParsingResultResponse) SetStatusCode(v int32) *GetDocParsingResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocParsingResultResponse) SetBody(v *GetDocParsingResultResponseBody) *GetDocParsingResultResponse {
	s.Body = v
	return s
}

func (s *GetDocParsingResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
