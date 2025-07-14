// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCollectionResponse
	GetStatusCode() *int32
	SetBody(v *CancelCollectionResponseBody) *CancelCollectionResponse
	GetBody() *CancelCollectionResponseBody
}

type CancelCollectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCollectionResponse) GoString() string {
	return s.String()
}

func (s *CancelCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCollectionResponse) GetBody() *CancelCollectionResponseBody {
	return s.Body
}

func (s *CancelCollectionResponse) SetHeaders(v map[string]*string) *CancelCollectionResponse {
	s.Headers = v
	return s
}

func (s *CancelCollectionResponse) SetStatusCode(v int32) *CancelCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCollectionResponse) SetBody(v *CancelCollectionResponseBody) *CancelCollectionResponse {
	s.Body = v
	return s
}

func (s *CancelCollectionResponse) Validate() error {
	return dara.Validate(s)
}
