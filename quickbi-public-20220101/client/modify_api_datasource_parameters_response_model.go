// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiDatasourceParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApiDatasourceParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApiDatasourceParametersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApiDatasourceParametersResponseBody) *ModifyApiDatasourceParametersResponse
	GetBody() *ModifyApiDatasourceParametersResponseBody
}

type ModifyApiDatasourceParametersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiDatasourceParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiDatasourceParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiDatasourceParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiDatasourceParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApiDatasourceParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApiDatasourceParametersResponse) GetBody() *ModifyApiDatasourceParametersResponseBody {
	return s.Body
}

func (s *ModifyApiDatasourceParametersResponse) SetHeaders(v map[string]*string) *ModifyApiDatasourceParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiDatasourceParametersResponse) SetStatusCode(v int32) *ModifyApiDatasourceParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiDatasourceParametersResponse) SetBody(v *ModifyApiDatasourceParametersResponseBody) *ModifyApiDatasourceParametersResponse {
	s.Body = v
	return s
}

func (s *ModifyApiDatasourceParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
