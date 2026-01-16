// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCloudGtmInstanceConfigLogSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCloudGtmInstanceConfigLogSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCloudGtmInstanceConfigLogSwitchResponse
	GetStatusCode() *int32
	SetBody(v *SetCloudGtmInstanceConfigLogSwitchResponseBody) *SetCloudGtmInstanceConfigLogSwitchResponse
	GetBody() *SetCloudGtmInstanceConfigLogSwitchResponseBody
}

type SetCloudGtmInstanceConfigLogSwitchResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCloudGtmInstanceConfigLogSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCloudGtmInstanceConfigLogSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCloudGtmInstanceConfigLogSwitchResponse) GoString() string {
	return s.String()
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponse) GetBody() *SetCloudGtmInstanceConfigLogSwitchResponseBody {
	return s.Body
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponse) SetHeaders(v map[string]*string) *SetCloudGtmInstanceConfigLogSwitchResponse {
	s.Headers = v
	return s
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponse) SetStatusCode(v int32) *SetCloudGtmInstanceConfigLogSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponse) SetBody(v *SetCloudGtmInstanceConfigLogSwitchResponseBody) *SetCloudGtmInstanceConfigLogSwitchResponse {
	s.Body = v
	return s
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
