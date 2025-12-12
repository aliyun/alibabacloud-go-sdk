// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceNodeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartInstanceNodeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartInstanceNodeListResponse
	GetStatusCode() *int32
	SetBody(v *RestartInstanceNodeListResponseBody) *RestartInstanceNodeListResponse
	GetBody() *RestartInstanceNodeListResponseBody
}

type RestartInstanceNodeListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartInstanceNodeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartInstanceNodeListResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceNodeListResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceNodeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartInstanceNodeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartInstanceNodeListResponse) GetBody() *RestartInstanceNodeListResponseBody {
	return s.Body
}

func (s *RestartInstanceNodeListResponse) SetHeaders(v map[string]*string) *RestartInstanceNodeListResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceNodeListResponse) SetStatusCode(v int32) *RestartInstanceNodeListResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartInstanceNodeListResponse) SetBody(v *RestartInstanceNodeListResponseBody) *RestartInstanceNodeListResponse {
	s.Body = v
	return s
}

func (s *RestartInstanceNodeListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
