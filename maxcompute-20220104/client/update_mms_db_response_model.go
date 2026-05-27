// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsDbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmsDbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmsDbResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmsDbResponseBody) *UpdateMmsDbResponse
	GetBody() *UpdateMmsDbResponseBody
}

type UpdateMmsDbResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmsDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmsDbResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsDbResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmsDbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmsDbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmsDbResponse) GetBody() *UpdateMmsDbResponseBody {
	return s.Body
}

func (s *UpdateMmsDbResponse) SetHeaders(v map[string]*string) *UpdateMmsDbResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmsDbResponse) SetStatusCode(v int32) *UpdateMmsDbResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmsDbResponse) SetBody(v *UpdateMmsDbResponseBody) *UpdateMmsDbResponse {
	s.Body = v
	return s
}

func (s *UpdateMmsDbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
