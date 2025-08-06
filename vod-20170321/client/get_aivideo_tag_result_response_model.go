// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIVideoTagResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIVideoTagResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIVideoTagResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAIVideoTagResultResponseBody) *GetAIVideoTagResultResponse
	GetBody() *GetAIVideoTagResultResponseBody
}

type GetAIVideoTagResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIVideoTagResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIVideoTagResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIVideoTagResultResponse) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIVideoTagResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIVideoTagResultResponse) GetBody() *GetAIVideoTagResultResponseBody {
	return s.Body
}

func (s *GetAIVideoTagResultResponse) SetHeaders(v map[string]*string) *GetAIVideoTagResultResponse {
	s.Headers = v
	return s
}

func (s *GetAIVideoTagResultResponse) SetStatusCode(v int32) *GetAIVideoTagResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIVideoTagResultResponse) SetBody(v *GetAIVideoTagResultResponseBody) *GetAIVideoTagResultResponse {
	s.Body = v
	return s
}

func (s *GetAIVideoTagResultResponse) Validate() error {
	return dara.Validate(s)
}
