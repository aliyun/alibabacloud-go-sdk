// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleSqlDryRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SingleSqlDryRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SingleSqlDryRunResponse
	GetStatusCode() *int32
	SetBody(v *SingleSqlDryRunResponseBody) *SingleSqlDryRunResponse
	GetBody() *SingleSqlDryRunResponseBody
}

type SingleSqlDryRunResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleSqlDryRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SingleSqlDryRunResponse) String() string {
	return dara.Prettify(s)
}

func (s SingleSqlDryRunResponse) GoString() string {
	return s.String()
}

func (s *SingleSqlDryRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SingleSqlDryRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SingleSqlDryRunResponse) GetBody() *SingleSqlDryRunResponseBody {
	return s.Body
}

func (s *SingleSqlDryRunResponse) SetHeaders(v map[string]*string) *SingleSqlDryRunResponse {
	s.Headers = v
	return s
}

func (s *SingleSqlDryRunResponse) SetStatusCode(v int32) *SingleSqlDryRunResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleSqlDryRunResponse) SetBody(v *SingleSqlDryRunResponseBody) *SingleSqlDryRunResponse {
	s.Body = v
	return s
}

func (s *SingleSqlDryRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
