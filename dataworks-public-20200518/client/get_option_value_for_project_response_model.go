// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOptionValueForProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOptionValueForProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOptionValueForProjectResponse
	GetStatusCode() *int32
	SetBody(v *GetOptionValueForProjectResponseBody) *GetOptionValueForProjectResponse
	GetBody() *GetOptionValueForProjectResponseBody
}

type GetOptionValueForProjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOptionValueForProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOptionValueForProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOptionValueForProjectResponse) GoString() string {
	return s.String()
}

func (s *GetOptionValueForProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOptionValueForProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOptionValueForProjectResponse) GetBody() *GetOptionValueForProjectResponseBody {
	return s.Body
}

func (s *GetOptionValueForProjectResponse) SetHeaders(v map[string]*string) *GetOptionValueForProjectResponse {
	s.Headers = v
	return s
}

func (s *GetOptionValueForProjectResponse) SetStatusCode(v int32) *GetOptionValueForProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOptionValueForProjectResponse) SetBody(v *GetOptionValueForProjectResponseBody) *GetOptionValueForProjectResponse {
	s.Body = v
	return s
}

func (s *GetOptionValueForProjectResponse) Validate() error {
	return dara.Validate(s)
}
