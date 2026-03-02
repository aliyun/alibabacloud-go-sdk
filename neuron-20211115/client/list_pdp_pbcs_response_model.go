// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpPbcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPdpPbcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPdpPbcsResponse
	GetStatusCode() *int32
	SetBody(v *PdpPbcListResult) *ListPdpPbcsResponse
	GetBody() *PdpPbcListResult
}

type ListPdpPbcsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpPbcListResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPdpPbcsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPdpPbcsResponse) GoString() string {
	return s.String()
}

func (s *ListPdpPbcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPdpPbcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPdpPbcsResponse) GetBody() *PdpPbcListResult {
	return s.Body
}

func (s *ListPdpPbcsResponse) SetHeaders(v map[string]*string) *ListPdpPbcsResponse {
	s.Headers = v
	return s
}

func (s *ListPdpPbcsResponse) SetStatusCode(v int32) *ListPdpPbcsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPdpPbcsResponse) SetBody(v *PdpPbcListResult) *ListPdpPbcsResponse {
	s.Body = v
	return s
}

func (s *ListPdpPbcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
