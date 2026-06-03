// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySdkVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySdkVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySdkVersionResponse
	GetStatusCode() *int32
	SetBody(v *QuerySdkVersionResponseBody) *QuerySdkVersionResponse
	GetBody() *QuerySdkVersionResponseBody
}

type QuerySdkVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySdkVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySdkVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySdkVersionResponse) GoString() string {
	return s.String()
}

func (s *QuerySdkVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySdkVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySdkVersionResponse) GetBody() *QuerySdkVersionResponseBody {
	return s.Body
}

func (s *QuerySdkVersionResponse) SetHeaders(v map[string]*string) *QuerySdkVersionResponse {
	s.Headers = v
	return s
}

func (s *QuerySdkVersionResponse) SetStatusCode(v int32) *QuerySdkVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySdkVersionResponse) SetBody(v *QuerySdkVersionResponseBody) *QuerySdkVersionResponse {
	s.Body = v
	return s
}

func (s *QuerySdkVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
