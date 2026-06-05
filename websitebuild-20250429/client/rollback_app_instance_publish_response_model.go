// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackAppInstancePublishResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackAppInstancePublishResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackAppInstancePublishResponse
	GetStatusCode() *int32
	SetBody(v *RollbackAppInstancePublishResponseBody) *RollbackAppInstancePublishResponse
	GetBody() *RollbackAppInstancePublishResponseBody
}

type RollbackAppInstancePublishResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackAppInstancePublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackAppInstancePublishResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppInstancePublishResponse) GoString() string {
	return s.String()
}

func (s *RollbackAppInstancePublishResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackAppInstancePublishResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackAppInstancePublishResponse) GetBody() *RollbackAppInstancePublishResponseBody {
	return s.Body
}

func (s *RollbackAppInstancePublishResponse) SetHeaders(v map[string]*string) *RollbackAppInstancePublishResponse {
	s.Headers = v
	return s
}

func (s *RollbackAppInstancePublishResponse) SetStatusCode(v int32) *RollbackAppInstancePublishResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackAppInstancePublishResponse) SetBody(v *RollbackAppInstancePublishResponseBody) *RollbackAppInstancePublishResponse {
	s.Body = v
	return s
}

func (s *RollbackAppInstancePublishResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
