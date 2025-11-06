// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMonitorSLRResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudMonitorSLRResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudMonitorSLRResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudMonitorSLRResponseBody) *CreateCloudMonitorSLRResponse
	GetBody() *CreateCloudMonitorSLRResponseBody
}

type CreateCloudMonitorSLRResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudMonitorSLRResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudMonitorSLRResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMonitorSLRResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudMonitorSLRResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudMonitorSLRResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudMonitorSLRResponse) GetBody() *CreateCloudMonitorSLRResponseBody {
	return s.Body
}

func (s *CreateCloudMonitorSLRResponse) SetHeaders(v map[string]*string) *CreateCloudMonitorSLRResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudMonitorSLRResponse) SetStatusCode(v int32) *CreateCloudMonitorSLRResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudMonitorSLRResponse) SetBody(v *CreateCloudMonitorSLRResponseBody) *CreateCloudMonitorSLRResponse {
	s.Body = v
	return s
}

func (s *CreateCloudMonitorSLRResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
