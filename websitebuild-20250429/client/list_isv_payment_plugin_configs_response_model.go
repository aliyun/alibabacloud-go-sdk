// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIsvPaymentPluginConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIsvPaymentPluginConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIsvPaymentPluginConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListIsvPaymentPluginConfigsResponseBody) *ListIsvPaymentPluginConfigsResponse
	GetBody() *ListIsvPaymentPluginConfigsResponseBody
}

type ListIsvPaymentPluginConfigsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIsvPaymentPluginConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIsvPaymentPluginConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIsvPaymentPluginConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListIsvPaymentPluginConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIsvPaymentPluginConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIsvPaymentPluginConfigsResponse) GetBody() *ListIsvPaymentPluginConfigsResponseBody {
	return s.Body
}

func (s *ListIsvPaymentPluginConfigsResponse) SetHeaders(v map[string]*string) *ListIsvPaymentPluginConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponse) SetStatusCode(v int32) *ListIsvPaymentPluginConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponse) SetBody(v *ListIsvPaymentPluginConfigsResponseBody) *ListIsvPaymentPluginConfigsResponse {
	s.Body = v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
