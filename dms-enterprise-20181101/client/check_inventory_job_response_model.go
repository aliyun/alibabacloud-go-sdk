// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInventoryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInventoryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInventoryJobResponse
	GetStatusCode() *int32
	SetBody(v *CheckInventoryJobResponseBody) *CheckInventoryJobResponse
	GetBody() *CheckInventoryJobResponseBody
}

type CheckInventoryJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInventoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInventoryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInventoryJobResponse) GoString() string {
	return s.String()
}

func (s *CheckInventoryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInventoryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInventoryJobResponse) GetBody() *CheckInventoryJobResponseBody {
	return s.Body
}

func (s *CheckInventoryJobResponse) SetHeaders(v map[string]*string) *CheckInventoryJobResponse {
	s.Headers = v
	return s
}

func (s *CheckInventoryJobResponse) SetStatusCode(v int32) *CheckInventoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInventoryJobResponse) SetBody(v *CheckInventoryJobResponseBody) *CheckInventoryJobResponse {
	s.Body = v
	return s
}

func (s *CheckInventoryJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
