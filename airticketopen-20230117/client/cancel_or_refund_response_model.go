// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrRefundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelOrRefundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelOrRefundResponse
	GetStatusCode() *int32
	SetBody(v *CancelOrRefundResponseBody) *CancelOrRefundResponse
	GetBody() *CancelOrRefundResponseBody
}

type CancelOrRefundResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOrRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOrRefundResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelOrRefundResponse) GoString() string {
	return s.String()
}

func (s *CancelOrRefundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelOrRefundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelOrRefundResponse) GetBody() *CancelOrRefundResponseBody {
	return s.Body
}

func (s *CancelOrRefundResponse) SetHeaders(v map[string]*string) *CancelOrRefundResponse {
	s.Headers = v
	return s
}

func (s *CancelOrRefundResponse) SetStatusCode(v int32) *CancelOrRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOrRefundResponse) SetBody(v *CancelOrRefundResponseBody) *CancelOrRefundResponse {
	s.Body = v
	return s
}

func (s *CancelOrRefundResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
