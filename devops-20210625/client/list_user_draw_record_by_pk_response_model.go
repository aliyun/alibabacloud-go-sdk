// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDrawRecordByPkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserDrawRecordByPkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserDrawRecordByPkResponse
	GetStatusCode() *int32
	SetBody(v *ListUserDrawRecordByPkResponseBody) *ListUserDrawRecordByPkResponse
	GetBody() *ListUserDrawRecordByPkResponseBody
}

type ListUserDrawRecordByPkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDrawRecordByPkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDrawRecordByPkResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserDrawRecordByPkResponse) GoString() string {
	return s.String()
}

func (s *ListUserDrawRecordByPkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserDrawRecordByPkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserDrawRecordByPkResponse) GetBody() *ListUserDrawRecordByPkResponseBody {
	return s.Body
}

func (s *ListUserDrawRecordByPkResponse) SetHeaders(v map[string]*string) *ListUserDrawRecordByPkResponse {
	s.Headers = v
	return s
}

func (s *ListUserDrawRecordByPkResponse) SetStatusCode(v int32) *ListUserDrawRecordByPkResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDrawRecordByPkResponse) SetBody(v *ListUserDrawRecordByPkResponseBody) *ListUserDrawRecordByPkResponse {
	s.Body = v
	return s
}

func (s *ListUserDrawRecordByPkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
