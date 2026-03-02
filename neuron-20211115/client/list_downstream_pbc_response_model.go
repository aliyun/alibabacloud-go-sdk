// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownstreamPbcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDownstreamPbcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDownstreamPbcResponse
	GetStatusCode() *int32
	SetBody(v *PbcUpDownstreamResult) *ListDownstreamPbcResponse
	GetBody() *PbcUpDownstreamResult
}

type ListDownstreamPbcResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcUpDownstreamResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDownstreamPbcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamPbcResponse) GoString() string {
	return s.String()
}

func (s *ListDownstreamPbcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDownstreamPbcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDownstreamPbcResponse) GetBody() *PbcUpDownstreamResult {
	return s.Body
}

func (s *ListDownstreamPbcResponse) SetHeaders(v map[string]*string) *ListDownstreamPbcResponse {
	s.Headers = v
	return s
}

func (s *ListDownstreamPbcResponse) SetStatusCode(v int32) *ListDownstreamPbcResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDownstreamPbcResponse) SetBody(v *PbcUpDownstreamResult) *ListDownstreamPbcResponse {
	s.Body = v
	return s
}

func (s *ListDownstreamPbcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
