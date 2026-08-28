// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateHttpApiOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateHttpApiOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateHttpApiOperationResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateHttpApiOperationResponseBody) *BatchUpdateHttpApiOperationResponse
	GetBody() *BatchUpdateHttpApiOperationResponseBody
}

type BatchUpdateHttpApiOperationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateHttpApiOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateHttpApiOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateHttpApiOperationResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateHttpApiOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateHttpApiOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateHttpApiOperationResponse) GetBody() *BatchUpdateHttpApiOperationResponseBody {
	return s.Body
}

func (s *BatchUpdateHttpApiOperationResponse) SetHeaders(v map[string]*string) *BatchUpdateHttpApiOperationResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateHttpApiOperationResponse) SetStatusCode(v int32) *BatchUpdateHttpApiOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateHttpApiOperationResponse) SetBody(v *BatchUpdateHttpApiOperationResponseBody) *BatchUpdateHttpApiOperationResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateHttpApiOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
