// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCalculateDBInstanceWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CalculateDBInstanceWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CalculateDBInstanceWeightResponse
	GetStatusCode() *int32
	SetBody(v *CalculateDBInstanceWeightResponseBody) *CalculateDBInstanceWeightResponse
	GetBody() *CalculateDBInstanceWeightResponseBody
}

type CalculateDBInstanceWeightResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CalculateDBInstanceWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CalculateDBInstanceWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s CalculateDBInstanceWeightResponse) GoString() string {
	return s.String()
}

func (s *CalculateDBInstanceWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CalculateDBInstanceWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CalculateDBInstanceWeightResponse) GetBody() *CalculateDBInstanceWeightResponseBody {
	return s.Body
}

func (s *CalculateDBInstanceWeightResponse) SetHeaders(v map[string]*string) *CalculateDBInstanceWeightResponse {
	s.Headers = v
	return s
}

func (s *CalculateDBInstanceWeightResponse) SetStatusCode(v int32) *CalculateDBInstanceWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *CalculateDBInstanceWeightResponse) SetBody(v *CalculateDBInstanceWeightResponseBody) *CalculateDBInstanceWeightResponse {
	s.Body = v
	return s
}

func (s *CalculateDBInstanceWeightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
