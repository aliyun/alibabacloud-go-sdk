// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEntityIntoMetaCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddEntityIntoMetaCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddEntityIntoMetaCollectionResponse
	GetStatusCode() *int32
	SetBody(v *AddEntityIntoMetaCollectionResponseBody) *AddEntityIntoMetaCollectionResponse
	GetBody() *AddEntityIntoMetaCollectionResponseBody
}

type AddEntityIntoMetaCollectionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEntityIntoMetaCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEntityIntoMetaCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddEntityIntoMetaCollectionResponse) GoString() string {
	return s.String()
}

func (s *AddEntityIntoMetaCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddEntityIntoMetaCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddEntityIntoMetaCollectionResponse) GetBody() *AddEntityIntoMetaCollectionResponseBody {
	return s.Body
}

func (s *AddEntityIntoMetaCollectionResponse) SetHeaders(v map[string]*string) *AddEntityIntoMetaCollectionResponse {
	s.Headers = v
	return s
}

func (s *AddEntityIntoMetaCollectionResponse) SetStatusCode(v int32) *AddEntityIntoMetaCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEntityIntoMetaCollectionResponse) SetBody(v *AddEntityIntoMetaCollectionResponseBody) *AddEntityIntoMetaCollectionResponse {
	s.Body = v
	return s
}

func (s *AddEntityIntoMetaCollectionResponse) Validate() error {
	return dara.Validate(s)
}
