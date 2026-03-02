// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPdpServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPdpServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPdpServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *PdpServiceGroupDetail) *GetPdpServiceGroupResponse
	GetBody() *PdpServiceGroupDetail
}

type GetPdpServiceGroupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServiceGroupDetail `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPdpServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPdpServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetPdpServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPdpServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPdpServiceGroupResponse) GetBody() *PdpServiceGroupDetail {
	return s.Body
}

func (s *GetPdpServiceGroupResponse) SetHeaders(v map[string]*string) *GetPdpServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetPdpServiceGroupResponse) SetStatusCode(v int32) *GetPdpServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPdpServiceGroupResponse) SetBody(v *PdpServiceGroupDetail) *GetPdpServiceGroupResponse {
	s.Body = v
	return s
}

func (s *GetPdpServiceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
