// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableLineagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableLineagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableLineagesResponse
	GetStatusCode() *int32
	SetBody(v *GetTableLineagesResponseBody) *GetTableLineagesResponse
	GetBody() *GetTableLineagesResponseBody
}

type GetTableLineagesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableLineagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableLineagesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineagesResponse) GoString() string {
	return s.String()
}

func (s *GetTableLineagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableLineagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableLineagesResponse) GetBody() *GetTableLineagesResponseBody {
	return s.Body
}

func (s *GetTableLineagesResponse) SetHeaders(v map[string]*string) *GetTableLineagesResponse {
	s.Headers = v
	return s
}

func (s *GetTableLineagesResponse) SetStatusCode(v int32) *GetTableLineagesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableLineagesResponse) SetBody(v *GetTableLineagesResponseBody) *GetTableLineagesResponse {
	s.Body = v
	return s
}

func (s *GetTableLineagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
