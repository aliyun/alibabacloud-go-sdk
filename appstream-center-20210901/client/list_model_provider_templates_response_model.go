// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelProviderTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelProviderTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelProviderTemplatesResponseBody) *ListModelProviderTemplatesResponse
	GetBody() *ListModelProviderTemplatesResponseBody
}

type ListModelProviderTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelProviderTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelProviderTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListModelProviderTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelProviderTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelProviderTemplatesResponse) GetBody() *ListModelProviderTemplatesResponseBody {
	return s.Body
}

func (s *ListModelProviderTemplatesResponse) SetHeaders(v map[string]*string) *ListModelProviderTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListModelProviderTemplatesResponse) SetStatusCode(v int32) *ListModelProviderTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelProviderTemplatesResponse) SetBody(v *ListModelProviderTemplatesResponseBody) *ListModelProviderTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListModelProviderTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
