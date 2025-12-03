// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterServiceConfigHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterServiceConfigHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterServiceConfigHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterServiceConfigHistoryResponseBody) *ListClusterServiceConfigHistoryResponse
	GetBody() *ListClusterServiceConfigHistoryResponseBody
}

type ListClusterServiceConfigHistoryResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterServiceConfigHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterServiceConfigHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterServiceConfigHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterServiceConfigHistoryResponse) GetBody() *ListClusterServiceConfigHistoryResponseBody {
	return s.Body
}

func (s *ListClusterServiceConfigHistoryResponse) SetHeaders(v map[string]*string) *ListClusterServiceConfigHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListClusterServiceConfigHistoryResponse) SetStatusCode(v int32) *ListClusterServiceConfigHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponse) SetBody(v *ListClusterServiceConfigHistoryResponseBody) *ListClusterServiceConfigHistoryResponse {
	s.Body = v
	return s
}

func (s *ListClusterServiceConfigHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
