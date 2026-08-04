// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContacterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyContacterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyContacterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyContacterResponseBody) *ModifyContacterResponse
	GetBody() *ModifyContacterResponseBody
}

type ModifyContacterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyContacterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyContacterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyContacterResponse) GoString() string {
	return s.String()
}

func (s *ModifyContacterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyContacterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyContacterResponse) GetBody() *ModifyContacterResponseBody {
	return s.Body
}

func (s *ModifyContacterResponse) SetHeaders(v map[string]*string) *ModifyContacterResponse {
	s.Headers = v
	return s
}

func (s *ModifyContacterResponse) SetStatusCode(v int32) *ModifyContacterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyContacterResponse) SetBody(v *ModifyContacterResponseBody) *ModifyContacterResponse {
	s.Body = v
	return s
}

func (s *ModifyContacterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
