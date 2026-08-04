// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncrByCacheOperateSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IncrByCacheOperateSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IncrByCacheOperateSyncResponse
	GetStatusCode() *int32
	SetBody(v *IncrByCacheOperateSyncResponseBody) *IncrByCacheOperateSyncResponse
	GetBody() *IncrByCacheOperateSyncResponseBody
}

type IncrByCacheOperateSyncResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IncrByCacheOperateSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IncrByCacheOperateSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s IncrByCacheOperateSyncResponse) GoString() string {
	return s.String()
}

func (s *IncrByCacheOperateSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IncrByCacheOperateSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IncrByCacheOperateSyncResponse) GetBody() *IncrByCacheOperateSyncResponseBody {
	return s.Body
}

func (s *IncrByCacheOperateSyncResponse) SetHeaders(v map[string]*string) *IncrByCacheOperateSyncResponse {
	s.Headers = v
	return s
}

func (s *IncrByCacheOperateSyncResponse) SetStatusCode(v int32) *IncrByCacheOperateSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *IncrByCacheOperateSyncResponse) SetBody(v *IncrByCacheOperateSyncResponseBody) *IncrByCacheOperateSyncResponse {
	s.Body = v
	return s
}

func (s *IncrByCacheOperateSyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
