// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetResourceConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetResourceConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetResourceConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetResourceConfigurationsResponseBody) *BatchGetResourceConfigurationsResponse
	GetBody() *BatchGetResourceConfigurationsResponseBody
}

type BatchGetResourceConfigurationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetResourceConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetResourceConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetResourceConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *BatchGetResourceConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetResourceConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetResourceConfigurationsResponse) GetBody() *BatchGetResourceConfigurationsResponseBody {
	return s.Body
}

func (s *BatchGetResourceConfigurationsResponse) SetHeaders(v map[string]*string) *BatchGetResourceConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *BatchGetResourceConfigurationsResponse) SetStatusCode(v int32) *BatchGetResourceConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponse) SetBody(v *BatchGetResourceConfigurationsResponseBody) *BatchGetResourceConfigurationsResponse {
	s.Body = v
	return s
}

func (s *BatchGetResourceConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
