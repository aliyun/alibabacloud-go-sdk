// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePdpLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePdpLaneResponse
	GetStatusCode() *int32
	SetBody(v *PdpLane) *CreatePdpLaneResponse
	GetBody() *PdpLane
}

type CreatePdpLaneResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpLane           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpLaneResponse) GoString() string {
	return s.String()
}

func (s *CreatePdpLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePdpLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePdpLaneResponse) GetBody() *PdpLane {
	return s.Body
}

func (s *CreatePdpLaneResponse) SetHeaders(v map[string]*string) *CreatePdpLaneResponse {
	s.Headers = v
	return s
}

func (s *CreatePdpLaneResponse) SetStatusCode(v int32) *CreatePdpLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePdpLaneResponse) SetBody(v *PdpLane) *CreatePdpLaneResponse {
	s.Body = v
	return s
}

func (s *CreatePdpLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
