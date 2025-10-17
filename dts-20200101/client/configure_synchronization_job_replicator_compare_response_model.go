// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSynchronizationJobReplicatorCompareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureSynchronizationJobReplicatorCompareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureSynchronizationJobReplicatorCompareResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureSynchronizationJobReplicatorCompareResponseBody) *ConfigureSynchronizationJobReplicatorCompareResponse
	GetBody() *ConfigureSynchronizationJobReplicatorCompareResponseBody
}

type ConfigureSynchronizationJobReplicatorCompareResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureSynchronizationJobReplicatorCompareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureSynchronizationJobReplicatorCompareResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobReplicatorCompareResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponse) GetBody() *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	return s.Body
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponse) SetHeaders(v map[string]*string) *ConfigureSynchronizationJobReplicatorCompareResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponse) SetStatusCode(v int32) *ConfigureSynchronizationJobReplicatorCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponse) SetBody(v *ConfigureSynchronizationJobReplicatorCompareResponseBody) *ConfigureSynchronizationJobReplicatorCompareResponse {
	s.Body = v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
