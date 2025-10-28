// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertApplicationResponse
	GetStatusCode() *int32
	SetBody(v *InsertApplicationResponseBody) *InsertApplicationResponse
	GetBody() *InsertApplicationResponseBody
}

type InsertApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertApplicationResponse) GoString() string {
	return s.String()
}

func (s *InsertApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertApplicationResponse) GetBody() *InsertApplicationResponseBody {
	return s.Body
}

func (s *InsertApplicationResponse) SetHeaders(v map[string]*string) *InsertApplicationResponse {
	s.Headers = v
	return s
}

func (s *InsertApplicationResponse) SetStatusCode(v int32) *InsertApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertApplicationResponse) SetBody(v *InsertApplicationResponseBody) *InsertApplicationResponse {
	s.Body = v
	return s
}

func (s *InsertApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
