// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTriggersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTriggersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTriggersResponse
	GetStatusCode() *int32
	SetBody(v *ListTriggersResponseBody) *ListTriggersResponse
	GetBody() *ListTriggersResponseBody
}

type ListTriggersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTriggersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTriggersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTriggersResponse) GoString() string {
	return s.String()
}

func (s *ListTriggersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTriggersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTriggersResponse) GetBody() *ListTriggersResponseBody {
	return s.Body
}

func (s *ListTriggersResponse) SetHeaders(v map[string]*string) *ListTriggersResponse {
	s.Headers = v
	return s
}

func (s *ListTriggersResponse) SetStatusCode(v int32) *ListTriggersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTriggersResponse) SetBody(v *ListTriggersResponseBody) *ListTriggersResponse {
	s.Body = v
	return s
}

func (s *ListTriggersResponse) Validate() error {
	return dara.Validate(s)
}
