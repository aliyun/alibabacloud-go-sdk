// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncEcsHostTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSyncEcsHostTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSyncEcsHostTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSyncEcsHostTaskResponseBody) *DescribeSyncEcsHostTaskResponse
	GetBody() *DescribeSyncEcsHostTaskResponseBody
}

type DescribeSyncEcsHostTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSyncEcsHostTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSyncEcsHostTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSyncEcsHostTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSyncEcsHostTaskResponse) GetBody() *DescribeSyncEcsHostTaskResponseBody {
	return s.Body
}

func (s *DescribeSyncEcsHostTaskResponse) SetHeaders(v map[string]*string) *DescribeSyncEcsHostTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponse) SetStatusCode(v int32) *DescribeSyncEcsHostTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponse) SetBody(v *DescribeSyncEcsHostTaskResponseBody) *DescribeSyncEcsHostTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
