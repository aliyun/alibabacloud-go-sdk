// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeDBInstanceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurgeDBInstanceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurgeDBInstanceLogResponse
	GetStatusCode() *int32
	SetBody(v *PurgeDBInstanceLogResponseBody) *PurgeDBInstanceLogResponse
	GetBody() *PurgeDBInstanceLogResponseBody
}

type PurgeDBInstanceLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurgeDBInstanceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurgeDBInstanceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s PurgeDBInstanceLogResponse) GoString() string {
	return s.String()
}

func (s *PurgeDBInstanceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurgeDBInstanceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurgeDBInstanceLogResponse) GetBody() *PurgeDBInstanceLogResponseBody {
	return s.Body
}

func (s *PurgeDBInstanceLogResponse) SetHeaders(v map[string]*string) *PurgeDBInstanceLogResponse {
	s.Headers = v
	return s
}

func (s *PurgeDBInstanceLogResponse) SetStatusCode(v int32) *PurgeDBInstanceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *PurgeDBInstanceLogResponse) SetBody(v *PurgeDBInstanceLogResponseBody) *PurgeDBInstanceLogResponse {
	s.Body = v
	return s
}

func (s *PurgeDBInstanceLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
