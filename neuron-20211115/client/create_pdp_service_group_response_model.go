// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePdpServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePdpServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *PdpServiceGroup) *CreatePdpServiceGroupResponse
	GetBody() *PdpServiceGroup
}

type CreatePdpServiceGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServiceGroup   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreatePdpServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePdpServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePdpServiceGroupResponse) GetBody() *PdpServiceGroup {
	return s.Body
}

func (s *CreatePdpServiceGroupResponse) SetHeaders(v map[string]*string) *CreatePdpServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreatePdpServiceGroupResponse) SetStatusCode(v int32) *CreatePdpServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePdpServiceGroupResponse) SetBody(v *PdpServiceGroup) *CreatePdpServiceGroupResponse {
	s.Body = v
	return s
}

func (s *CreatePdpServiceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
