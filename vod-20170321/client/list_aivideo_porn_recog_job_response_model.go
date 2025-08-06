// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoPornRecogJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIVideoPornRecogJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIVideoPornRecogJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIVideoPornRecogJobResponseBody) *ListAIVideoPornRecogJobResponse
	GetBody() *ListAIVideoPornRecogJobResponseBody
}

type ListAIVideoPornRecogJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIVideoPornRecogJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIVideoPornRecogJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoPornRecogJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIVideoPornRecogJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIVideoPornRecogJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIVideoPornRecogJobResponse) GetBody() *ListAIVideoPornRecogJobResponseBody {
	return s.Body
}

func (s *ListAIVideoPornRecogJobResponse) SetHeaders(v map[string]*string) *ListAIVideoPornRecogJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIVideoPornRecogJobResponse) SetStatusCode(v int32) *ListAIVideoPornRecogJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIVideoPornRecogJobResponse) SetBody(v *ListAIVideoPornRecogJobResponseBody) *ListAIVideoPornRecogJobResponse {
	s.Body = v
	return s
}

func (s *ListAIVideoPornRecogJobResponse) Validate() error {
	return dara.Validate(s)
}
