// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformDBInstancePayTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransformDBInstancePayTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransformDBInstancePayTypeResponse
	GetStatusCode() *int32
	SetBody(v *TransformDBInstancePayTypeResponseBody) *TransformDBInstancePayTypeResponse
	GetBody() *TransformDBInstancePayTypeResponseBody
}

type TransformDBInstancePayTypeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransformDBInstancePayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransformDBInstancePayTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s TransformDBInstancePayTypeResponse) GoString() string {
	return s.String()
}

func (s *TransformDBInstancePayTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransformDBInstancePayTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransformDBInstancePayTypeResponse) GetBody() *TransformDBInstancePayTypeResponseBody {
	return s.Body
}

func (s *TransformDBInstancePayTypeResponse) SetHeaders(v map[string]*string) *TransformDBInstancePayTypeResponse {
	s.Headers = v
	return s
}

func (s *TransformDBInstancePayTypeResponse) SetStatusCode(v int32) *TransformDBInstancePayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *TransformDBInstancePayTypeResponse) SetBody(v *TransformDBInstancePayTypeResponseBody) *TransformDBInstancePayTypeResponse {
	s.Body = v
	return s
}

func (s *TransformDBInstancePayTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
