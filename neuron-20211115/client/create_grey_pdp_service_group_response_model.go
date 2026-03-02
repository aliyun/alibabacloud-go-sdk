// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGreyPdpServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGreyPdpServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGreyPdpServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *PdpServiceGroup) *CreateGreyPdpServiceGroupResponse
	GetBody() *PdpServiceGroup
}

type CreateGreyPdpServiceGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServiceGroup   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGreyPdpServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGreyPdpServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGreyPdpServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGreyPdpServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGreyPdpServiceGroupResponse) GetBody() *PdpServiceGroup {
	return s.Body
}

func (s *CreateGreyPdpServiceGroupResponse) SetHeaders(v map[string]*string) *CreateGreyPdpServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGreyPdpServiceGroupResponse) SetStatusCode(v int32) *CreateGreyPdpServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGreyPdpServiceGroupResponse) SetBody(v *PdpServiceGroup) *CreateGreyPdpServiceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateGreyPdpServiceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
