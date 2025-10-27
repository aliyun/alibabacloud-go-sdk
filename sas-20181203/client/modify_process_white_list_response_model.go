// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProcessWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyProcessWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyProcessWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyProcessWhiteListResponseBody) *ModifyProcessWhiteListResponse
	GetBody() *ModifyProcessWhiteListResponseBody
}

type ModifyProcessWhiteListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProcessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyProcessWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyProcessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyProcessWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyProcessWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyProcessWhiteListResponse) GetBody() *ModifyProcessWhiteListResponseBody {
	return s.Body
}

func (s *ModifyProcessWhiteListResponse) SetHeaders(v map[string]*string) *ModifyProcessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyProcessWhiteListResponse) SetStatusCode(v int32) *ModifyProcessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProcessWhiteListResponse) SetBody(v *ModifyProcessWhiteListResponseBody) *ModifyProcessWhiteListResponse {
	s.Body = v
	return s
}

func (s *ModifyProcessWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
