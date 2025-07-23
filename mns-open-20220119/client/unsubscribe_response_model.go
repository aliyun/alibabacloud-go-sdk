// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnsubscribeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnsubscribeResponse
	GetStatusCode() *int32
	SetBody(v *UnsubscribeResponseBody) *UnsubscribeResponse
	GetBody() *UnsubscribeResponseBody
}

type UnsubscribeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnsubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnsubscribeResponse) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnsubscribeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnsubscribeResponse) GetBody() *UnsubscribeResponseBody {
	return s.Body
}

func (s *UnsubscribeResponse) SetHeaders(v map[string]*string) *UnsubscribeResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeResponse) SetStatusCode(v int32) *UnsubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *UnsubscribeResponse) SetBody(v *UnsubscribeResponseBody) *UnsubscribeResponse {
	s.Body = v
	return s
}

func (s *UnsubscribeResponse) Validate() error {
	return dara.Validate(s)
}
