// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePdpLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePdpLaneResponse
	GetStatusCode() *int32
	SetBody(v *PdpCommonResult) *DeletePdpLaneResponse
	GetBody() *PdpCommonResult
}

type DeletePdpLaneResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpCommonResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePdpLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpLaneResponse) GoString() string {
	return s.String()
}

func (s *DeletePdpLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePdpLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePdpLaneResponse) GetBody() *PdpCommonResult {
	return s.Body
}

func (s *DeletePdpLaneResponse) SetHeaders(v map[string]*string) *DeletePdpLaneResponse {
	s.Headers = v
	return s
}

func (s *DeletePdpLaneResponse) SetStatusCode(v int32) *DeletePdpLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePdpLaneResponse) SetBody(v *PdpCommonResult) *DeletePdpLaneResponse {
	s.Body = v
	return s
}

func (s *DeletePdpLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
