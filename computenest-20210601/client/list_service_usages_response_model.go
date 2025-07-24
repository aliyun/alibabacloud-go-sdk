// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceUsagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceUsagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceUsagesResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceUsagesResponseBody) *ListServiceUsagesResponse
	GetBody() *ListServiceUsagesResponseBody
}

type ListServiceUsagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceUsagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceUsagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceUsagesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceUsagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceUsagesResponse) GetBody() *ListServiceUsagesResponseBody {
	return s.Body
}

func (s *ListServiceUsagesResponse) SetHeaders(v map[string]*string) *ListServiceUsagesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceUsagesResponse) SetStatusCode(v int32) *ListServiceUsagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceUsagesResponse) SetBody(v *ListServiceUsagesResponseBody) *ListServiceUsagesResponse {
	s.Body = v
	return s
}

func (s *ListServiceUsagesResponse) Validate() error {
	return dara.Validate(s)
}
