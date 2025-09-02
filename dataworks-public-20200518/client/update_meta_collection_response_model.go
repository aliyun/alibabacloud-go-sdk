// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMetaCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMetaCollectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMetaCollectionResponseBody) *UpdateMetaCollectionResponse
	GetBody() *UpdateMetaCollectionResponseBody
}

type UpdateMetaCollectionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetaCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetaCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCollectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetaCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMetaCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMetaCollectionResponse) GetBody() *UpdateMetaCollectionResponseBody {
	return s.Body
}

func (s *UpdateMetaCollectionResponse) SetHeaders(v map[string]*string) *UpdateMetaCollectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetaCollectionResponse) SetStatusCode(v int32) *UpdateMetaCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetaCollectionResponse) SetBody(v *UpdateMetaCollectionResponseBody) *UpdateMetaCollectionResponse {
	s.Body = v
	return s
}

func (s *UpdateMetaCollectionResponse) Validate() error {
	return dara.Validate(s)
}
