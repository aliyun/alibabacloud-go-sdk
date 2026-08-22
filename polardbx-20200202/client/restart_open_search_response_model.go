// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartOpenSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartOpenSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartOpenSearchResponse
	GetStatusCode() *int32
	SetBody(v *RestartOpenSearchResponseBody) *RestartOpenSearchResponse
	GetBody() *RestartOpenSearchResponseBody
}

type RestartOpenSearchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartOpenSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartOpenSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartOpenSearchResponse) GoString() string {
	return s.String()
}

func (s *RestartOpenSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartOpenSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartOpenSearchResponse) GetBody() *RestartOpenSearchResponseBody {
	return s.Body
}

func (s *RestartOpenSearchResponse) SetHeaders(v map[string]*string) *RestartOpenSearchResponse {
	s.Headers = v
	return s
}

func (s *RestartOpenSearchResponse) SetStatusCode(v int32) *RestartOpenSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartOpenSearchResponse) SetBody(v *RestartOpenSearchResponseBody) *RestartOpenSearchResponse {
	s.Body = v
	return s
}

func (s *RestartOpenSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
