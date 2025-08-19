// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageContactResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageContactResponseBody) *GetMessageContactResponse
	GetBody() *GetMessageContactResponseBody
}

type GetMessageContactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageContactResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageContactResponse) GoString() string {
	return s.String()
}

func (s *GetMessageContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageContactResponse) GetBody() *GetMessageContactResponseBody {
	return s.Body
}

func (s *GetMessageContactResponse) SetHeaders(v map[string]*string) *GetMessageContactResponse {
	s.Headers = v
	return s
}

func (s *GetMessageContactResponse) SetStatusCode(v int32) *GetMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageContactResponse) SetBody(v *GetMessageContactResponseBody) *GetMessageContactResponse {
	s.Body = v
	return s
}

func (s *GetMessageContactResponse) Validate() error {
	return dara.Validate(s)
}
