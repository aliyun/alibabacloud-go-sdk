// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMessageContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMessageContactResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMessageContactResponseBody) *UpdateMessageContactResponse
	GetBody() *UpdateMessageContactResponseBody
}

type UpdateMessageContactResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMessageContactResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageContactResponse) GoString() string {
	return s.String()
}

func (s *UpdateMessageContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMessageContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMessageContactResponse) GetBody() *UpdateMessageContactResponseBody {
	return s.Body
}

func (s *UpdateMessageContactResponse) SetHeaders(v map[string]*string) *UpdateMessageContactResponse {
	s.Headers = v
	return s
}

func (s *UpdateMessageContactResponse) SetStatusCode(v int32) *UpdateMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageContactResponse) SetBody(v *UpdateMessageContactResponseBody) *UpdateMessageContactResponse {
	s.Body = v
	return s
}

func (s *UpdateMessageContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
