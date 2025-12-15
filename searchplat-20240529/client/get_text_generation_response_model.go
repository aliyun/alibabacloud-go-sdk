// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTextGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTextGenerationResponse
	GetStatusCode() *int32
	SetBody(v *GetTextGenerationResponseBody) *GetTextGenerationResponse
	GetBody() *GetTextGenerationResponseBody
}

type GetTextGenerationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTextGenerationResponse) GoString() string {
	return s.String()
}

func (s *GetTextGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTextGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTextGenerationResponse) GetBody() *GetTextGenerationResponseBody {
	return s.Body
}

func (s *GetTextGenerationResponse) SetHeaders(v map[string]*string) *GetTextGenerationResponse {
	s.Headers = v
	return s
}

func (s *GetTextGenerationResponse) SetStatusCode(v int32) *GetTextGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextGenerationResponse) SetBody(v *GetTextGenerationResponseBody) *GetTextGenerationResponse {
	s.Body = v
	return s
}

func (s *GetTextGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
