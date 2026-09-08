// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMiniEngineVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMiniEngineVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMiniEngineVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListMiniEngineVersionsResponseBody) *ListMiniEngineVersionsResponse
	GetBody() *ListMiniEngineVersionsResponseBody
}

type ListMiniEngineVersionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMiniEngineVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMiniEngineVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMiniEngineVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListMiniEngineVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMiniEngineVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMiniEngineVersionsResponse) GetBody() *ListMiniEngineVersionsResponseBody {
	return s.Body
}

func (s *ListMiniEngineVersionsResponse) SetHeaders(v map[string]*string) *ListMiniEngineVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListMiniEngineVersionsResponse) SetStatusCode(v int32) *ListMiniEngineVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMiniEngineVersionsResponse) SetBody(v *ListMiniEngineVersionsResponseBody) *ListMiniEngineVersionsResponse {
	s.Body = v
	return s
}

func (s *ListMiniEngineVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
