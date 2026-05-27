// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmsTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmsTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmsTableResponseBody) *UpdateMmsTableResponse
	GetBody() *UpdateMmsTableResponseBody
}

type UpdateMmsTableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmsTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmsTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmsTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmsTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmsTableResponse) GetBody() *UpdateMmsTableResponseBody {
	return s.Body
}

func (s *UpdateMmsTableResponse) SetHeaders(v map[string]*string) *UpdateMmsTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmsTableResponse) SetStatusCode(v int32) *UpdateMmsTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmsTableResponse) SetBody(v *UpdateMmsTableResponseBody) *UpdateMmsTableResponse {
	s.Body = v
	return s
}

func (s *UpdateMmsTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
