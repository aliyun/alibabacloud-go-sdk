// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLabelConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceLabelConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceLabelConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceLabelConfigResponseBody) *GetServiceLabelConfigResponse
	GetBody() *GetServiceLabelConfigResponseBody
}

type GetServiceLabelConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceLabelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceLabelConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLabelConfigResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLabelConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceLabelConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceLabelConfigResponse) GetBody() *GetServiceLabelConfigResponseBody {
	return s.Body
}

func (s *GetServiceLabelConfigResponse) SetHeaders(v map[string]*string) *GetServiceLabelConfigResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLabelConfigResponse) SetStatusCode(v int32) *GetServiceLabelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceLabelConfigResponse) SetBody(v *GetServiceLabelConfigResponseBody) *GetServiceLabelConfigResponse {
	s.Body = v
	return s
}

func (s *GetServiceLabelConfigResponse) Validate() error {
	return dara.Validate(s)
}
