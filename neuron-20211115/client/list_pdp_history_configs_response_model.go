// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpHistoryConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPdpHistoryConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPdpHistoryConfigsResponse
	GetStatusCode() *int32
	SetBody(v *PdpHistoryConfigPageResult) *ListPdpHistoryConfigsResponse
	GetBody() *PdpHistoryConfigPageResult
}

type ListPdpHistoryConfigsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpHistoryConfigPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPdpHistoryConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPdpHistoryConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListPdpHistoryConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPdpHistoryConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPdpHistoryConfigsResponse) GetBody() *PdpHistoryConfigPageResult {
	return s.Body
}

func (s *ListPdpHistoryConfigsResponse) SetHeaders(v map[string]*string) *ListPdpHistoryConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListPdpHistoryConfigsResponse) SetStatusCode(v int32) *ListPdpHistoryConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPdpHistoryConfigsResponse) SetBody(v *PdpHistoryConfigPageResult) *ListPdpHistoryConfigsResponse {
	s.Body = v
	return s
}

func (s *ListPdpHistoryConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
