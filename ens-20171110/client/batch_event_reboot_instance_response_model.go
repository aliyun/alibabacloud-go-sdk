// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventRebootInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchEventRebootInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchEventRebootInstanceResponse
	GetStatusCode() *int32
	SetBody(v *BatchEventRebootInstanceResponseBody) *BatchEventRebootInstanceResponse
	GetBody() *BatchEventRebootInstanceResponseBody
}

type BatchEventRebootInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchEventRebootInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchEventRebootInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRebootInstanceResponse) GoString() string {
	return s.String()
}

func (s *BatchEventRebootInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchEventRebootInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchEventRebootInstanceResponse) GetBody() *BatchEventRebootInstanceResponseBody {
	return s.Body
}

func (s *BatchEventRebootInstanceResponse) SetHeaders(v map[string]*string) *BatchEventRebootInstanceResponse {
	s.Headers = v
	return s
}

func (s *BatchEventRebootInstanceResponse) SetStatusCode(v int32) *BatchEventRebootInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchEventRebootInstanceResponse) SetBody(v *BatchEventRebootInstanceResponseBody) *BatchEventRebootInstanceResponse {
	s.Body = v
	return s
}

func (s *BatchEventRebootInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
