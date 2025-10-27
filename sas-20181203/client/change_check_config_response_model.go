// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeCheckConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeCheckConfigResponse
	GetStatusCode() *int32
	SetBody(v *ChangeCheckConfigResponseBody) *ChangeCheckConfigResponse
	GetBody() *ChangeCheckConfigResponseBody
}

type ChangeCheckConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCheckConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeCheckConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeCheckConfigResponse) GetBody() *ChangeCheckConfigResponseBody {
	return s.Body
}

func (s *ChangeCheckConfigResponse) SetHeaders(v map[string]*string) *ChangeCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *ChangeCheckConfigResponse) SetStatusCode(v int32) *ChangeCheckConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCheckConfigResponse) SetBody(v *ChangeCheckConfigResponseBody) *ChangeCheckConfigResponse {
	s.Body = v
	return s
}

func (s *ChangeCheckConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
