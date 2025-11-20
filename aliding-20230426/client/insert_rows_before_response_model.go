// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertRowsBeforeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertRowsBeforeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertRowsBeforeResponse
	GetStatusCode() *int32
	SetBody(v *InsertRowsBeforeResponseBody) *InsertRowsBeforeResponse
	GetBody() *InsertRowsBeforeResponseBody
}

type InsertRowsBeforeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertRowsBeforeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertRowsBeforeResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertRowsBeforeResponse) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertRowsBeforeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertRowsBeforeResponse) GetBody() *InsertRowsBeforeResponseBody {
	return s.Body
}

func (s *InsertRowsBeforeResponse) SetHeaders(v map[string]*string) *InsertRowsBeforeResponse {
	s.Headers = v
	return s
}

func (s *InsertRowsBeforeResponse) SetStatusCode(v int32) *InsertRowsBeforeResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertRowsBeforeResponse) SetBody(v *InsertRowsBeforeResponseBody) *InsertRowsBeforeResponse {
	s.Body = v
	return s
}

func (s *InsertRowsBeforeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
