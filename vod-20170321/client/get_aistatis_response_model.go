// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIStatisResponse
	GetStatusCode() *int32
	SetBody(v *GetAIStatisResponseBody) *GetAIStatisResponse
	GetBody() *GetAIStatisResponseBody
}

type GetAIStatisResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIStatisResponse) GoString() string {
	return s.String()
}

func (s *GetAIStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIStatisResponse) GetBody() *GetAIStatisResponseBody {
	return s.Body
}

func (s *GetAIStatisResponse) SetHeaders(v map[string]*string) *GetAIStatisResponse {
	s.Headers = v
	return s
}

func (s *GetAIStatisResponse) SetStatusCode(v int32) *GetAIStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIStatisResponse) SetBody(v *GetAIStatisResponseBody) *GetAIStatisResponse {
	s.Body = v
	return s
}

func (s *GetAIStatisResponse) Validate() error {
	return dara.Validate(s)
}
