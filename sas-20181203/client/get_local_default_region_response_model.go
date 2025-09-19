// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLocalDefaultRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLocalDefaultRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLocalDefaultRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetLocalDefaultRegionResponseBody) *GetLocalDefaultRegionResponse
	GetBody() *GetLocalDefaultRegionResponseBody
}

type GetLocalDefaultRegionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLocalDefaultRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLocalDefaultRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLocalDefaultRegionResponse) GoString() string {
	return s.String()
}

func (s *GetLocalDefaultRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLocalDefaultRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLocalDefaultRegionResponse) GetBody() *GetLocalDefaultRegionResponseBody {
	return s.Body
}

func (s *GetLocalDefaultRegionResponse) SetHeaders(v map[string]*string) *GetLocalDefaultRegionResponse {
	s.Headers = v
	return s
}

func (s *GetLocalDefaultRegionResponse) SetStatusCode(v int32) *GetLocalDefaultRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLocalDefaultRegionResponse) SetBody(v *GetLocalDefaultRegionResponseBody) *GetLocalDefaultRegionResponse {
	s.Body = v
	return s
}

func (s *GetLocalDefaultRegionResponse) Validate() error {
	return dara.Validate(s)
}
