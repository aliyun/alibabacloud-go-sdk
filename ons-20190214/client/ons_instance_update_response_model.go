// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsInstanceUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsInstanceUpdateResponse
	GetStatusCode() *int32
	SetBody(v *OnsInstanceUpdateResponseBody) *OnsInstanceUpdateResponse
	GetBody() *OnsInstanceUpdateResponseBody
}

type OnsInstanceUpdateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceUpdateResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsInstanceUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsInstanceUpdateResponse) GetBody() *OnsInstanceUpdateResponseBody {
	return s.Body
}

func (s *OnsInstanceUpdateResponse) SetHeaders(v map[string]*string) *OnsInstanceUpdateResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceUpdateResponse) SetStatusCode(v int32) *OnsInstanceUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceUpdateResponse) SetBody(v *OnsInstanceUpdateResponseBody) *OnsInstanceUpdateResponse {
	s.Body = v
	return s
}

func (s *OnsInstanceUpdateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
