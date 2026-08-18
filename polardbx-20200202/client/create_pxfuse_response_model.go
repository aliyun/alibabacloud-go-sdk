// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfuseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePxfuseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePxfuseResponse
	GetStatusCode() *int32
	SetBody(v *CreatePxfuseResponseBody) *CreatePxfuseResponse
	GetBody() *CreatePxfuseResponseBody
}

type CreatePxfuseResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePxfuseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePxfuseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfuseResponse) GoString() string {
	return s.String()
}

func (s *CreatePxfuseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePxfuseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePxfuseResponse) GetBody() *CreatePxfuseResponseBody {
	return s.Body
}

func (s *CreatePxfuseResponse) SetHeaders(v map[string]*string) *CreatePxfuseResponse {
	s.Headers = v
	return s
}

func (s *CreatePxfuseResponse) SetStatusCode(v int32) *CreatePxfuseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePxfuseResponse) SetBody(v *CreatePxfuseResponseBody) *CreatePxfuseResponse {
	s.Body = v
	return s
}

func (s *CreatePxfuseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
