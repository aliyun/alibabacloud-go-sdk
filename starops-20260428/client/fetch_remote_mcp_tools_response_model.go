// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchRemoteMcpToolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchRemoteMcpToolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchRemoteMcpToolsResponse
	GetStatusCode() *int32
	SetBody(v *FetchRemoteMcpToolsResponseBody) *FetchRemoteMcpToolsResponse
	GetBody() *FetchRemoteMcpToolsResponseBody
}

type FetchRemoteMcpToolsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchRemoteMcpToolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchRemoteMcpToolsResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchRemoteMcpToolsResponse) GoString() string {
	return s.String()
}

func (s *FetchRemoteMcpToolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchRemoteMcpToolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchRemoteMcpToolsResponse) GetBody() *FetchRemoteMcpToolsResponseBody {
	return s.Body
}

func (s *FetchRemoteMcpToolsResponse) SetHeaders(v map[string]*string) *FetchRemoteMcpToolsResponse {
	s.Headers = v
	return s
}

func (s *FetchRemoteMcpToolsResponse) SetStatusCode(v int32) *FetchRemoteMcpToolsResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchRemoteMcpToolsResponse) SetBody(v *FetchRemoteMcpToolsResponseBody) *FetchRemoteMcpToolsResponse {
	s.Body = v
	return s
}

func (s *FetchRemoteMcpToolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
