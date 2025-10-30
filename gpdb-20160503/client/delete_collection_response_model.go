// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCollectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCollectionResponseBody) *DeleteCollectionResponse
	GetBody() *DeleteCollectionResponseBody
}

type DeleteCollectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCollectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCollectionResponse) GetBody() *DeleteCollectionResponseBody {
	return s.Body
}

func (s *DeleteCollectionResponse) SetHeaders(v map[string]*string) *DeleteCollectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteCollectionResponse) SetStatusCode(v int32) *DeleteCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCollectionResponse) SetBody(v *DeleteCollectionResponseBody) *DeleteCollectionResponse {
	s.Body = v
	return s
}

func (s *DeleteCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
