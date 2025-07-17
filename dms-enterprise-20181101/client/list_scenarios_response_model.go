// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScenariosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScenariosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScenariosResponse
	GetStatusCode() *int32
	SetBody(v *ListScenariosResponseBody) *ListScenariosResponse
	GetBody() *ListScenariosResponseBody
}

type ListScenariosResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScenariosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScenariosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScenariosResponse) GoString() string {
	return s.String()
}

func (s *ListScenariosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScenariosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScenariosResponse) GetBody() *ListScenariosResponseBody {
	return s.Body
}

func (s *ListScenariosResponse) SetHeaders(v map[string]*string) *ListScenariosResponse {
	s.Headers = v
	return s
}

func (s *ListScenariosResponse) SetStatusCode(v int32) *ListScenariosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScenariosResponse) SetBody(v *ListScenariosResponseBody) *ListScenariosResponse {
	s.Body = v
	return s
}

func (s *ListScenariosResponse) Validate() error {
	return dara.Validate(s)
}
