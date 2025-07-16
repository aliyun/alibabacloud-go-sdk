// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertContentWithOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertContentWithOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertContentWithOptionsResponse
	GetStatusCode() *int32
	SetBody(v *InsertContentWithOptionsResponseBody) *InsertContentWithOptionsResponse
	GetBody() *InsertContentWithOptionsResponseBody
}

type InsertContentWithOptionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertContentWithOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertContentWithOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertContentWithOptionsResponse) GoString() string {
	return s.String()
}

func (s *InsertContentWithOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertContentWithOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertContentWithOptionsResponse) GetBody() *InsertContentWithOptionsResponseBody {
	return s.Body
}

func (s *InsertContentWithOptionsResponse) SetHeaders(v map[string]*string) *InsertContentWithOptionsResponse {
	s.Headers = v
	return s
}

func (s *InsertContentWithOptionsResponse) SetStatusCode(v int32) *InsertContentWithOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertContentWithOptionsResponse) SetBody(v *InsertContentWithOptionsResponseBody) *InsertContentWithOptionsResponse {
	s.Body = v
	return s
}

func (s *InsertContentWithOptionsResponse) Validate() error {
	return dara.Validate(s)
}
