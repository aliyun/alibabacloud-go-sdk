// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableGenerationResponse
	GetStatusCode() *int32
	SetBody(v *GetTableGenerationResponseBody) *GetTableGenerationResponse
	GetBody() *GetTableGenerationResponseBody
}

type GetTableGenerationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableGenerationResponse) GoString() string {
	return s.String()
}

func (s *GetTableGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableGenerationResponse) GetBody() *GetTableGenerationResponseBody {
	return s.Body
}

func (s *GetTableGenerationResponse) SetHeaders(v map[string]*string) *GetTableGenerationResponse {
	s.Headers = v
	return s
}

func (s *GetTableGenerationResponse) SetStatusCode(v int32) *GetTableGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableGenerationResponse) SetBody(v *GetTableGenerationResponseBody) *GetTableGenerationResponse {
	s.Body = v
	return s
}

func (s *GetTableGenerationResponse) Validate() error {
	return dara.Validate(s)
}
