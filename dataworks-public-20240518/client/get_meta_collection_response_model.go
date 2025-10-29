// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaCollectionResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaCollectionResponseBody) *GetMetaCollectionResponse
	GetBody() *GetMetaCollectionResponseBody
}

type GetMetaCollectionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCollectionResponse) GoString() string {
	return s.String()
}

func (s *GetMetaCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaCollectionResponse) GetBody() *GetMetaCollectionResponseBody {
	return s.Body
}

func (s *GetMetaCollectionResponse) SetHeaders(v map[string]*string) *GetMetaCollectionResponse {
	s.Headers = v
	return s
}

func (s *GetMetaCollectionResponse) SetStatusCode(v int32) *GetMetaCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaCollectionResponse) SetBody(v *GetMetaCollectionResponseBody) *GetMetaCollectionResponse {
	s.Body = v
	return s
}

func (s *GetMetaCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
