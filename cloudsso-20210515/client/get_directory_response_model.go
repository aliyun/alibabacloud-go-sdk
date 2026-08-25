// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *GetDirectoryResponseBody) *GetDirectoryResponse
	GetBody() *GetDirectoryResponseBody
}

type GetDirectoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryResponse) GoString() string {
	return s.String()
}

func (s *GetDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDirectoryResponse) GetBody() *GetDirectoryResponseBody {
	return s.Body
}

func (s *GetDirectoryResponse) SetHeaders(v map[string]*string) *GetDirectoryResponse {
	s.Headers = v
	return s
}

func (s *GetDirectoryResponse) SetStatusCode(v int32) *GetDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectoryResponse) SetBody(v *GetDirectoryResponseBody) *GetDirectoryResponse {
	s.Body = v
	return s
}

func (s *GetDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
