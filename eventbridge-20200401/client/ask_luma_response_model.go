// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAskLumaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AskLumaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AskLumaResponse
	GetStatusCode() *int32
	SetBody(v *AskLumaResponseBody) *AskLumaResponse
	GetBody() *AskLumaResponseBody
}

type AskLumaResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AskLumaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AskLumaResponse) String() string {
	return dara.Prettify(s)
}

func (s AskLumaResponse) GoString() string {
	return s.String()
}

func (s *AskLumaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AskLumaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AskLumaResponse) GetBody() *AskLumaResponseBody {
	return s.Body
}

func (s *AskLumaResponse) SetHeaders(v map[string]*string) *AskLumaResponse {
	s.Headers = v
	return s
}

func (s *AskLumaResponse) SetStatusCode(v int32) *AskLumaResponse {
	s.StatusCode = &v
	return s
}

func (s *AskLumaResponse) SetBody(v *AskLumaResponseBody) *AskLumaResponse {
	s.Body = v
	return s
}

func (s *AskLumaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
