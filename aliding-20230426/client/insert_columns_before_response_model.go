// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertColumnsBeforeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertColumnsBeforeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertColumnsBeforeResponse
	GetStatusCode() *int32
	SetBody(v *InsertColumnsBeforeResponseBody) *InsertColumnsBeforeResponse
	GetBody() *InsertColumnsBeforeResponseBody
}

type InsertColumnsBeforeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertColumnsBeforeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertColumnsBeforeResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertColumnsBeforeResponse) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertColumnsBeforeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertColumnsBeforeResponse) GetBody() *InsertColumnsBeforeResponseBody {
	return s.Body
}

func (s *InsertColumnsBeforeResponse) SetHeaders(v map[string]*string) *InsertColumnsBeforeResponse {
	s.Headers = v
	return s
}

func (s *InsertColumnsBeforeResponse) SetStatusCode(v int32) *InsertColumnsBeforeResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertColumnsBeforeResponse) SetBody(v *InsertColumnsBeforeResponseBody) *InsertColumnsBeforeResponse {
	s.Body = v
	return s
}

func (s *InsertColumnsBeforeResponse) Validate() error {
	return dara.Validate(s)
}
