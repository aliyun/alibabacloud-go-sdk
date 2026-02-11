// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataSourceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDataSourceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDataSourceLogResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDataSourceLogResponseBody) *ModifyDataSourceLogResponse
	GetBody() *ModifyDataSourceLogResponseBody
}

type ModifyDataSourceLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataSourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataSourceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceLogResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDataSourceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDataSourceLogResponse) GetBody() *ModifyDataSourceLogResponseBody {
	return s.Body
}

func (s *ModifyDataSourceLogResponse) SetHeaders(v map[string]*string) *ModifyDataSourceLogResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceLogResponse) SetStatusCode(v int32) *ModifyDataSourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataSourceLogResponse) SetBody(v *ModifyDataSourceLogResponseBody) *ModifyDataSourceLogResponse {
	s.Body = v
	return s
}

func (s *ModifyDataSourceLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
