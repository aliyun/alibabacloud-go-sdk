// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBucketScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateBucketScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateBucketScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *OperateBucketScanTaskResponseBody) *OperateBucketScanTaskResponse
	GetBody() *OperateBucketScanTaskResponseBody
}

type OperateBucketScanTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateBucketScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateBucketScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateBucketScanTaskResponse) GoString() string {
	return s.String()
}

func (s *OperateBucketScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateBucketScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateBucketScanTaskResponse) GetBody() *OperateBucketScanTaskResponseBody {
	return s.Body
}

func (s *OperateBucketScanTaskResponse) SetHeaders(v map[string]*string) *OperateBucketScanTaskResponse {
	s.Headers = v
	return s
}

func (s *OperateBucketScanTaskResponse) SetStatusCode(v int32) *OperateBucketScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateBucketScanTaskResponse) SetBody(v *OperateBucketScanTaskResponseBody) *OperateBucketScanTaskResponse {
	s.Body = v
	return s
}

func (s *OperateBucketScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
