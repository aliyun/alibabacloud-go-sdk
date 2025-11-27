// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBInstanceTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDBInstanceTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDBInstanceTopologyResponse
	GetStatusCode() *int32
	SetBody(v *GetDBInstanceTopologyResponseBody) *GetDBInstanceTopologyResponse
	GetBody() *GetDBInstanceTopologyResponseBody
}

type GetDBInstanceTopologyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDBInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDBInstanceTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDBInstanceTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDBInstanceTopologyResponse) GetBody() *GetDBInstanceTopologyResponseBody {
	return s.Body
}

func (s *GetDBInstanceTopologyResponse) SetHeaders(v map[string]*string) *GetDBInstanceTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetDBInstanceTopologyResponse) SetStatusCode(v int32) *GetDBInstanceTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDBInstanceTopologyResponse) SetBody(v *GetDBInstanceTopologyResponseBody) *GetDBInstanceTopologyResponse {
	s.Body = v
	return s
}

func (s *GetDBInstanceTopologyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
