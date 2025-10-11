// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceCountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceCountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceCountsResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceCountsResponseBody) *GetResourceCountsResponse
	GetBody() *GetResourceCountsResponseBody
}

type GetResourceCountsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceCountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceCountsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCountsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceCountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceCountsResponse) GetBody() *GetResourceCountsResponseBody {
	return s.Body
}

func (s *GetResourceCountsResponse) SetHeaders(v map[string]*string) *GetResourceCountsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceCountsResponse) SetStatusCode(v int32) *GetResourceCountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceCountsResponse) SetBody(v *GetResourceCountsResponseBody) *GetResourceCountsResponse {
	s.Body = v
	return s
}

func (s *GetResourceCountsResponse) Validate() error {
	return dara.Validate(s)
}
