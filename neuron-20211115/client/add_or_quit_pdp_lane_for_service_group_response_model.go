// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrQuitPdpLaneForServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddOrQuitPdpLaneForServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddOrQuitPdpLaneForServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *PdpCommonResult) *AddOrQuitPdpLaneForServiceGroupResponse
	GetBody() *PdpCommonResult
}

type AddOrQuitPdpLaneForServiceGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpCommonResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrQuitPdpLaneForServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddOrQuitPdpLaneForServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *AddOrQuitPdpLaneForServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddOrQuitPdpLaneForServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrQuitPdpLaneForServiceGroupResponse) GetBody() *PdpCommonResult {
	return s.Body
}

func (s *AddOrQuitPdpLaneForServiceGroupResponse) SetHeaders(v map[string]*string) *AddOrQuitPdpLaneForServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *AddOrQuitPdpLaneForServiceGroupResponse) SetStatusCode(v int32) *AddOrQuitPdpLaneForServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrQuitPdpLaneForServiceGroupResponse) SetBody(v *PdpCommonResult) *AddOrQuitPdpLaneForServiceGroupResponse {
	s.Body = v
	return s
}

func (s *AddOrQuitPdpLaneForServiceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
