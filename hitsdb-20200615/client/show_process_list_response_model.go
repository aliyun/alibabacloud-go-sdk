// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShowProcessListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ShowProcessListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ShowProcessListResponse
	GetStatusCode() *int32
	SetBody(v *ShowProcessListResponseBody) *ShowProcessListResponse
	GetBody() *ShowProcessListResponseBody
}

type ShowProcessListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShowProcessListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShowProcessListResponse) String() string {
	return dara.Prettify(s)
}

func (s ShowProcessListResponse) GoString() string {
	return s.String()
}

func (s *ShowProcessListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ShowProcessListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ShowProcessListResponse) GetBody() *ShowProcessListResponseBody {
	return s.Body
}

func (s *ShowProcessListResponse) SetHeaders(v map[string]*string) *ShowProcessListResponse {
	s.Headers = v
	return s
}

func (s *ShowProcessListResponse) SetStatusCode(v int32) *ShowProcessListResponse {
	s.StatusCode = &v
	return s
}

func (s *ShowProcessListResponse) SetBody(v *ShowProcessListResponseBody) *ShowProcessListResponse {
	s.Body = v
	return s
}

func (s *ShowProcessListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
