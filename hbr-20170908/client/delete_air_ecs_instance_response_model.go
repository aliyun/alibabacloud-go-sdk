// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAirEcsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAirEcsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAirEcsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAirEcsInstanceResponseBody) *DeleteAirEcsInstanceResponse
	GetBody() *DeleteAirEcsInstanceResponseBody
}

type DeleteAirEcsInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAirEcsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAirEcsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAirEcsInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteAirEcsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAirEcsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAirEcsInstanceResponse) GetBody() *DeleteAirEcsInstanceResponseBody {
	return s.Body
}

func (s *DeleteAirEcsInstanceResponse) SetHeaders(v map[string]*string) *DeleteAirEcsInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteAirEcsInstanceResponse) SetStatusCode(v int32) *DeleteAirEcsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAirEcsInstanceResponse) SetBody(v *DeleteAirEcsInstanceResponseBody) *DeleteAirEcsInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteAirEcsInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
