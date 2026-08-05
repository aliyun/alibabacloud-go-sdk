// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarFsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolarFsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolarFsResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolarFsResponseBody) *CreatePolarFsResponse
	GetBody() *CreatePolarFsResponseBody
}

type CreatePolarFsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolarFsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolarFsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarFsResponse) GoString() string {
	return s.String()
}

func (s *CreatePolarFsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolarFsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolarFsResponse) GetBody() *CreatePolarFsResponseBody {
	return s.Body
}

func (s *CreatePolarFsResponse) SetHeaders(v map[string]*string) *CreatePolarFsResponse {
	s.Headers = v
	return s
}

func (s *CreatePolarFsResponse) SetStatusCode(v int32) *CreatePolarFsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolarFsResponse) SetBody(v *CreatePolarFsResponseBody) *CreatePolarFsResponse {
	s.Body = v
	return s
}

func (s *CreatePolarFsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
