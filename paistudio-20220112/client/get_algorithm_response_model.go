// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlgorithmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlgorithmResponse
	GetStatusCode() *int32
	SetBody(v *GetAlgorithmResponseBody) *GetAlgorithmResponse
	GetBody() *GetAlgorithmResponseBody
}

type GetAlgorithmResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlgorithmResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlgorithmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlgorithmResponse) GetBody() *GetAlgorithmResponseBody {
	return s.Body
}

func (s *GetAlgorithmResponse) SetHeaders(v map[string]*string) *GetAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmResponse) SetStatusCode(v int32) *GetAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmResponse) SetBody(v *GetAlgorithmResponseBody) *GetAlgorithmResponse {
	s.Body = v
	return s
}

func (s *GetAlgorithmResponse) Validate() error {
	return dara.Validate(s)
}
