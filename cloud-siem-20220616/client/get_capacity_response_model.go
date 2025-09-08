// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCapacityResponse
	GetStatusCode() *int32
	SetBody(v *GetCapacityResponseBody) *GetCapacityResponse
	GetBody() *GetCapacityResponseBody
}

type GetCapacityResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCapacityResponse) GoString() string {
	return s.String()
}

func (s *GetCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCapacityResponse) GetBody() *GetCapacityResponseBody {
	return s.Body
}

func (s *GetCapacityResponse) SetHeaders(v map[string]*string) *GetCapacityResponse {
	s.Headers = v
	return s
}

func (s *GetCapacityResponse) SetStatusCode(v int32) *GetCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCapacityResponse) SetBody(v *GetCapacityResponseBody) *GetCapacityResponse {
	s.Body = v
	return s
}

func (s *GetCapacityResponse) Validate() error {
	return dara.Validate(s)
}
