// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMetaCollectionEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMetaCollectionEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMetaCollectionEntityResponse
	GetStatusCode() *int32
	SetBody(v *AddMetaCollectionEntityResponseBody) *AddMetaCollectionEntityResponse
	GetBody() *AddMetaCollectionEntityResponseBody
}

type AddMetaCollectionEntityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMetaCollectionEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMetaCollectionEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMetaCollectionEntityResponse) GoString() string {
	return s.String()
}

func (s *AddMetaCollectionEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMetaCollectionEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMetaCollectionEntityResponse) GetBody() *AddMetaCollectionEntityResponseBody {
	return s.Body
}

func (s *AddMetaCollectionEntityResponse) SetHeaders(v map[string]*string) *AddMetaCollectionEntityResponse {
	s.Headers = v
	return s
}

func (s *AddMetaCollectionEntityResponse) SetStatusCode(v int32) *AddMetaCollectionEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMetaCollectionEntityResponse) SetBody(v *AddMetaCollectionEntityResponseBody) *AddMetaCollectionEntityResponse {
	s.Body = v
	return s
}

func (s *AddMetaCollectionEntityResponse) Validate() error {
	return dara.Validate(s)
}
