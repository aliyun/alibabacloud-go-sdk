// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExploreUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExploreUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExploreUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetExploreUrlResponseBody) *GetExploreUrlResponse
	GetBody() *GetExploreUrlResponseBody
}

type GetExploreUrlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExploreUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExploreUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExploreUrlResponse) GoString() string {
	return s.String()
}

func (s *GetExploreUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExploreUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExploreUrlResponse) GetBody() *GetExploreUrlResponseBody {
	return s.Body
}

func (s *GetExploreUrlResponse) SetHeaders(v map[string]*string) *GetExploreUrlResponse {
	s.Headers = v
	return s
}

func (s *GetExploreUrlResponse) SetStatusCode(v int32) *GetExploreUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExploreUrlResponse) SetBody(v *GetExploreUrlResponseBody) *GetExploreUrlResponse {
	s.Body = v
	return s
}

func (s *GetExploreUrlResponse) Validate() error {
	return dara.Validate(s)
}
