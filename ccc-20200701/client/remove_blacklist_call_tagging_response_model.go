// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBlacklistCallTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveBlacklistCallTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveBlacklistCallTaggingResponse
	GetStatusCode() *int32
	SetBody(v *RemoveBlacklistCallTaggingResponseBody) *RemoveBlacklistCallTaggingResponse
	GetBody() *RemoveBlacklistCallTaggingResponseBody
}

type RemoveBlacklistCallTaggingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveBlacklistCallTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveBlacklistCallTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveBlacklistCallTaggingResponse) GoString() string {
	return s.String()
}

func (s *RemoveBlacklistCallTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveBlacklistCallTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveBlacklistCallTaggingResponse) GetBody() *RemoveBlacklistCallTaggingResponseBody {
	return s.Body
}

func (s *RemoveBlacklistCallTaggingResponse) SetHeaders(v map[string]*string) *RemoveBlacklistCallTaggingResponse {
	s.Headers = v
	return s
}

func (s *RemoveBlacklistCallTaggingResponse) SetStatusCode(v int32) *RemoveBlacklistCallTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveBlacklistCallTaggingResponse) SetBody(v *RemoveBlacklistCallTaggingResponseBody) *RemoveBlacklistCallTaggingResponse {
	s.Body = v
	return s
}

func (s *RemoveBlacklistCallTaggingResponse) Validate() error {
	return dara.Validate(s)
}
