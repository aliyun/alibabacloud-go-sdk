// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobScriptHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobScriptHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobScriptHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListJobScriptHistoryResponseBody) *ListJobScriptHistoryResponse
	GetBody() *ListJobScriptHistoryResponseBody
}

type ListJobScriptHistoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobScriptHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobScriptHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobScriptHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListJobScriptHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobScriptHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobScriptHistoryResponse) GetBody() *ListJobScriptHistoryResponseBody {
	return s.Body
}

func (s *ListJobScriptHistoryResponse) SetHeaders(v map[string]*string) *ListJobScriptHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListJobScriptHistoryResponse) SetStatusCode(v int32) *ListJobScriptHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobScriptHistoryResponse) SetBody(v *ListJobScriptHistoryResponseBody) *ListJobScriptHistoryResponse {
	s.Body = v
	return s
}

func (s *ListJobScriptHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
