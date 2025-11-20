// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProxiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelProxiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelProxiesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelProxiesResult) *ListModelProxiesResponse
	GetBody() *ListModelProxiesResult
}

type ListModelProxiesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelProxiesResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelProxiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelProxiesResponse) GoString() string {
	return s.String()
}

func (s *ListModelProxiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelProxiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelProxiesResponse) GetBody() *ListModelProxiesResult {
	return s.Body
}

func (s *ListModelProxiesResponse) SetHeaders(v map[string]*string) *ListModelProxiesResponse {
	s.Headers = v
	return s
}

func (s *ListModelProxiesResponse) SetStatusCode(v int32) *ListModelProxiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelProxiesResponse) SetBody(v *ListModelProxiesResult) *ListModelProxiesResponse {
	s.Body = v
	return s
}

func (s *ListModelProxiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
