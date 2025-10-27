// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJenkinsImageRegistryPersistenceDayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJenkinsImageRegistryPersistenceDayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJenkinsImageRegistryPersistenceDayResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJenkinsImageRegistryPersistenceDayResponseBody) *UpdateJenkinsImageRegistryPersistenceDayResponse
	GetBody() *UpdateJenkinsImageRegistryPersistenceDayResponseBody
}

type UpdateJenkinsImageRegistryPersistenceDayResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJenkinsImageRegistryPersistenceDayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJenkinsImageRegistryPersistenceDayResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJenkinsImageRegistryPersistenceDayResponse) GoString() string {
	return s.String()
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponse) GetBody() *UpdateJenkinsImageRegistryPersistenceDayResponseBody {
	return s.Body
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponse) SetHeaders(v map[string]*string) *UpdateJenkinsImageRegistryPersistenceDayResponse {
	s.Headers = v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponse) SetStatusCode(v int32) *UpdateJenkinsImageRegistryPersistenceDayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponse) SetBody(v *UpdateJenkinsImageRegistryPersistenceDayResponseBody) *UpdateJenkinsImageRegistryPersistenceDayResponse {
	s.Body = v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
