// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBlacklistCallTaggingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBlacklistCallTaggingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBlacklistCallTaggingsResponse
	GetStatusCode() *int32
	SetBody(v *ListBlacklistCallTaggingsResponseBody) *ListBlacklistCallTaggingsResponse
	GetBody() *ListBlacklistCallTaggingsResponseBody
}

type ListBlacklistCallTaggingsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBlacklistCallTaggingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBlacklistCallTaggingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBlacklistCallTaggingsResponse) GoString() string {
	return s.String()
}

func (s *ListBlacklistCallTaggingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBlacklistCallTaggingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBlacklistCallTaggingsResponse) GetBody() *ListBlacklistCallTaggingsResponseBody {
	return s.Body
}

func (s *ListBlacklistCallTaggingsResponse) SetHeaders(v map[string]*string) *ListBlacklistCallTaggingsResponse {
	s.Headers = v
	return s
}

func (s *ListBlacklistCallTaggingsResponse) SetStatusCode(v int32) *ListBlacklistCallTaggingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBlacklistCallTaggingsResponse) SetBody(v *ListBlacklistCallTaggingsResponseBody) *ListBlacklistCallTaggingsResponse {
	s.Body = v
	return s
}

func (s *ListBlacklistCallTaggingsResponse) Validate() error {
	return dara.Validate(s)
}
