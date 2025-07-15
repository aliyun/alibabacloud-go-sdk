// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlacklistCallTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBlacklistCallTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBlacklistCallTaggingResponse
	GetStatusCode() *int32
	SetBody(v *AddBlacklistCallTaggingResponseBody) *AddBlacklistCallTaggingResponse
	GetBody() *AddBlacklistCallTaggingResponseBody
}

type AddBlacklistCallTaggingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBlacklistCallTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBlacklistCallTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBlacklistCallTaggingResponse) GoString() string {
	return s.String()
}

func (s *AddBlacklistCallTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBlacklistCallTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBlacklistCallTaggingResponse) GetBody() *AddBlacklistCallTaggingResponseBody {
	return s.Body
}

func (s *AddBlacklistCallTaggingResponse) SetHeaders(v map[string]*string) *AddBlacklistCallTaggingResponse {
	s.Headers = v
	return s
}

func (s *AddBlacklistCallTaggingResponse) SetStatusCode(v int32) *AddBlacklistCallTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBlacklistCallTaggingResponse) SetBody(v *AddBlacklistCallTaggingResponseBody) *AddBlacklistCallTaggingResponse {
	s.Body = v
	return s
}

func (s *AddBlacklistCallTaggingResponse) Validate() error {
	return dara.Validate(s)
}
