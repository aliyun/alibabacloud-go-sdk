// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelServiceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelServiceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelServiceUsageResponse
	GetStatusCode() *int32
	SetBody(v *CancelServiceUsageResponseBody) *CancelServiceUsageResponse
	GetBody() *CancelServiceUsageResponseBody
}

type CancelServiceUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelServiceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *CancelServiceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelServiceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelServiceUsageResponse) GetBody() *CancelServiceUsageResponseBody {
	return s.Body
}

func (s *CancelServiceUsageResponse) SetHeaders(v map[string]*string) *CancelServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *CancelServiceUsageResponse) SetStatusCode(v int32) *CancelServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelServiceUsageResponse) SetBody(v *CancelServiceUsageResponseBody) *CancelServiceUsageResponse {
	s.Body = v
	return s
}

func (s *CancelServiceUsageResponse) Validate() error {
	return dara.Validate(s)
}
