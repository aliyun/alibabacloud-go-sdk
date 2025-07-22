// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeShareUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryOptimizeShareUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryOptimizeShareUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryOptimizeShareUrlResponseBody) *GetQueryOptimizeShareUrlResponse
	GetBody() *GetQueryOptimizeShareUrlResponseBody
}

type GetQueryOptimizeShareUrlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryOptimizeShareUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryOptimizeShareUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeShareUrlResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeShareUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryOptimizeShareUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryOptimizeShareUrlResponse) GetBody() *GetQueryOptimizeShareUrlResponseBody {
	return s.Body
}

func (s *GetQueryOptimizeShareUrlResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeShareUrlResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeShareUrlResponse) SetStatusCode(v int32) *GetQueryOptimizeShareUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryOptimizeShareUrlResponse) SetBody(v *GetQueryOptimizeShareUrlResponseBody) *GetQueryOptimizeShareUrlResponse {
	s.Body = v
	return s
}

func (s *GetQueryOptimizeShareUrlResponse) Validate() error {
	return dara.Validate(s)
}
