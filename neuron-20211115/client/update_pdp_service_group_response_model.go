// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePdpServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePdpServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *PdpServiceGroup) *UpdatePdpServiceGroupResponse
	GetBody() *PdpServiceGroup
}

type UpdatePdpServiceGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServiceGroup   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdatePdpServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePdpServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePdpServiceGroupResponse) GetBody() *PdpServiceGroup {
	return s.Body
}

func (s *UpdatePdpServiceGroupResponse) SetHeaders(v map[string]*string) *UpdatePdpServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdatePdpServiceGroupResponse) SetStatusCode(v int32) *UpdatePdpServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePdpServiceGroupResponse) SetBody(v *PdpServiceGroup) *UpdatePdpServiceGroupResponse {
	s.Body = v
	return s
}

func (s *UpdatePdpServiceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
