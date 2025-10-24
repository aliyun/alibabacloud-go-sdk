// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMmsDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMmsDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetMmsDataSourceResponseBody) *GetMmsDataSourceResponse
	GetBody() *GetMmsDataSourceResponseBody
}

type GetMmsDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMmsDataSourceResponse) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMmsDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMmsDataSourceResponse) GetBody() *GetMmsDataSourceResponseBody {
	return s.Body
}

func (s *GetMmsDataSourceResponse) SetHeaders(v map[string]*string) *GetMmsDataSourceResponse {
	s.Headers = v
	return s
}

func (s *GetMmsDataSourceResponse) SetStatusCode(v int32) *GetMmsDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsDataSourceResponse) SetBody(v *GetMmsDataSourceResponseBody) *GetMmsDataSourceResponse {
	s.Body = v
	return s
}

func (s *GetMmsDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
