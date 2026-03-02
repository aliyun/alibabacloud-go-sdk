// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpLanesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPdpLanesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPdpLanesResponse
	GetStatusCode() *int32
	SetBody(v *PdpLanesPageResult) *ListPdpLanesResponse
	GetBody() *PdpLanesPageResult
}

type ListPdpLanesResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpLanesPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPdpLanesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPdpLanesResponse) GoString() string {
	return s.String()
}

func (s *ListPdpLanesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPdpLanesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPdpLanesResponse) GetBody() *PdpLanesPageResult {
	return s.Body
}

func (s *ListPdpLanesResponse) SetHeaders(v map[string]*string) *ListPdpLanesResponse {
	s.Headers = v
	return s
}

func (s *ListPdpLanesResponse) SetStatusCode(v int32) *ListPdpLanesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPdpLanesResponse) SetBody(v *PdpLanesPageResult) *ListPdpLanesResponse {
	s.Body = v
	return s
}

func (s *ListPdpLanesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
