// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelSourceModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiModelSourceModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiModelSourceModelResponse
	GetStatusCode() *int32
	SetBody(v *GetAiModelSourceModelResponseBody) *GetAiModelSourceModelResponse
	GetBody() *GetAiModelSourceModelResponseBody
}

type GetAiModelSourceModelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiModelSourceModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiModelSourceModelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelSourceModelResponse) GoString() string {
	return s.String()
}

func (s *GetAiModelSourceModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiModelSourceModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiModelSourceModelResponse) GetBody() *GetAiModelSourceModelResponseBody {
	return s.Body
}

func (s *GetAiModelSourceModelResponse) SetHeaders(v map[string]*string) *GetAiModelSourceModelResponse {
	s.Headers = v
	return s
}

func (s *GetAiModelSourceModelResponse) SetStatusCode(v int32) *GetAiModelSourceModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiModelSourceModelResponse) SetBody(v *GetAiModelSourceModelResponseBody) *GetAiModelSourceModelResponse {
	s.Body = v
	return s
}

func (s *GetAiModelSourceModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
