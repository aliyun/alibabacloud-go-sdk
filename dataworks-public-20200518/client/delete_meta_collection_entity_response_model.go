// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCollectionEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMetaCollectionEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMetaCollectionEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMetaCollectionEntityResponseBody) *DeleteMetaCollectionEntityResponse
	GetBody() *DeleteMetaCollectionEntityResponseBody
}

type DeleteMetaCollectionEntityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetaCollectionEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetaCollectionEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCollectionEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetaCollectionEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMetaCollectionEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMetaCollectionEntityResponse) GetBody() *DeleteMetaCollectionEntityResponseBody {
	return s.Body
}

func (s *DeleteMetaCollectionEntityResponse) SetHeaders(v map[string]*string) *DeleteMetaCollectionEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetaCollectionEntityResponse) SetStatusCode(v int32) *DeleteMetaCollectionEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetaCollectionEntityResponse) SetBody(v *DeleteMetaCollectionEntityResponseBody) *DeleteMetaCollectionEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteMetaCollectionEntityResponse) Validate() error {
	return dara.Validate(s)
}
