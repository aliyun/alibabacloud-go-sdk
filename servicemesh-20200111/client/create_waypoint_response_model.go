// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaypointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWaypointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWaypointResponse
	GetStatusCode() *int32
	SetBody(v *CreateWaypointResponseBody) *CreateWaypointResponse
	GetBody() *CreateWaypointResponseBody
}

type CreateWaypointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWaypointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWaypointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWaypointResponse) GoString() string {
	return s.String()
}

func (s *CreateWaypointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWaypointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWaypointResponse) GetBody() *CreateWaypointResponseBody {
	return s.Body
}

func (s *CreateWaypointResponse) SetHeaders(v map[string]*string) *CreateWaypointResponse {
	s.Headers = v
	return s
}

func (s *CreateWaypointResponse) SetStatusCode(v int32) *CreateWaypointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWaypointResponse) SetBody(v *CreateWaypointResponseBody) *CreateWaypointResponse {
	s.Body = v
	return s
}

func (s *CreateWaypointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
