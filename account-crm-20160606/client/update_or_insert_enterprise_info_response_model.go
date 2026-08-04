// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrInsertEnterpriseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOrInsertEnterpriseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOrInsertEnterpriseInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOrInsertEnterpriseInfoResponseBody) *UpdateOrInsertEnterpriseInfoResponse
	GetBody() *UpdateOrInsertEnterpriseInfoResponseBody
}

type UpdateOrInsertEnterpriseInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOrInsertEnterpriseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOrInsertEnterpriseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrInsertEnterpriseInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrInsertEnterpriseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOrInsertEnterpriseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOrInsertEnterpriseInfoResponse) GetBody() *UpdateOrInsertEnterpriseInfoResponseBody {
	return s.Body
}

func (s *UpdateOrInsertEnterpriseInfoResponse) SetHeaders(v map[string]*string) *UpdateOrInsertEnterpriseInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoResponse) SetStatusCode(v int32) *UpdateOrInsertEnterpriseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoResponse) SetBody(v *UpdateOrInsertEnterpriseInfoResponseBody) *UpdateOrInsertEnterpriseInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
