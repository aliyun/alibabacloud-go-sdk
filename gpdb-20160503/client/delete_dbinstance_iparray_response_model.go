// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceIPArrayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBInstanceIPArrayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBInstanceIPArrayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBInstanceIPArrayResponseBody) *DeleteDBInstanceIPArrayResponse
	GetBody() *DeleteDBInstanceIPArrayResponseBody
}

type DeleteDBInstanceIPArrayResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBInstanceIPArrayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBInstanceIPArrayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceIPArrayResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceIPArrayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBInstanceIPArrayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBInstanceIPArrayResponse) GetBody() *DeleteDBInstanceIPArrayResponseBody {
	return s.Body
}

func (s *DeleteDBInstanceIPArrayResponse) SetHeaders(v map[string]*string) *DeleteDBInstanceIPArrayResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstanceIPArrayResponse) SetStatusCode(v int32) *DeleteDBInstanceIPArrayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstanceIPArrayResponse) SetBody(v *DeleteDBInstanceIPArrayResponseBody) *DeleteDBInstanceIPArrayResponse {
	s.Body = v
	return s
}

func (s *DeleteDBInstanceIPArrayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
