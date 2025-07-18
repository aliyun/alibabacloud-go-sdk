// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCollectionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCollectionResponseBody) *ModifyCollectionResponse
	GetBody() *ModifyCollectionResponseBody
}

type ModifyCollectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCollectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCollectionResponse) GetBody() *ModifyCollectionResponseBody {
	return s.Body
}

func (s *ModifyCollectionResponse) SetHeaders(v map[string]*string) *ModifyCollectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyCollectionResponse) SetStatusCode(v int32) *ModifyCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCollectionResponse) SetBody(v *ModifyCollectionResponseBody) *ModifyCollectionResponse {
	s.Body = v
	return s
}

func (s *ModifyCollectionResponse) Validate() error {
	return dara.Validate(s)
}
