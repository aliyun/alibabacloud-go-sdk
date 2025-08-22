// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnSubTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDcdnSubTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDcdnSubTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDcdnSubTaskResponseBody) *CreateDcdnSubTaskResponse
	GetBody() *CreateDcdnSubTaskResponseBody
}

type CreateDcdnSubTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDcdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDcdnSubTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnSubTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDcdnSubTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDcdnSubTaskResponse) GetBody() *CreateDcdnSubTaskResponseBody {
	return s.Body
}

func (s *CreateDcdnSubTaskResponse) SetHeaders(v map[string]*string) *CreateDcdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnSubTaskResponse) SetStatusCode(v int32) *CreateDcdnSubTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDcdnSubTaskResponse) SetBody(v *CreateDcdnSubTaskResponseBody) *CreateDcdnSubTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDcdnSubTaskResponse) Validate() error {
	return dara.Validate(s)
}
