// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackConfigModificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackConfigModificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackConfigModificationResponse
	GetStatusCode() *int32
	SetBody(v *RollbackConfigModificationResponseBody) *RollbackConfigModificationResponse
	GetBody() *RollbackConfigModificationResponseBody
}

type RollbackConfigModificationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackConfigModificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackConfigModificationResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackConfigModificationResponse) GoString() string {
	return s.String()
}

func (s *RollbackConfigModificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackConfigModificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackConfigModificationResponse) GetBody() *RollbackConfigModificationResponseBody {
	return s.Body
}

func (s *RollbackConfigModificationResponse) SetHeaders(v map[string]*string) *RollbackConfigModificationResponse {
	s.Headers = v
	return s
}

func (s *RollbackConfigModificationResponse) SetStatusCode(v int32) *RollbackConfigModificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackConfigModificationResponse) SetBody(v *RollbackConfigModificationResponseBody) *RollbackConfigModificationResponse {
	s.Body = v
	return s
}

func (s *RollbackConfigModificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
