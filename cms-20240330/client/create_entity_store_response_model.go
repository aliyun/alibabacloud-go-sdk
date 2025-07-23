// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEntityStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEntityStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEntityStoreResponse
	GetStatusCode() *int32
	SetBody(v *CreateEntityStoreResponseBody) *CreateEntityStoreResponse
	GetBody() *CreateEntityStoreResponseBody
}

type CreateEntityStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEntityStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEntityStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEntityStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateEntityStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEntityStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEntityStoreResponse) GetBody() *CreateEntityStoreResponseBody {
	return s.Body
}

func (s *CreateEntityStoreResponse) SetHeaders(v map[string]*string) *CreateEntityStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateEntityStoreResponse) SetStatusCode(v int32) *CreateEntityStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEntityStoreResponse) SetBody(v *CreateEntityStoreResponseBody) *CreateEntityStoreResponse {
	s.Body = v
	return s
}

func (s *CreateEntityStoreResponse) Validate() error {
	return dara.Validate(s)
}
