// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseUserIntentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseUserIntentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseUserIntentionResponse
	GetStatusCode() *int32
	SetBody(v *CloseUserIntentionResponseBody) *CloseUserIntentionResponse
	GetBody() *CloseUserIntentionResponseBody
}

type CloseUserIntentionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseUserIntentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseUserIntentionResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseUserIntentionResponse) GoString() string {
	return s.String()
}

func (s *CloseUserIntentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseUserIntentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseUserIntentionResponse) GetBody() *CloseUserIntentionResponseBody {
	return s.Body
}

func (s *CloseUserIntentionResponse) SetHeaders(v map[string]*string) *CloseUserIntentionResponse {
	s.Headers = v
	return s
}

func (s *CloseUserIntentionResponse) SetStatusCode(v int32) *CloseUserIntentionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseUserIntentionResponse) SetBody(v *CloseUserIntentionResponseBody) *CloseUserIntentionResponse {
	s.Body = v
	return s
}

func (s *CloseUserIntentionResponse) Validate() error {
	return dara.Validate(s)
}
