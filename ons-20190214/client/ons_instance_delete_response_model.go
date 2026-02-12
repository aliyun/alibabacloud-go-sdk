// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsInstanceDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsInstanceDeleteResponse
	GetStatusCode() *int32
	SetBody(v *OnsInstanceDeleteResponseBody) *OnsInstanceDeleteResponse
	GetBody() *OnsInstanceDeleteResponseBody
}

type OnsInstanceDeleteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceDeleteResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsInstanceDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsInstanceDeleteResponse) GetBody() *OnsInstanceDeleteResponseBody {
	return s.Body
}

func (s *OnsInstanceDeleteResponse) SetHeaders(v map[string]*string) *OnsInstanceDeleteResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceDeleteResponse) SetStatusCode(v int32) *OnsInstanceDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceDeleteResponse) SetBody(v *OnsInstanceDeleteResponseBody) *OnsInstanceDeleteResponse {
	s.Body = v
	return s
}

func (s *OnsInstanceDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
