// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestModelFileSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RequestModelFileSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RequestModelFileSyncResponse
	GetStatusCode() *int32
	SetBody(v *RequestModelFileSyncResponseBody) *RequestModelFileSyncResponse
	GetBody() *RequestModelFileSyncResponseBody
}

type RequestModelFileSyncResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestModelFileSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestModelFileSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s RequestModelFileSyncResponse) GoString() string {
	return s.String()
}

func (s *RequestModelFileSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RequestModelFileSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RequestModelFileSyncResponse) GetBody() *RequestModelFileSyncResponseBody {
	return s.Body
}

func (s *RequestModelFileSyncResponse) SetHeaders(v map[string]*string) *RequestModelFileSyncResponse {
	s.Headers = v
	return s
}

func (s *RequestModelFileSyncResponse) SetStatusCode(v int32) *RequestModelFileSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestModelFileSyncResponse) SetBody(v *RequestModelFileSyncResponseBody) *RequestModelFileSyncResponse {
	s.Body = v
	return s
}

func (s *RequestModelFileSyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
