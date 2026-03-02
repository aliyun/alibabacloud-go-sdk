// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePdpLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePdpLaneResponse
	GetStatusCode() *int32
	SetBody(v *PdpLane) *UpdatePdpLaneResponse
	GetBody() *PdpLane
}

type UpdatePdpLaneResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpLane           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpLaneResponse) GoString() string {
	return s.String()
}

func (s *UpdatePdpLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePdpLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePdpLaneResponse) GetBody() *PdpLane {
	return s.Body
}

func (s *UpdatePdpLaneResponse) SetHeaders(v map[string]*string) *UpdatePdpLaneResponse {
	s.Headers = v
	return s
}

func (s *UpdatePdpLaneResponse) SetStatusCode(v int32) *UpdatePdpLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePdpLaneResponse) SetBody(v *PdpLane) *UpdatePdpLaneResponse {
	s.Body = v
	return s
}

func (s *UpdatePdpLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
