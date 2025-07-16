// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegionStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetRegionStatusResponseBody) *GetRegionStatusResponse
	GetBody() *GetRegionStatusResponseBody
}

type GetRegionStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetRegionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegionStatusResponse) GetBody() *GetRegionStatusResponseBody {
	return s.Body
}

func (s *GetRegionStatusResponse) SetHeaders(v map[string]*string) *GetRegionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetRegionStatusResponse) SetStatusCode(v int32) *GetRegionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegionStatusResponse) SetBody(v *GetRegionStatusResponseBody) *GetRegionStatusResponse {
	s.Body = v
	return s
}

func (s *GetRegionStatusResponse) Validate() error {
	return dara.Validate(s)
}
