// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySaasServiceDeletionProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySaasServiceDeletionProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySaasServiceDeletionProtectionResponse
	GetStatusCode() *int32
	SetBody(v *ModifySaasServiceDeletionProtectionResponseBody) *ModifySaasServiceDeletionProtectionResponse
	GetBody() *ModifySaasServiceDeletionProtectionResponseBody
}

type ModifySaasServiceDeletionProtectionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySaasServiceDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySaasServiceDeletionProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySaasServiceDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *ModifySaasServiceDeletionProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySaasServiceDeletionProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySaasServiceDeletionProtectionResponse) GetBody() *ModifySaasServiceDeletionProtectionResponseBody {
	return s.Body
}

func (s *ModifySaasServiceDeletionProtectionResponse) SetHeaders(v map[string]*string) *ModifySaasServiceDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *ModifySaasServiceDeletionProtectionResponse) SetStatusCode(v int32) *ModifySaasServiceDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySaasServiceDeletionProtectionResponse) SetBody(v *ModifySaasServiceDeletionProtectionResponseBody) *ModifySaasServiceDeletionProtectionResponse {
	s.Body = v
	return s
}

func (s *ModifySaasServiceDeletionProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
