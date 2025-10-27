// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSearchConditionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSearchConditionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSearchConditionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSearchConditionResponseBody) *DeleteSearchConditionResponse
	GetBody() *DeleteSearchConditionResponseBody
}

type DeleteSearchConditionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSearchConditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSearchConditionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSearchConditionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSearchConditionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSearchConditionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSearchConditionResponse) GetBody() *DeleteSearchConditionResponseBody {
	return s.Body
}

func (s *DeleteSearchConditionResponse) SetHeaders(v map[string]*string) *DeleteSearchConditionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSearchConditionResponse) SetStatusCode(v int32) *DeleteSearchConditionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSearchConditionResponse) SetBody(v *DeleteSearchConditionResponseBody) *DeleteSearchConditionResponse {
	s.Body = v
	return s
}

func (s *DeleteSearchConditionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
