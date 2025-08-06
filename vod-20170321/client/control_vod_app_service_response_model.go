// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iControlVodAppServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ControlVodAppServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ControlVodAppServiceResponse
	GetStatusCode() *int32
	SetBody(v *ControlVodAppServiceResponseBody) *ControlVodAppServiceResponse
	GetBody() *ControlVodAppServiceResponseBody
}

type ControlVodAppServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ControlVodAppServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ControlVodAppServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ControlVodAppServiceResponse) GoString() string {
	return s.String()
}

func (s *ControlVodAppServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ControlVodAppServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ControlVodAppServiceResponse) GetBody() *ControlVodAppServiceResponseBody {
	return s.Body
}

func (s *ControlVodAppServiceResponse) SetHeaders(v map[string]*string) *ControlVodAppServiceResponse {
	s.Headers = v
	return s
}

func (s *ControlVodAppServiceResponse) SetStatusCode(v int32) *ControlVodAppServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ControlVodAppServiceResponse) SetBody(v *ControlVodAppServiceResponseBody) *ControlVodAppServiceResponse {
	s.Body = v
	return s
}

func (s *ControlVodAppServiceResponse) Validate() error {
	return dara.Validate(s)
}
