// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIStatisTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIStatisTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIStatisTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListAIStatisTypeResponseBody) *ListAIStatisTypeResponse
	GetBody() *ListAIStatisTypeResponseBody
}

type ListAIStatisTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIStatisTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIStatisTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIStatisTypeResponse) GoString() string {
	return s.String()
}

func (s *ListAIStatisTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIStatisTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIStatisTypeResponse) GetBody() *ListAIStatisTypeResponseBody {
	return s.Body
}

func (s *ListAIStatisTypeResponse) SetHeaders(v map[string]*string) *ListAIStatisTypeResponse {
	s.Headers = v
	return s
}

func (s *ListAIStatisTypeResponse) SetStatusCode(v int32) *ListAIStatisTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIStatisTypeResponse) SetBody(v *ListAIStatisTypeResponseBody) *ListAIStatisTypeResponse {
	s.Body = v
	return s
}

func (s *ListAIStatisTypeResponse) Validate() error {
	return dara.Validate(s)
}
