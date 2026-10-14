// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotfixStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHotfixStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHotfixStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryHotfixStatusResponseBody) *QueryHotfixStatusResponse
	GetBody() *QueryHotfixStatusResponseBody
}

type QueryHotfixStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHotfixStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHotfixStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHotfixStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryHotfixStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHotfixStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHotfixStatusResponse) GetBody() *QueryHotfixStatusResponseBody {
	return s.Body
}

func (s *QueryHotfixStatusResponse) SetHeaders(v map[string]*string) *QueryHotfixStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryHotfixStatusResponse) SetStatusCode(v int32) *QueryHotfixStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotfixStatusResponse) SetBody(v *QueryHotfixStatusResponseBody) *QueryHotfixStatusResponse {
	s.Body = v
	return s
}

func (s *QueryHotfixStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
