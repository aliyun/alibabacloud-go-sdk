// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteMetaEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteMetaEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteMetaEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteMetaEntitiesResponseBody) *BatchDeleteMetaEntitiesResponse
	GetBody() *BatchDeleteMetaEntitiesResponseBody
}

type BatchDeleteMetaEntitiesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteMetaEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteMetaEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteMetaEntitiesResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteMetaEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteMetaEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteMetaEntitiesResponse) GetBody() *BatchDeleteMetaEntitiesResponseBody {
	return s.Body
}

func (s *BatchDeleteMetaEntitiesResponse) SetHeaders(v map[string]*string) *BatchDeleteMetaEntitiesResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteMetaEntitiesResponse) SetStatusCode(v int32) *BatchDeleteMetaEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteMetaEntitiesResponse) SetBody(v *BatchDeleteMetaEntitiesResponseBody) *BatchDeleteMetaEntitiesResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteMetaEntitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
