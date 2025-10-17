// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterStoragePerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterStoragePerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterStoragePerformanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterStoragePerformanceResponseBody) *ModifyDBClusterStoragePerformanceResponse
	GetBody() *ModifyDBClusterStoragePerformanceResponseBody
}

type ModifyDBClusterStoragePerformanceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterStoragePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterStoragePerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterStoragePerformanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterStoragePerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterStoragePerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterStoragePerformanceResponse) GetBody() *ModifyDBClusterStoragePerformanceResponseBody {
	return s.Body
}

func (s *ModifyDBClusterStoragePerformanceResponse) SetHeaders(v map[string]*string) *ModifyDBClusterStoragePerformanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterStoragePerformanceResponse) SetStatusCode(v int32) *ModifyDBClusterStoragePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceResponse) SetBody(v *ModifyDBClusterStoragePerformanceResponseBody) *ModifyDBClusterStoragePerformanceResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterStoragePerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
