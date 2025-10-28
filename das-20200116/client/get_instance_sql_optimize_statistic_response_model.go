// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSqlOptimizeStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceSqlOptimizeStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceSqlOptimizeStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceSqlOptimizeStatisticResponseBody) *GetInstanceSqlOptimizeStatisticResponse
	GetBody() *GetInstanceSqlOptimizeStatisticResponseBody
}

type GetInstanceSqlOptimizeStatisticResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceSqlOptimizeStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceSqlOptimizeStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSqlOptimizeStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceSqlOptimizeStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceSqlOptimizeStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceSqlOptimizeStatisticResponse) GetBody() *GetInstanceSqlOptimizeStatisticResponseBody {
	return s.Body
}

func (s *GetInstanceSqlOptimizeStatisticResponse) SetHeaders(v map[string]*string) *GetInstanceSqlOptimizeStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponse) SetStatusCode(v int32) *GetInstanceSqlOptimizeStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponse) SetBody(v *GetInstanceSqlOptimizeStatisticResponseBody) *GetInstanceSqlOptimizeStatisticResponse {
	s.Body = v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
