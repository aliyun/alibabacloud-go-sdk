// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertK8sResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertK8sResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertK8sResourceResponse
	GetStatusCode() *int32
	SetBody(v *ConvertK8sResourceResponseBody) *ConvertK8sResourceResponse
	GetBody() *ConvertK8sResourceResponseBody
}

type ConvertK8sResourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertK8sResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertK8sResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertK8sResourceResponse) GoString() string {
	return s.String()
}

func (s *ConvertK8sResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertK8sResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertK8sResourceResponse) GetBody() *ConvertK8sResourceResponseBody {
	return s.Body
}

func (s *ConvertK8sResourceResponse) SetHeaders(v map[string]*string) *ConvertK8sResourceResponse {
	s.Headers = v
	return s
}

func (s *ConvertK8sResourceResponse) SetStatusCode(v int32) *ConvertK8sResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertK8sResourceResponse) SetBody(v *ConvertK8sResourceResponseBody) *ConvertK8sResourceResponse {
	s.Body = v
	return s
}

func (s *ConvertK8sResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
