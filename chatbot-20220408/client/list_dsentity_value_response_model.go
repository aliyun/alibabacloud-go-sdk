// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDSEntityValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDSEntityValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDSEntityValueResponse
	GetStatusCode() *int32
	SetBody(v *ListDSEntityValueResponseBody) *ListDSEntityValueResponse
	GetBody() *ListDSEntityValueResponseBody
}

type ListDSEntityValueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDSEntityValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDSEntityValueResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDSEntityValueResponse) GoString() string {
	return s.String()
}

func (s *ListDSEntityValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDSEntityValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDSEntityValueResponse) GetBody() *ListDSEntityValueResponseBody {
	return s.Body
}

func (s *ListDSEntityValueResponse) SetHeaders(v map[string]*string) *ListDSEntityValueResponse {
	s.Headers = v
	return s
}

func (s *ListDSEntityValueResponse) SetStatusCode(v int32) *ListDSEntityValueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDSEntityValueResponse) SetBody(v *ListDSEntityValueResponseBody) *ListDSEntityValueResponse {
	s.Body = v
	return s
}

func (s *ListDSEntityValueResponse) Validate() error {
	return dara.Validate(s)
}
