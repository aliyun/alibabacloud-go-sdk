// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPdpImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPdpImageResponse
	GetStatusCode() *int32
	SetBody(v *PdpImagePageResult) *ListPdpImageResponse
	GetBody() *PdpImagePageResult
}

type ListPdpImageResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpImagePageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPdpImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPdpImageResponse) GoString() string {
	return s.String()
}

func (s *ListPdpImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPdpImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPdpImageResponse) GetBody() *PdpImagePageResult {
	return s.Body
}

func (s *ListPdpImageResponse) SetHeaders(v map[string]*string) *ListPdpImageResponse {
	s.Headers = v
	return s
}

func (s *ListPdpImageResponse) SetStatusCode(v int32) *ListPdpImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPdpImageResponse) SetBody(v *PdpImagePageResult) *ListPdpImageResponse {
	s.Body = v
	return s
}

func (s *ListPdpImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
