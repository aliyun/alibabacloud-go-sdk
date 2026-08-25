// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteModelsResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteModelsResponseBody) *BatchDeleteModelsResponse
	GetBody() *BatchDeleteModelsResponseBody
}

type BatchDeleteModelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteModelsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteModelsResponse) GetBody() *BatchDeleteModelsResponseBody {
	return s.Body
}

func (s *BatchDeleteModelsResponse) SetHeaders(v map[string]*string) *BatchDeleteModelsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteModelsResponse) SetStatusCode(v int32) *BatchDeleteModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteModelsResponse) SetBody(v *BatchDeleteModelsResponseBody) *BatchDeleteModelsResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
