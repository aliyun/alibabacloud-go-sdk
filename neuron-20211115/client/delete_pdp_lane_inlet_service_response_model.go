// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpLaneInletServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePdpLaneInletServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePdpLaneInletServiceResponse
	GetStatusCode() *int32
	SetBody(v *PdpCommonResult) *DeletePdpLaneInletServiceResponse
	GetBody() *PdpCommonResult
}

type DeletePdpLaneInletServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpCommonResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePdpLaneInletServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpLaneInletServiceResponse) GoString() string {
	return s.String()
}

func (s *DeletePdpLaneInletServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePdpLaneInletServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePdpLaneInletServiceResponse) GetBody() *PdpCommonResult {
	return s.Body
}

func (s *DeletePdpLaneInletServiceResponse) SetHeaders(v map[string]*string) *DeletePdpLaneInletServiceResponse {
	s.Headers = v
	return s
}

func (s *DeletePdpLaneInletServiceResponse) SetStatusCode(v int32) *DeletePdpLaneInletServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePdpLaneInletServiceResponse) SetBody(v *PdpCommonResult) *DeletePdpLaneInletServiceResponse {
	s.Body = v
	return s
}

func (s *DeletePdpLaneInletServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
