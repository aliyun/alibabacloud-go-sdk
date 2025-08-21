// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDictInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDictInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDictInformationResponse
	GetStatusCode() *int32
	SetBody(v *ListDictInformationResponseBody) *ListDictInformationResponse
	GetBody() *ListDictInformationResponseBody
}

type ListDictInformationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDictInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDictInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDictInformationResponse) GoString() string {
	return s.String()
}

func (s *ListDictInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDictInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDictInformationResponse) GetBody() *ListDictInformationResponseBody {
	return s.Body
}

func (s *ListDictInformationResponse) SetHeaders(v map[string]*string) *ListDictInformationResponse {
	s.Headers = v
	return s
}

func (s *ListDictInformationResponse) SetStatusCode(v int32) *ListDictInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDictInformationResponse) SetBody(v *ListDictInformationResponseBody) *ListDictInformationResponse {
	s.Body = v
	return s
}

func (s *ListDictInformationResponse) Validate() error {
	return dara.Validate(s)
}
