// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaasDataUnitUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSaasDataUnitUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSaasDataUnitUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetSaasDataUnitUsageResponseBody) *GetSaasDataUnitUsageResponse
	GetBody() *GetSaasDataUnitUsageResponseBody
}

type GetSaasDataUnitUsageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSaasDataUnitUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSaasDataUnitUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponse) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSaasDataUnitUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSaasDataUnitUsageResponse) GetBody() *GetSaasDataUnitUsageResponseBody {
	return s.Body
}

func (s *GetSaasDataUnitUsageResponse) SetHeaders(v map[string]*string) *GetSaasDataUnitUsageResponse {
	s.Headers = v
	return s
}

func (s *GetSaasDataUnitUsageResponse) SetStatusCode(v int32) *GetSaasDataUnitUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSaasDataUnitUsageResponse) SetBody(v *GetSaasDataUnitUsageResponseBody) *GetSaasDataUnitUsageResponse {
	s.Body = v
	return s
}

func (s *GetSaasDataUnitUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
