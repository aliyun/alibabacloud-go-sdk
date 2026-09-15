// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConnectorModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConnectorModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListConnectorModelsResponseBody) *ListConnectorModelsResponse
	GetBody() *ListConnectorModelsResponseBody
}

type ListConnectorModelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConnectorModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConnectorModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorModelsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectorModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConnectorModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConnectorModelsResponse) GetBody() *ListConnectorModelsResponseBody {
	return s.Body
}

func (s *ListConnectorModelsResponse) SetHeaders(v map[string]*string) *ListConnectorModelsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectorModelsResponse) SetStatusCode(v int32) *ListConnectorModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectorModelsResponse) SetBody(v *ListConnectorModelsResponseBody) *ListConnectorModelsResponse {
	s.Body = v
	return s
}

func (s *ListConnectorModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
