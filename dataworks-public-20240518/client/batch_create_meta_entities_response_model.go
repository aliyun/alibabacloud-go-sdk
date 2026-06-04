// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateMetaEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateMetaEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateMetaEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateMetaEntitiesResponseBody) *BatchCreateMetaEntitiesResponse
	GetBody() *BatchCreateMetaEntitiesResponseBody
}

type BatchCreateMetaEntitiesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateMetaEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateMetaEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateMetaEntitiesResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateMetaEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateMetaEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateMetaEntitiesResponse) GetBody() *BatchCreateMetaEntitiesResponseBody {
	return s.Body
}

func (s *BatchCreateMetaEntitiesResponse) SetHeaders(v map[string]*string) *BatchCreateMetaEntitiesResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateMetaEntitiesResponse) SetStatusCode(v int32) *BatchCreateMetaEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateMetaEntitiesResponse) SetBody(v *BatchCreateMetaEntitiesResponseBody) *BatchCreateMetaEntitiesResponse {
	s.Body = v
	return s
}

func (s *BatchCreateMetaEntitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
