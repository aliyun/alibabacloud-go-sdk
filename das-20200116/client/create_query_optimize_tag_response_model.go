// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueryOptimizeTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQueryOptimizeTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQueryOptimizeTagResponse
	GetStatusCode() *int32
	SetBody(v *CreateQueryOptimizeTagResponseBody) *CreateQueryOptimizeTagResponse
	GetBody() *CreateQueryOptimizeTagResponseBody
}

type CreateQueryOptimizeTagResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQueryOptimizeTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQueryOptimizeTagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryOptimizeTagResponse) GoString() string {
	return s.String()
}

func (s *CreateQueryOptimizeTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQueryOptimizeTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQueryOptimizeTagResponse) GetBody() *CreateQueryOptimizeTagResponseBody {
	return s.Body
}

func (s *CreateQueryOptimizeTagResponse) SetHeaders(v map[string]*string) *CreateQueryOptimizeTagResponse {
	s.Headers = v
	return s
}

func (s *CreateQueryOptimizeTagResponse) SetStatusCode(v int32) *CreateQueryOptimizeTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQueryOptimizeTagResponse) SetBody(v *CreateQueryOptimizeTagResponseBody) *CreateQueryOptimizeTagResponse {
	s.Body = v
	return s
}

func (s *CreateQueryOptimizeTagResponse) Validate() error {
	return dara.Validate(s)
}
