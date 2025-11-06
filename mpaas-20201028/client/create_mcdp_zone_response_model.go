// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcdpZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcdpZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcdpZoneResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcdpZoneResponseBody) *CreateMcdpZoneResponse
	GetBody() *CreateMcdpZoneResponseBody
}

type CreateMcdpZoneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcdpZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcdpZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpZoneResponse) GoString() string {
	return s.String()
}

func (s *CreateMcdpZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcdpZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcdpZoneResponse) GetBody() *CreateMcdpZoneResponseBody {
	return s.Body
}

func (s *CreateMcdpZoneResponse) SetHeaders(v map[string]*string) *CreateMcdpZoneResponse {
	s.Headers = v
	return s
}

func (s *CreateMcdpZoneResponse) SetStatusCode(v int32) *CreateMcdpZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcdpZoneResponse) SetBody(v *CreateMcdpZoneResponseBody) *CreateMcdpZoneResponse {
	s.Body = v
	return s
}

func (s *CreateMcdpZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
