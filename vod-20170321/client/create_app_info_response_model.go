// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppInfoResponseBody) *CreateAppInfoResponse
	GetBody() *CreateAppInfoResponseBody
}

type CreateAppInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppInfoResponse) GetBody() *CreateAppInfoResponseBody {
	return s.Body
}

func (s *CreateAppInfoResponse) SetHeaders(v map[string]*string) *CreateAppInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateAppInfoResponse) SetStatusCode(v int32) *CreateAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppInfoResponse) SetBody(v *CreateAppInfoResponseBody) *CreateAppInfoResponse {
	s.Body = v
	return s
}

func (s *CreateAppInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
