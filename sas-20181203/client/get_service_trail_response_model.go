// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTrailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceTrailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceTrailResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceTrailResponseBody) *GetServiceTrailResponse
	GetBody() *GetServiceTrailResponseBody
}

type GetServiceTrailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTrailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTrailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTrailResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTrailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceTrailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceTrailResponse) GetBody() *GetServiceTrailResponseBody {
	return s.Body
}

func (s *GetServiceTrailResponse) SetHeaders(v map[string]*string) *GetServiceTrailResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTrailResponse) SetStatusCode(v int32) *GetServiceTrailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTrailResponse) SetBody(v *GetServiceTrailResponseBody) *GetServiceTrailResponse {
	s.Body = v
	return s
}

func (s *GetServiceTrailResponse) Validate() error {
	return dara.Validate(s)
}
