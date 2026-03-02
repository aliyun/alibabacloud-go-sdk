// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPdpServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPdpServicesResponse
	GetStatusCode() *int32
	SetBody(v *PdpServicePageResult) *ListPdpServicesResponse
	GetBody() *PdpServicePageResult
}

type ListPdpServicesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServicePageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPdpServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPdpServicesResponse) GoString() string {
	return s.String()
}

func (s *ListPdpServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPdpServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPdpServicesResponse) GetBody() *PdpServicePageResult {
	return s.Body
}

func (s *ListPdpServicesResponse) SetHeaders(v map[string]*string) *ListPdpServicesResponse {
	s.Headers = v
	return s
}

func (s *ListPdpServicesResponse) SetStatusCode(v int32) *ListPdpServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPdpServicesResponse) SetBody(v *PdpServicePageResult) *ListPdpServicesResponse {
	s.Body = v
	return s
}

func (s *ListPdpServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
