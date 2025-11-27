// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedStoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomizedStoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomizedStoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomizedStoryResponseBody) *CreateCustomizedStoryResponse
	GetBody() *CreateCustomizedStoryResponseBody
}

type CreateCustomizedStoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomizedStoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomizedStoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedStoryResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomizedStoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomizedStoryResponse) GetBody() *CreateCustomizedStoryResponseBody {
	return s.Body
}

func (s *CreateCustomizedStoryResponse) SetHeaders(v map[string]*string) *CreateCustomizedStoryResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomizedStoryResponse) SetStatusCode(v int32) *CreateCustomizedStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomizedStoryResponse) SetBody(v *CreateCustomizedStoryResponseBody) *CreateCustomizedStoryResponse {
	s.Body = v
	return s
}

func (s *CreateCustomizedStoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
