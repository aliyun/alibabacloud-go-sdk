// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMessageContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMessageContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMessageContactResponse
	GetStatusCode() *int32
	SetBody(v *AddMessageContactResponseBody) *AddMessageContactResponse
	GetBody() *AddMessageContactResponseBody
}

type AddMessageContactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMessageContactResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMessageContactResponse) GoString() string {
	return s.String()
}

func (s *AddMessageContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMessageContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMessageContactResponse) GetBody() *AddMessageContactResponseBody {
	return s.Body
}

func (s *AddMessageContactResponse) SetHeaders(v map[string]*string) *AddMessageContactResponse {
	s.Headers = v
	return s
}

func (s *AddMessageContactResponse) SetStatusCode(v int32) *AddMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMessageContactResponse) SetBody(v *AddMessageContactResponseBody) *AddMessageContactResponse {
	s.Body = v
	return s
}

func (s *AddMessageContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
