// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSwimLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSwimLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSwimLaneResponse
	GetStatusCode() *int32
	SetBody(v *CreateSwimLaneResponseBody) *CreateSwimLaneResponse
	GetBody() *CreateSwimLaneResponseBody
}

type CreateSwimLaneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSwimLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSwimLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSwimLaneResponse) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSwimLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSwimLaneResponse) GetBody() *CreateSwimLaneResponseBody {
	return s.Body
}

func (s *CreateSwimLaneResponse) SetHeaders(v map[string]*string) *CreateSwimLaneResponse {
	s.Headers = v
	return s
}

func (s *CreateSwimLaneResponse) SetStatusCode(v int32) *CreateSwimLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSwimLaneResponse) SetBody(v *CreateSwimLaneResponseBody) *CreateSwimLaneResponse {
	s.Body = v
	return s
}

func (s *CreateSwimLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
