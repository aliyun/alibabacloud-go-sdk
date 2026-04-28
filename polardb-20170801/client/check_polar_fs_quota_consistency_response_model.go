// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckPolarFsQuotaConsistencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckPolarFsQuotaConsistencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckPolarFsQuotaConsistencyResponse
	GetStatusCode() *int32
	SetBody(v *CheckPolarFsQuotaConsistencyResponseBody) *CheckPolarFsQuotaConsistencyResponse
	GetBody() *CheckPolarFsQuotaConsistencyResponseBody
}

type CheckPolarFsQuotaConsistencyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckPolarFsQuotaConsistencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckPolarFsQuotaConsistencyResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckPolarFsQuotaConsistencyResponse) GoString() string {
	return s.String()
}

func (s *CheckPolarFsQuotaConsistencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckPolarFsQuotaConsistencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckPolarFsQuotaConsistencyResponse) GetBody() *CheckPolarFsQuotaConsistencyResponseBody {
	return s.Body
}

func (s *CheckPolarFsQuotaConsistencyResponse) SetHeaders(v map[string]*string) *CheckPolarFsQuotaConsistencyResponse {
	s.Headers = v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponse) SetStatusCode(v int32) *CheckPolarFsQuotaConsistencyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponse) SetBody(v *CheckPolarFsQuotaConsistencyResponseBody) *CheckPolarFsQuotaConsistencyResponse {
	s.Body = v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
