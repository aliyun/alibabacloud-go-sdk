// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDagComplementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDagComplementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDagComplementResponse
	GetStatusCode() *int32
	SetBody(v *CreateDagComplementResponseBody) *CreateDagComplementResponse
	GetBody() *CreateDagComplementResponseBody
}

type CreateDagComplementResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDagComplementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDagComplementResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDagComplementResponse) GoString() string {
	return s.String()
}

func (s *CreateDagComplementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDagComplementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDagComplementResponse) GetBody() *CreateDagComplementResponseBody {
	return s.Body
}

func (s *CreateDagComplementResponse) SetHeaders(v map[string]*string) *CreateDagComplementResponse {
	s.Headers = v
	return s
}

func (s *CreateDagComplementResponse) SetStatusCode(v int32) *CreateDagComplementResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDagComplementResponse) SetBody(v *CreateDagComplementResponseBody) *CreateDagComplementResponse {
	s.Body = v
	return s
}

func (s *CreateDagComplementResponse) Validate() error {
	return dara.Validate(s)
}
