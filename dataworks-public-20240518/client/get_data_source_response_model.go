// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetDataSourceResponseBody) *GetDataSourceResponse
	GetBody() *GetDataSourceResponseBody
}

type GetDataSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataSourceResponse) GetBody() *GetDataSourceResponseBody {
	return s.Body
}

func (s *GetDataSourceResponse) SetHeaders(v map[string]*string) *GetDataSourceResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceResponse) SetStatusCode(v int32) *GetDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceResponse) SetBody(v *GetDataSourceResponseBody) *GetDataSourceResponse {
	s.Body = v
	return s
}

func (s *GetDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
