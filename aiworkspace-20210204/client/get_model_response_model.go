// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelResponse
	GetStatusCode() *int32
	SetBody(v *GetModelResponseBody) *GetModelResponse
	GetBody() *GetModelResponseBody
}

type GetModelResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponse) GoString() string {
	return s.String()
}

func (s *GetModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelResponse) GetBody() *GetModelResponseBody {
	return s.Body
}

func (s *GetModelResponse) SetHeaders(v map[string]*string) *GetModelResponse {
	s.Headers = v
	return s
}

func (s *GetModelResponse) SetStatusCode(v int32) *GetModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelResponse) SetBody(v *GetModelResponseBody) *GetModelResponse {
	s.Body = v
	return s
}

func (s *GetModelResponse) Validate() error {
	return dara.Validate(s)
}
