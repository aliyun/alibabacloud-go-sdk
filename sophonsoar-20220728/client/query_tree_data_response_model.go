// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTreeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTreeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTreeDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryTreeDataResponseBody) *QueryTreeDataResponse
	GetBody() *QueryTreeDataResponseBody
}

type QueryTreeDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTreeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTreeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTreeDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTreeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTreeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTreeDataResponse) GetBody() *QueryTreeDataResponseBody {
	return s.Body
}

func (s *QueryTreeDataResponse) SetHeaders(v map[string]*string) *QueryTreeDataResponse {
	s.Headers = v
	return s
}

func (s *QueryTreeDataResponse) SetStatusCode(v int32) *QueryTreeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTreeDataResponse) SetBody(v *QueryTreeDataResponseBody) *QueryTreeDataResponse {
	s.Body = v
	return s
}

func (s *QueryTreeDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
