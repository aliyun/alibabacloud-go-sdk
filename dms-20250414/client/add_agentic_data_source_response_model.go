// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgenticDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAgenticDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAgenticDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *AddAgenticDataSourceResponseBody) *AddAgenticDataSourceResponse
	GetBody() *AddAgenticDataSourceResponseBody
}

type AddAgenticDataSourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAgenticDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAgenticDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAgenticDataSourceResponse) GoString() string {
	return s.String()
}

func (s *AddAgenticDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAgenticDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAgenticDataSourceResponse) GetBody() *AddAgenticDataSourceResponseBody {
	return s.Body
}

func (s *AddAgenticDataSourceResponse) SetHeaders(v map[string]*string) *AddAgenticDataSourceResponse {
	s.Headers = v
	return s
}

func (s *AddAgenticDataSourceResponse) SetStatusCode(v int32) *AddAgenticDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAgenticDataSourceResponse) SetBody(v *AddAgenticDataSourceResponseBody) *AddAgenticDataSourceResponse {
	s.Body = v
	return s
}

func (s *AddAgenticDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
