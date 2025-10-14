// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataBatchIngestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataBatchIngestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataBatchIngestionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataBatchIngestionResponseBody) *UpdateDataBatchIngestionResponse
	GetBody() *UpdateDataBatchIngestionResponseBody
}

type UpdateDataBatchIngestionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataBatchIngestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataBatchIngestionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataBatchIngestionResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataBatchIngestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataBatchIngestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataBatchIngestionResponse) GetBody() *UpdateDataBatchIngestionResponseBody {
	return s.Body
}

func (s *UpdateDataBatchIngestionResponse) SetHeaders(v map[string]*string) *UpdateDataBatchIngestionResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataBatchIngestionResponse) SetStatusCode(v int32) *UpdateDataBatchIngestionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataBatchIngestionResponse) SetBody(v *UpdateDataBatchIngestionResponseBody) *UpdateDataBatchIngestionResponse {
	s.Body = v
	return s
}

func (s *UpdateDataBatchIngestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
