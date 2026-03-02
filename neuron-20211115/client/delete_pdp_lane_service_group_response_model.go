// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpLaneServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePdpLaneServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePdpLaneServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *PdpCommonResult) *DeletePdpLaneServiceGroupResponse
	GetBody() *PdpCommonResult
}

type DeletePdpLaneServiceGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpCommonResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePdpLaneServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpLaneServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeletePdpLaneServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePdpLaneServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePdpLaneServiceGroupResponse) GetBody() *PdpCommonResult {
	return s.Body
}

func (s *DeletePdpLaneServiceGroupResponse) SetHeaders(v map[string]*string) *DeletePdpLaneServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeletePdpLaneServiceGroupResponse) SetStatusCode(v int32) *DeletePdpLaneServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePdpLaneServiceGroupResponse) SetBody(v *PdpCommonResult) *DeletePdpLaneServiceGroupResponse {
	s.Body = v
	return s
}

func (s *DeletePdpLaneServiceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
