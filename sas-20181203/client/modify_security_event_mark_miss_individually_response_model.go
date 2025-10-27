// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityEventMarkMissIndividuallyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityEventMarkMissIndividuallyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityEventMarkMissIndividuallyResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityEventMarkMissIndividuallyResponseBody) *ModifySecurityEventMarkMissIndividuallyResponse
	GetBody() *ModifySecurityEventMarkMissIndividuallyResponseBody
}

type ModifySecurityEventMarkMissIndividuallyResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityEventMarkMissIndividuallyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityEventMarkMissIndividuallyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityEventMarkMissIndividuallyResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityEventMarkMissIndividuallyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityEventMarkMissIndividuallyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityEventMarkMissIndividuallyResponse) GetBody() *ModifySecurityEventMarkMissIndividuallyResponseBody {
	return s.Body
}

func (s *ModifySecurityEventMarkMissIndividuallyResponse) SetHeaders(v map[string]*string) *ModifySecurityEventMarkMissIndividuallyResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyResponse) SetStatusCode(v int32) *ModifySecurityEventMarkMissIndividuallyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyResponse) SetBody(v *ModifySecurityEventMarkMissIndividuallyResponseBody) *ModifySecurityEventMarkMissIndividuallyResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
