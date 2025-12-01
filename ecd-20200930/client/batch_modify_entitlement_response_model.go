// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchModifyEntitlementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchModifyEntitlementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchModifyEntitlementResponse
	GetStatusCode() *int32
	SetBody(v *BatchModifyEntitlementResponseBody) *BatchModifyEntitlementResponse
	GetBody() *BatchModifyEntitlementResponseBody
}

type BatchModifyEntitlementResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchModifyEntitlementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchModifyEntitlementResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchModifyEntitlementResponse) GoString() string {
	return s.String()
}

func (s *BatchModifyEntitlementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchModifyEntitlementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchModifyEntitlementResponse) GetBody() *BatchModifyEntitlementResponseBody {
	return s.Body
}

func (s *BatchModifyEntitlementResponse) SetHeaders(v map[string]*string) *BatchModifyEntitlementResponse {
	s.Headers = v
	return s
}

func (s *BatchModifyEntitlementResponse) SetStatusCode(v int32) *BatchModifyEntitlementResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchModifyEntitlementResponse) SetBody(v *BatchModifyEntitlementResponseBody) *BatchModifyEntitlementResponse {
	s.Body = v
	return s
}

func (s *BatchModifyEntitlementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
