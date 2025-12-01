// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDataLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDataLimitResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDataLimitResponseBody) *ModifyDataLimitResponse
	GetBody() *ModifyDataLimitResponseBody
}

type ModifyDataLimitResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataLimitResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDataLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDataLimitResponse) GetBody() *ModifyDataLimitResponseBody {
	return s.Body
}

func (s *ModifyDataLimitResponse) SetHeaders(v map[string]*string) *ModifyDataLimitResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataLimitResponse) SetStatusCode(v int32) *ModifyDataLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataLimitResponse) SetBody(v *ModifyDataLimitResponseBody) *ModifyDataLimitResponse {
	s.Body = v
	return s
}

func (s *ModifyDataLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
