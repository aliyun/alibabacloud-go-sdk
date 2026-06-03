// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPnsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPnsConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryPnsConfigResponseBody) *QueryPnsConfigResponse
	GetBody() *QueryPnsConfigResponseBody
}

type QueryPnsConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPnsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPnsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryPnsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPnsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPnsConfigResponse) GetBody() *QueryPnsConfigResponseBody {
	return s.Body
}

func (s *QueryPnsConfigResponse) SetHeaders(v map[string]*string) *QueryPnsConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryPnsConfigResponse) SetStatusCode(v int32) *QueryPnsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPnsConfigResponse) SetBody(v *QueryPnsConfigResponseBody) *QueryPnsConfigResponse {
	s.Body = v
	return s
}

func (s *QueryPnsConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
