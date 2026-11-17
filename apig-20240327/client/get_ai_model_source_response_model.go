// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiModelSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiModelSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetAiModelSourceResponseBody) *GetAiModelSourceResponse
	GetBody() *GetAiModelSourceResponseBody
}

type GetAiModelSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiModelSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiModelSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelSourceResponse) GoString() string {
	return s.String()
}

func (s *GetAiModelSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiModelSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiModelSourceResponse) GetBody() *GetAiModelSourceResponseBody {
	return s.Body
}

func (s *GetAiModelSourceResponse) SetHeaders(v map[string]*string) *GetAiModelSourceResponse {
	s.Headers = v
	return s
}

func (s *GetAiModelSourceResponse) SetStatusCode(v int32) *GetAiModelSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiModelSourceResponse) SetBody(v *GetAiModelSourceResponseBody) *GetAiModelSourceResponse {
	s.Body = v
	return s
}

func (s *GetAiModelSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
