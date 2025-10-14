// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKvResponse
	GetStatusCode() *int32
	SetBody(v *GetKvResponseBody) *GetKvResponse
	GetBody() *GetKvResponseBody
}

type GetKvResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKvResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKvResponse) GoString() string {
	return s.String()
}

func (s *GetKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKvResponse) GetBody() *GetKvResponseBody {
	return s.Body
}

func (s *GetKvResponse) SetHeaders(v map[string]*string) *GetKvResponse {
	s.Headers = v
	return s
}

func (s *GetKvResponse) SetStatusCode(v int32) *GetKvResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKvResponse) SetBody(v *GetKvResponseBody) *GetKvResponse {
	s.Body = v
	return s
}

func (s *GetKvResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
