// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnifiedSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnifiedSearchResponse
	GetStatusCode() *int32
	SetBody(v *UnifiedSearchOutput) *UnifiedSearchResponse
	GetBody() *UnifiedSearchOutput
}

type UnifiedSearchResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnifiedSearchOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnifiedSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s UnifiedSearchResponse) GoString() string {
	return s.String()
}

func (s *UnifiedSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnifiedSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnifiedSearchResponse) GetBody() *UnifiedSearchOutput {
	return s.Body
}

func (s *UnifiedSearchResponse) SetHeaders(v map[string]*string) *UnifiedSearchResponse {
	s.Headers = v
	return s
}

func (s *UnifiedSearchResponse) SetStatusCode(v int32) *UnifiedSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *UnifiedSearchResponse) SetBody(v *UnifiedSearchOutput) *UnifiedSearchResponse {
	s.Body = v
	return s
}

func (s *UnifiedSearchResponse) Validate() error {
	return dara.Validate(s)
}
