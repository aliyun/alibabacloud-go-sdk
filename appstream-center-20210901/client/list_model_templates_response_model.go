// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelTemplatesResponseBody) *ListModelTemplatesResponse
	GetBody() *ListModelTemplatesResponseBody
}

type ListModelTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListModelTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelTemplatesResponse) GetBody() *ListModelTemplatesResponseBody {
	return s.Body
}

func (s *ListModelTemplatesResponse) SetHeaders(v map[string]*string) *ListModelTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListModelTemplatesResponse) SetStatusCode(v int32) *ListModelTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelTemplatesResponse) SetBody(v *ListModelTemplatesResponseBody) *ListModelTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListModelTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
