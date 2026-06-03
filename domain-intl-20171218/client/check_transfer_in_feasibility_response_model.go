// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTransferInFeasibilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckTransferInFeasibilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckTransferInFeasibilityResponse
	GetStatusCode() *int32
	SetBody(v *CheckTransferInFeasibilityResponseBody) *CheckTransferInFeasibilityResponse
	GetBody() *CheckTransferInFeasibilityResponseBody
}

type CheckTransferInFeasibilityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckTransferInFeasibilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckTransferInFeasibilityResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckTransferInFeasibilityResponse) GoString() string {
	return s.String()
}

func (s *CheckTransferInFeasibilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckTransferInFeasibilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckTransferInFeasibilityResponse) GetBody() *CheckTransferInFeasibilityResponseBody {
	return s.Body
}

func (s *CheckTransferInFeasibilityResponse) SetHeaders(v map[string]*string) *CheckTransferInFeasibilityResponse {
	s.Headers = v
	return s
}

func (s *CheckTransferInFeasibilityResponse) SetStatusCode(v int32) *CheckTransferInFeasibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckTransferInFeasibilityResponse) SetBody(v *CheckTransferInFeasibilityResponseBody) *CheckTransferInFeasibilityResponse {
	s.Body = v
	return s
}

func (s *CheckTransferInFeasibilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
