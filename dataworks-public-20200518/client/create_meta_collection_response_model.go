// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMetaCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMetaCollectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateMetaCollectionResponseBody) *CreateMetaCollectionResponse
	GetBody() *CreateMetaCollectionResponseBody
}

type CreateMetaCollectionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMetaCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMetaCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCollectionResponse) GoString() string {
	return s.String()
}

func (s *CreateMetaCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMetaCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMetaCollectionResponse) GetBody() *CreateMetaCollectionResponseBody {
	return s.Body
}

func (s *CreateMetaCollectionResponse) SetHeaders(v map[string]*string) *CreateMetaCollectionResponse {
	s.Headers = v
	return s
}

func (s *CreateMetaCollectionResponse) SetStatusCode(v int32) *CreateMetaCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMetaCollectionResponse) SetBody(v *CreateMetaCollectionResponseBody) *CreateMetaCollectionResponse {
	s.Body = v
	return s
}

func (s *CreateMetaCollectionResponse) Validate() error {
	return dara.Validate(s)
}
