// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationVswitchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationVswitchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationVswitchesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationVswitchesResponseBody) *UpdateApplicationVswitchesResponse
	GetBody() *UpdateApplicationVswitchesResponseBody
}

type UpdateApplicationVswitchesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationVswitchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationVswitchesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVswitchesResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVswitchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationVswitchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationVswitchesResponse) GetBody() *UpdateApplicationVswitchesResponseBody {
	return s.Body
}

func (s *UpdateApplicationVswitchesResponse) SetHeaders(v map[string]*string) *UpdateApplicationVswitchesResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationVswitchesResponse) SetStatusCode(v int32) *UpdateApplicationVswitchesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationVswitchesResponse) SetBody(v *UpdateApplicationVswitchesResponseBody) *UpdateApplicationVswitchesResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationVswitchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
