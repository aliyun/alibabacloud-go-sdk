// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpstreamPbcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUpstreamPbcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUpstreamPbcResponse
	GetStatusCode() *int32
	SetBody(v *PbcUpDownstreamResult) *ListUpstreamPbcResponse
	GetBody() *PbcUpDownstreamResult
}

type ListUpstreamPbcResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcUpDownstreamResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUpstreamPbcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamPbcResponse) GoString() string {
	return s.String()
}

func (s *ListUpstreamPbcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUpstreamPbcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUpstreamPbcResponse) GetBody() *PbcUpDownstreamResult {
	return s.Body
}

func (s *ListUpstreamPbcResponse) SetHeaders(v map[string]*string) *ListUpstreamPbcResponse {
	s.Headers = v
	return s
}

func (s *ListUpstreamPbcResponse) SetStatusCode(v int32) *ListUpstreamPbcResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUpstreamPbcResponse) SetBody(v *PbcUpDownstreamResult) *ListUpstreamPbcResponse {
	s.Body = v
	return s
}

func (s *ListUpstreamPbcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
