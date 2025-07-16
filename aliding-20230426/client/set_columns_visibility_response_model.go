// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetColumnsVisibilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetColumnsVisibilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetColumnsVisibilityResponse
	GetStatusCode() *int32
	SetBody(v *SetColumnsVisibilityResponseBody) *SetColumnsVisibilityResponse
	GetBody() *SetColumnsVisibilityResponseBody
}

type SetColumnsVisibilityResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetColumnsVisibilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetColumnsVisibilityResponse) String() string {
	return dara.Prettify(s)
}

func (s SetColumnsVisibilityResponse) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetColumnsVisibilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetColumnsVisibilityResponse) GetBody() *SetColumnsVisibilityResponseBody {
	return s.Body
}

func (s *SetColumnsVisibilityResponse) SetHeaders(v map[string]*string) *SetColumnsVisibilityResponse {
	s.Headers = v
	return s
}

func (s *SetColumnsVisibilityResponse) SetStatusCode(v int32) *SetColumnsVisibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *SetColumnsVisibilityResponse) SetBody(v *SetColumnsVisibilityResponseBody) *SetColumnsVisibilityResponse {
	s.Body = v
	return s
}

func (s *SetColumnsVisibilityResponse) Validate() error {
	return dara.Validate(s)
}
