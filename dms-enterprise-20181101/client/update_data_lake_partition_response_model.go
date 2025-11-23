// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakePartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataLakePartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataLakePartitionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataLakePartitionResponseBody) *UpdateDataLakePartitionResponse
	GetBody() *UpdateDataLakePartitionResponseBody
}

type UpdateDataLakePartitionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataLakePartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataLakePartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakePartitionResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataLakePartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataLakePartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataLakePartitionResponse) GetBody() *UpdateDataLakePartitionResponseBody {
	return s.Body
}

func (s *UpdateDataLakePartitionResponse) SetHeaders(v map[string]*string) *UpdateDataLakePartitionResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataLakePartitionResponse) SetStatusCode(v int32) *UpdateDataLakePartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataLakePartitionResponse) SetBody(v *UpdateDataLakePartitionResponseBody) *UpdateDataLakePartitionResponse {
	s.Body = v
	return s
}

func (s *UpdateDataLakePartitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
