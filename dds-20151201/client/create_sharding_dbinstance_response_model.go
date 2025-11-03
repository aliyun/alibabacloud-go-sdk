// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShardingDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateShardingDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateShardingDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateShardingDBInstanceResponseBody) *CreateShardingDBInstanceResponse
	GetBody() *CreateShardingDBInstanceResponseBody
}

type CreateShardingDBInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateShardingDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateShardingDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateShardingDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateShardingDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateShardingDBInstanceResponse) GetBody() *CreateShardingDBInstanceResponseBody {
	return s.Body
}

func (s *CreateShardingDBInstanceResponse) SetHeaders(v map[string]*string) *CreateShardingDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateShardingDBInstanceResponse) SetStatusCode(v int32) *CreateShardingDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateShardingDBInstanceResponse) SetBody(v *CreateShardingDBInstanceResponseBody) *CreateShardingDBInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateShardingDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
