// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareSampleConsistencyJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompareSampleConsistencyJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompareSampleConsistencyJobResponse
	GetStatusCode() *int32
	SetBody(v *CompareSampleConsistencyJobResponseBody) *CompareSampleConsistencyJobResponse
	GetBody() *CompareSampleConsistencyJobResponseBody
}

type CompareSampleConsistencyJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareSampleConsistencyJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareSampleConsistencyJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CompareSampleConsistencyJobResponse) GoString() string {
	return s.String()
}

func (s *CompareSampleConsistencyJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompareSampleConsistencyJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompareSampleConsistencyJobResponse) GetBody() *CompareSampleConsistencyJobResponseBody {
	return s.Body
}

func (s *CompareSampleConsistencyJobResponse) SetHeaders(v map[string]*string) *CompareSampleConsistencyJobResponse {
	s.Headers = v
	return s
}

func (s *CompareSampleConsistencyJobResponse) SetStatusCode(v int32) *CompareSampleConsistencyJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareSampleConsistencyJobResponse) SetBody(v *CompareSampleConsistencyJobResponseBody) *CompareSampleConsistencyJobResponse {
	s.Body = v
	return s
}

func (s *CompareSampleConsistencyJobResponse) Validate() error {
	return dara.Validate(s)
}
