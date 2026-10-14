// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgenticDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgenticDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetAgenticDataSourceResponseBody) *GetAgenticDataSourceResponse
	GetBody() *GetAgenticDataSourceResponseBody
}

type GetAgenticDataSourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgenticDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgenticDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticDataSourceResponse) GoString() string {
	return s.String()
}

func (s *GetAgenticDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgenticDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgenticDataSourceResponse) GetBody() *GetAgenticDataSourceResponseBody {
	return s.Body
}

func (s *GetAgenticDataSourceResponse) SetHeaders(v map[string]*string) *GetAgenticDataSourceResponse {
	s.Headers = v
	return s
}

func (s *GetAgenticDataSourceResponse) SetStatusCode(v int32) *GetAgenticDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgenticDataSourceResponse) SetBody(v *GetAgenticDataSourceResponseBody) *GetAgenticDataSourceResponse {
	s.Body = v
	return s
}

func (s *GetAgenticDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
