// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsInstanceBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsInstanceBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *OnsInstanceBaseInfoResponseBody) *OnsInstanceBaseInfoResponse
	GetBody() *OnsInstanceBaseInfoResponseBody
}

type OnsInstanceBaseInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsInstanceBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsInstanceBaseInfoResponse) GetBody() *OnsInstanceBaseInfoResponseBody {
	return s.Body
}

func (s *OnsInstanceBaseInfoResponse) SetHeaders(v map[string]*string) *OnsInstanceBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceBaseInfoResponse) SetStatusCode(v int32) *OnsInstanceBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceBaseInfoResponse) SetBody(v *OnsInstanceBaseInfoResponseBody) *OnsInstanceBaseInfoResponse {
	s.Body = v
	return s
}

func (s *OnsInstanceBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
