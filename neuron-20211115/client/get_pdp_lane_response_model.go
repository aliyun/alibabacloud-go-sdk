// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPdpLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPdpLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPdpLaneResponse
	GetStatusCode() *int32
	SetBody(v *PdpLane) *GetPdpLaneResponse
	GetBody() *PdpLane
}

type GetPdpLaneResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpLane           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPdpLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPdpLaneResponse) GoString() string {
	return s.String()
}

func (s *GetPdpLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPdpLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPdpLaneResponse) GetBody() *PdpLane {
	return s.Body
}

func (s *GetPdpLaneResponse) SetHeaders(v map[string]*string) *GetPdpLaneResponse {
	s.Headers = v
	return s
}

func (s *GetPdpLaneResponse) SetStatusCode(v int32) *GetPdpLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPdpLaneResponse) SetBody(v *PdpLane) *GetPdpLaneResponse {
	s.Body = v
	return s
}

func (s *GetPdpLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
