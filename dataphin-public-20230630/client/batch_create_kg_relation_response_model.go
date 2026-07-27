// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateKgRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateKgRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateKgRelationResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateKgRelationResponseBody) *BatchCreateKgRelationResponse
	GetBody() *BatchCreateKgRelationResponseBody
}

type BatchCreateKgRelationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateKgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateKgRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgRelationResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateKgRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateKgRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateKgRelationResponse) GetBody() *BatchCreateKgRelationResponseBody {
	return s.Body
}

func (s *BatchCreateKgRelationResponse) SetHeaders(v map[string]*string) *BatchCreateKgRelationResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateKgRelationResponse) SetStatusCode(v int32) *BatchCreateKgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateKgRelationResponse) SetBody(v *BatchCreateKgRelationResponseBody) *BatchCreateKgRelationResponse {
	s.Body = v
	return s
}

func (s *BatchCreateKgRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
