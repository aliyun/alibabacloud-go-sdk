// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyLogisticsSmsMailNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyLogisticsSmsMailNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyLogisticsSmsMailNoResponse
	GetStatusCode() *int32
	SetBody(v *VerifyLogisticsSmsMailNoResponseBody) *VerifyLogisticsSmsMailNoResponse
	GetBody() *VerifyLogisticsSmsMailNoResponseBody
}

type VerifyLogisticsSmsMailNoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyLogisticsSmsMailNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyLogisticsSmsMailNoResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyLogisticsSmsMailNoResponse) GoString() string {
	return s.String()
}

func (s *VerifyLogisticsSmsMailNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyLogisticsSmsMailNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyLogisticsSmsMailNoResponse) GetBody() *VerifyLogisticsSmsMailNoResponseBody {
	return s.Body
}

func (s *VerifyLogisticsSmsMailNoResponse) SetHeaders(v map[string]*string) *VerifyLogisticsSmsMailNoResponse {
	s.Headers = v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponse) SetStatusCode(v int32) *VerifyLogisticsSmsMailNoResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponse) SetBody(v *VerifyLogisticsSmsMailNoResponseBody) *VerifyLogisticsSmsMailNoResponse {
	s.Body = v
	return s
}

func (s *VerifyLogisticsSmsMailNoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
