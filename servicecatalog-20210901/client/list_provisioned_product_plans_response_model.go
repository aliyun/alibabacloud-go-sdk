// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionedProductPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProvisionedProductPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProvisionedProductPlansResponse
	GetStatusCode() *int32
	SetBody(v *ListProvisionedProductPlansResponseBody) *ListProvisionedProductPlansResponse
	GetBody() *ListProvisionedProductPlansResponseBody
}

type ListProvisionedProductPlansResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProvisionedProductPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProvisionedProductPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlansResponse) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProvisionedProductPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProvisionedProductPlansResponse) GetBody() *ListProvisionedProductPlansResponseBody {
	return s.Body
}

func (s *ListProvisionedProductPlansResponse) SetHeaders(v map[string]*string) *ListProvisionedProductPlansResponse {
	s.Headers = v
	return s
}

func (s *ListProvisionedProductPlansResponse) SetStatusCode(v int32) *ListProvisionedProductPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProvisionedProductPlansResponse) SetBody(v *ListProvisionedProductPlansResponseBody) *ListProvisionedProductPlansResponse {
	s.Body = v
	return s
}

func (s *ListProvisionedProductPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
