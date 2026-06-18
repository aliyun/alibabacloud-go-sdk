// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAssignSeatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchAssignSeatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchAssignSeatsResponse
	GetStatusCode() *int32
	SetBody(v *BatchAssignSeatsResponseBody) *BatchAssignSeatsResponse
	GetBody() *BatchAssignSeatsResponseBody
}

type BatchAssignSeatsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAssignSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAssignSeatsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchAssignSeatsResponse) GoString() string {
	return s.String()
}

func (s *BatchAssignSeatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchAssignSeatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchAssignSeatsResponse) GetBody() *BatchAssignSeatsResponseBody {
	return s.Body
}

func (s *BatchAssignSeatsResponse) SetHeaders(v map[string]*string) *BatchAssignSeatsResponse {
	s.Headers = v
	return s
}

func (s *BatchAssignSeatsResponse) SetStatusCode(v int32) *BatchAssignSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAssignSeatsResponse) SetBody(v *BatchAssignSeatsResponseBody) *BatchAssignSeatsResponse {
	s.Body = v
	return s
}

func (s *BatchAssignSeatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
