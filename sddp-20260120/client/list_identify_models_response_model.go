// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentifyModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIdentifyModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIdentifyModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListIdentifyModelsResponseBody) *ListIdentifyModelsResponse
	GetBody() *ListIdentifyModelsResponseBody
}

type ListIdentifyModelsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdentifyModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdentifyModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIdentifyModelsResponse) GoString() string {
	return s.String()
}

func (s *ListIdentifyModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIdentifyModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIdentifyModelsResponse) GetBody() *ListIdentifyModelsResponseBody {
	return s.Body
}

func (s *ListIdentifyModelsResponse) SetHeaders(v map[string]*string) *ListIdentifyModelsResponse {
	s.Headers = v
	return s
}

func (s *ListIdentifyModelsResponse) SetStatusCode(v int32) *ListIdentifyModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdentifyModelsResponse) SetBody(v *ListIdentifyModelsResponseBody) *ListIdentifyModelsResponse {
	s.Body = v
	return s
}

func (s *ListIdentifyModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
