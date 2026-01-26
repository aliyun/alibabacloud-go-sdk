// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPromInstallStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPromInstallStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPromInstallStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryPromInstallStatusResponseBody) *QueryPromInstallStatusResponse
	GetBody() *QueryPromInstallStatusResponseBody
}

type QueryPromInstallStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPromInstallStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPromInstallStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPromInstallStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryPromInstallStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPromInstallStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPromInstallStatusResponse) GetBody() *QueryPromInstallStatusResponseBody {
	return s.Body
}

func (s *QueryPromInstallStatusResponse) SetHeaders(v map[string]*string) *QueryPromInstallStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryPromInstallStatusResponse) SetStatusCode(v int32) *QueryPromInstallStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPromInstallStatusResponse) SetBody(v *QueryPromInstallStatusResponseBody) *QueryPromInstallStatusResponse {
	s.Body = v
	return s
}

func (s *QueryPromInstallStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
