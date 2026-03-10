// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoriesResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoriesResponseBody) *GetMemoriesResponse
	GetBody() *GetMemoriesResponseBody
}

type GetMemoriesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoriesResponse) GoString() string {
	return s.String()
}

func (s *GetMemoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoriesResponse) GetBody() *GetMemoriesResponseBody {
	return s.Body
}

func (s *GetMemoriesResponse) SetHeaders(v map[string]*string) *GetMemoriesResponse {
	s.Headers = v
	return s
}

func (s *GetMemoriesResponse) SetStatusCode(v int32) *GetMemoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoriesResponse) SetBody(v *GetMemoriesResponseBody) *GetMemoriesResponse {
	s.Body = v
	return s
}

func (s *GetMemoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
