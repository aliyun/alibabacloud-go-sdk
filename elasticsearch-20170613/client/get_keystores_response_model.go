// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeystoresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKeystoresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKeystoresResponse
	GetStatusCode() *int32
	SetBody(v *GetKeystoresResponseBody) *GetKeystoresResponse
	GetBody() *GetKeystoresResponseBody
}

type GetKeystoresResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeystoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeystoresResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKeystoresResponse) GoString() string {
	return s.String()
}

func (s *GetKeystoresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKeystoresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKeystoresResponse) GetBody() *GetKeystoresResponseBody {
	return s.Body
}

func (s *GetKeystoresResponse) SetHeaders(v map[string]*string) *GetKeystoresResponse {
	s.Headers = v
	return s
}

func (s *GetKeystoresResponse) SetStatusCode(v int32) *GetKeystoresResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeystoresResponse) SetBody(v *GetKeystoresResponseBody) *GetKeystoresResponse {
	s.Body = v
	return s
}

func (s *GetKeystoresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
