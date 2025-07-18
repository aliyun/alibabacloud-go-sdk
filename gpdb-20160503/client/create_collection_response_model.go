// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCollectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateCollectionResponseBody) *CreateCollectionResponse
	GetBody() *CreateCollectionResponseBody
}

type CreateCollectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectionResponse) GoString() string {
	return s.String()
}

func (s *CreateCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCollectionResponse) GetBody() *CreateCollectionResponseBody {
	return s.Body
}

func (s *CreateCollectionResponse) SetHeaders(v map[string]*string) *CreateCollectionResponse {
	s.Headers = v
	return s
}

func (s *CreateCollectionResponse) SetStatusCode(v int32) *CreateCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCollectionResponse) SetBody(v *CreateCollectionResponseBody) *CreateCollectionResponse {
	s.Body = v
	return s
}

func (s *CreateCollectionResponse) Validate() error {
	return dara.Validate(s)
}
