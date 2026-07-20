// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryResponseBody) *GetQueryResponse
	GetBody() *GetQueryResponseBody
}

type GetQueryResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryResponse) GoString() string {
	return s.String()
}

func (s *GetQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryResponse) GetBody() *GetQueryResponseBody {
	return s.Body
}

func (s *GetQueryResponse) SetHeaders(v map[string]*string) *GetQueryResponse {
	s.Headers = v
	return s
}

func (s *GetQueryResponse) SetStatusCode(v int32) *GetQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryResponse) SetBody(v *GetQueryResponseBody) *GetQueryResponse {
	s.Body = v
	return s
}

func (s *GetQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
