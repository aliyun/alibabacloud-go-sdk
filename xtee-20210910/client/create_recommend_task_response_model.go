// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecommendTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecommendTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecommendTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecommendTaskResponseBody) *CreateRecommendTaskResponse
	GetBody() *CreateRecommendTaskResponseBody
}

type CreateRecommendTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecommendTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecommendTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecommendTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRecommendTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecommendTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecommendTaskResponse) GetBody() *CreateRecommendTaskResponseBody {
	return s.Body
}

func (s *CreateRecommendTaskResponse) SetHeaders(v map[string]*string) *CreateRecommendTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRecommendTaskResponse) SetStatusCode(v int32) *CreateRecommendTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecommendTaskResponse) SetBody(v *CreateRecommendTaskResponseBody) *CreateRecommendTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRecommendTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
