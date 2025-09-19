// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJenkinsImageRegistryPersistenceDayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryJenkinsImageRegistryPersistenceDayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryJenkinsImageRegistryPersistenceDayResponse
	GetStatusCode() *int32
	SetBody(v *QueryJenkinsImageRegistryPersistenceDayResponseBody) *QueryJenkinsImageRegistryPersistenceDayResponse
	GetBody() *QueryJenkinsImageRegistryPersistenceDayResponseBody
}

type QueryJenkinsImageRegistryPersistenceDayResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryJenkinsImageRegistryPersistenceDayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryJenkinsImageRegistryPersistenceDayResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryJenkinsImageRegistryPersistenceDayResponse) GoString() string {
	return s.String()
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponse) GetBody() *QueryJenkinsImageRegistryPersistenceDayResponseBody {
	return s.Body
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponse) SetHeaders(v map[string]*string) *QueryJenkinsImageRegistryPersistenceDayResponse {
	s.Headers = v
	return s
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponse) SetStatusCode(v int32) *QueryJenkinsImageRegistryPersistenceDayResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponse) SetBody(v *QueryJenkinsImageRegistryPersistenceDayResponseBody) *QueryJenkinsImageRegistryPersistenceDayResponse {
	s.Body = v
	return s
}

func (s *QueryJenkinsImageRegistryPersistenceDayResponse) Validate() error {
	return dara.Validate(s)
}
