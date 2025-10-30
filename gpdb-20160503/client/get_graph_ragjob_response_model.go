// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphRAGJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGraphRAGJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGraphRAGJobResponse
	GetStatusCode() *int32
	SetBody(v *GetGraphRAGJobResponseBody) *GetGraphRAGJobResponse
	GetBody() *GetGraphRAGJobResponseBody
}

type GetGraphRAGJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGraphRAGJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGraphRAGJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGraphRAGJobResponse) GoString() string {
	return s.String()
}

func (s *GetGraphRAGJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGraphRAGJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGraphRAGJobResponse) GetBody() *GetGraphRAGJobResponseBody {
	return s.Body
}

func (s *GetGraphRAGJobResponse) SetHeaders(v map[string]*string) *GetGraphRAGJobResponse {
	s.Headers = v
	return s
}

func (s *GetGraphRAGJobResponse) SetStatusCode(v int32) *GetGraphRAGJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGraphRAGJobResponse) SetBody(v *GetGraphRAGJobResponseBody) *GetGraphRAGJobResponse {
	s.Body = v
	return s
}

func (s *GetGraphRAGJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
