// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceConnectionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataSourceConnectionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataSourceConnectionInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDataSourceConnectionInfoResponseBody) *GetDataSourceConnectionInfoResponse
	GetBody() *GetDataSourceConnectionInfoResponseBody
}

type GetDataSourceConnectionInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceConnectionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceConnectionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceConnectionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceConnectionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataSourceConnectionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataSourceConnectionInfoResponse) GetBody() *GetDataSourceConnectionInfoResponseBody {
	return s.Body
}

func (s *GetDataSourceConnectionInfoResponse) SetHeaders(v map[string]*string) *GetDataSourceConnectionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceConnectionInfoResponse) SetStatusCode(v int32) *GetDataSourceConnectionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponse) SetBody(v *GetDataSourceConnectionInfoResponseBody) *GetDataSourceConnectionInfoResponse {
	s.Body = v
	return s
}

func (s *GetDataSourceConnectionInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
