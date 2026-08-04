// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountRealNameInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountRealNameInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountRealNameInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountRealNameInfoResponseBody) *QueryAccountRealNameInfoResponse
	GetBody() *QueryAccountRealNameInfoResponseBody
}

type QueryAccountRealNameInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountRealNameInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountRealNameInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountRealNameInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountRealNameInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountRealNameInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountRealNameInfoResponse) GetBody() *QueryAccountRealNameInfoResponseBody {
	return s.Body
}

func (s *QueryAccountRealNameInfoResponse) SetHeaders(v map[string]*string) *QueryAccountRealNameInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountRealNameInfoResponse) SetStatusCode(v int32) *QueryAccountRealNameInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountRealNameInfoResponse) SetBody(v *QueryAccountRealNameInfoResponseBody) *QueryAccountRealNameInfoResponse {
	s.Body = v
	return s
}

func (s *QueryAccountRealNameInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
