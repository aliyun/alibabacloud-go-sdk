// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmsDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmsDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmsDataSourceResponseBody) *UpdateMmsDataSourceResponse
	GetBody() *UpdateMmsDataSourceResponseBody
}

type UpdateMmsDataSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmsDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmsDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsDataSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmsDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmsDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmsDataSourceResponse) GetBody() *UpdateMmsDataSourceResponseBody {
	return s.Body
}

func (s *UpdateMmsDataSourceResponse) SetHeaders(v map[string]*string) *UpdateMmsDataSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmsDataSourceResponse) SetStatusCode(v int32) *UpdateMmsDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmsDataSourceResponse) SetBody(v *UpdateMmsDataSourceResponseBody) *UpdateMmsDataSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateMmsDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
