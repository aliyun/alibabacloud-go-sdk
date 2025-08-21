// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuouslyPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContinuouslyPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContinuouslyPushResponse
	GetStatusCode() *int32
	SetBody(v *ContinuouslyPushResponseBody) *ContinuouslyPushResponse
	GetBody() *ContinuouslyPushResponseBody
}

type ContinuouslyPushResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinuouslyPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinuouslyPushResponse) String() string {
	return dara.Prettify(s)
}

func (s ContinuouslyPushResponse) GoString() string {
	return s.String()
}

func (s *ContinuouslyPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContinuouslyPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContinuouslyPushResponse) GetBody() *ContinuouslyPushResponseBody {
	return s.Body
}

func (s *ContinuouslyPushResponse) SetHeaders(v map[string]*string) *ContinuouslyPushResponse {
	s.Headers = v
	return s
}

func (s *ContinuouslyPushResponse) SetStatusCode(v int32) *ContinuouslyPushResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinuouslyPushResponse) SetBody(v *ContinuouslyPushResponseBody) *ContinuouslyPushResponse {
	s.Body = v
	return s
}

func (s *ContinuouslyPushResponse) Validate() error {
	return dara.Validate(s)
}
