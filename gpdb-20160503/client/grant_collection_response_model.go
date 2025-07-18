// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantCollectionResponse
	GetStatusCode() *int32
	SetBody(v *GrantCollectionResponseBody) *GrantCollectionResponse
	GetBody() *GrantCollectionResponseBody
}

type GrantCollectionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantCollectionResponse) GoString() string {
	return s.String()
}

func (s *GrantCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantCollectionResponse) GetBody() *GrantCollectionResponseBody {
	return s.Body
}

func (s *GrantCollectionResponse) SetHeaders(v map[string]*string) *GrantCollectionResponse {
	s.Headers = v
	return s
}

func (s *GrantCollectionResponse) SetStatusCode(v int32) *GrantCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantCollectionResponse) SetBody(v *GrantCollectionResponseBody) *GrantCollectionResponse {
	s.Body = v
	return s
}

func (s *GrantCollectionResponse) Validate() error {
	return dara.Validate(s)
}
