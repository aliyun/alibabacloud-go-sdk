// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateKgEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateKgEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateKgEntityResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateKgEntityResponseBody) *BatchCreateKgEntityResponse
	GetBody() *BatchCreateKgEntityResponseBody
}

type BatchCreateKgEntityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateKgEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateKgEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgEntityResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateKgEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateKgEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateKgEntityResponse) GetBody() *BatchCreateKgEntityResponseBody {
	return s.Body
}

func (s *BatchCreateKgEntityResponse) SetHeaders(v map[string]*string) *BatchCreateKgEntityResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateKgEntityResponse) SetStatusCode(v int32) *BatchCreateKgEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateKgEntityResponse) SetBody(v *BatchCreateKgEntityResponseBody) *BatchCreateKgEntityResponse {
	s.Body = v
	return s
}

func (s *BatchCreateKgEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
