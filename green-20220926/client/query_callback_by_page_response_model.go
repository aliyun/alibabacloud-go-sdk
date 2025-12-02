// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallbackByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCallbackByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCallbackByPageResponse
	GetStatusCode() *int32
	SetBody(v *QueryCallbackByPageResponseBody) *QueryCallbackByPageResponse
	GetBody() *QueryCallbackByPageResponseBody
}

type QueryCallbackByPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallbackByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallbackByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCallbackByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryCallbackByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCallbackByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCallbackByPageResponse) GetBody() *QueryCallbackByPageResponseBody {
	return s.Body
}

func (s *QueryCallbackByPageResponse) SetHeaders(v map[string]*string) *QueryCallbackByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryCallbackByPageResponse) SetStatusCode(v int32) *QueryCallbackByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallbackByPageResponse) SetBody(v *QueryCallbackByPageResponseBody) *QueryCallbackByPageResponse {
	s.Body = v
	return s
}

func (s *QueryCallbackByPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
