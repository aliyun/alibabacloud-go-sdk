// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntitiyStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEntitiyStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEntitiyStatResponse
	GetStatusCode() *int32
	SetBody(v *GetEntitiyStatResponseBody) *GetEntitiyStatResponse
	GetBody() *GetEntitiyStatResponseBody
}

type GetEntitiyStatResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEntitiyStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEntitiyStatResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEntitiyStatResponse) GoString() string {
	return s.String()
}

func (s *GetEntitiyStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEntitiyStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEntitiyStatResponse) GetBody() *GetEntitiyStatResponseBody {
	return s.Body
}

func (s *GetEntitiyStatResponse) SetHeaders(v map[string]*string) *GetEntitiyStatResponse {
	s.Headers = v
	return s
}

func (s *GetEntitiyStatResponse) SetStatusCode(v int32) *GetEntitiyStatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEntitiyStatResponse) SetBody(v *GetEntitiyStatResponseBody) *GetEntitiyStatResponse {
	s.Body = v
	return s
}

func (s *GetEntitiyStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
