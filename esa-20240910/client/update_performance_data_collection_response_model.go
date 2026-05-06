// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePerformanceDataCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePerformanceDataCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePerformanceDataCollectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePerformanceDataCollectionResponseBody) *UpdatePerformanceDataCollectionResponse
	GetBody() *UpdatePerformanceDataCollectionResponseBody
}

type UpdatePerformanceDataCollectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePerformanceDataCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePerformanceDataCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePerformanceDataCollectionResponse) GoString() string {
	return s.String()
}

func (s *UpdatePerformanceDataCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePerformanceDataCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePerformanceDataCollectionResponse) GetBody() *UpdatePerformanceDataCollectionResponseBody {
	return s.Body
}

func (s *UpdatePerformanceDataCollectionResponse) SetHeaders(v map[string]*string) *UpdatePerformanceDataCollectionResponse {
	s.Headers = v
	return s
}

func (s *UpdatePerformanceDataCollectionResponse) SetStatusCode(v int32) *UpdatePerformanceDataCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePerformanceDataCollectionResponse) SetBody(v *UpdatePerformanceDataCollectionResponseBody) *UpdatePerformanceDataCollectionResponse {
	s.Body = v
	return s
}

func (s *UpdatePerformanceDataCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
