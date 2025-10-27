// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicesByAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDevicesByAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDevicesByAliasResponse
	GetStatusCode() *int32
	SetBody(v *QueryDevicesByAliasResponseBody) *QueryDevicesByAliasResponse
	GetBody() *QueryDevicesByAliasResponseBody
}

type QueryDevicesByAliasResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicesByAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDevicesByAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicesByAliasResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDevicesByAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDevicesByAliasResponse) GetBody() *QueryDevicesByAliasResponseBody {
	return s.Body
}

func (s *QueryDevicesByAliasResponse) SetHeaders(v map[string]*string) *QueryDevicesByAliasResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicesByAliasResponse) SetStatusCode(v int32) *QueryDevicesByAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicesByAliasResponse) SetBody(v *QueryDevicesByAliasResponseBody) *QueryDevicesByAliasResponse {
	s.Body = v
	return s
}

func (s *QueryDevicesByAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
