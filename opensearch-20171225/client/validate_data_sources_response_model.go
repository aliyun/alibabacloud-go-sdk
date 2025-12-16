// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ValidateDataSourcesResponseBody) *ValidateDataSourcesResponse
	GetBody() *ValidateDataSourcesResponseBody
}

type ValidateDataSourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateDataSourcesResponse) GetBody() *ValidateDataSourcesResponseBody {
	return s.Body
}

func (s *ValidateDataSourcesResponse) SetHeaders(v map[string]*string) *ValidateDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ValidateDataSourcesResponse) SetStatusCode(v int32) *ValidateDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateDataSourcesResponse) SetBody(v *ValidateDataSourcesResponseBody) *ValidateDataSourcesResponse {
	s.Body = v
	return s
}

func (s *ValidateDataSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
