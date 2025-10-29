// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveDelayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveDelayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveDelayConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveDelayConfigResponseBody) *ListLiveDelayConfigResponse
	GetBody() *ListLiveDelayConfigResponseBody
}

type ListLiveDelayConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveDelayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveDelayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveDelayConfigResponse) GoString() string {
	return s.String()
}

func (s *ListLiveDelayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveDelayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveDelayConfigResponse) GetBody() *ListLiveDelayConfigResponseBody {
	return s.Body
}

func (s *ListLiveDelayConfigResponse) SetHeaders(v map[string]*string) *ListLiveDelayConfigResponse {
	s.Headers = v
	return s
}

func (s *ListLiveDelayConfigResponse) SetStatusCode(v int32) *ListLiveDelayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveDelayConfigResponse) SetBody(v *ListLiveDelayConfigResponseBody) *ListLiveDelayConfigResponse {
	s.Body = v
	return s
}

func (s *ListLiveDelayConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
