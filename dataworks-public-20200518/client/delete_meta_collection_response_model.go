// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMetaCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMetaCollectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMetaCollectionResponseBody) *DeleteMetaCollectionResponse
	GetBody() *DeleteMetaCollectionResponseBody
}

type DeleteMetaCollectionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetaCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetaCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCollectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetaCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMetaCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMetaCollectionResponse) GetBody() *DeleteMetaCollectionResponseBody {
	return s.Body
}

func (s *DeleteMetaCollectionResponse) SetHeaders(v map[string]*string) *DeleteMetaCollectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetaCollectionResponse) SetStatusCode(v int32) *DeleteMetaCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetaCollectionResponse) SetBody(v *DeleteMetaCollectionResponseBody) *DeleteMetaCollectionResponse {
	s.Body = v
	return s
}

func (s *DeleteMetaCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
