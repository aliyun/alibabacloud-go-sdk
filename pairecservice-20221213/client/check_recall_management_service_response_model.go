// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRecallManagementServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckRecallManagementServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckRecallManagementServiceResponse
	GetStatusCode() *int32
	SetBody(v *CheckRecallManagementServiceResponseBody) *CheckRecallManagementServiceResponse
	GetBody() *CheckRecallManagementServiceResponseBody
}

type CheckRecallManagementServiceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckRecallManagementServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckRecallManagementServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckRecallManagementServiceResponse) GoString() string {
	return s.String()
}

func (s *CheckRecallManagementServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckRecallManagementServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckRecallManagementServiceResponse) GetBody() *CheckRecallManagementServiceResponseBody {
	return s.Body
}

func (s *CheckRecallManagementServiceResponse) SetHeaders(v map[string]*string) *CheckRecallManagementServiceResponse {
	s.Headers = v
	return s
}

func (s *CheckRecallManagementServiceResponse) SetStatusCode(v int32) *CheckRecallManagementServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRecallManagementServiceResponse) SetBody(v *CheckRecallManagementServiceResponseBody) *CheckRecallManagementServiceResponse {
	s.Body = v
	return s
}

func (s *CheckRecallManagementServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
