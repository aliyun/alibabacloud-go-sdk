// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushItemDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushItemDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushItemDataResponse
	GetStatusCode() *int32
	SetBody(v *PushItemDataResponseBody) *PushItemDataResponse
	GetBody() *PushItemDataResponseBody
}

type PushItemDataResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushItemDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushItemDataResponse) String() string {
	return dara.Prettify(s)
}

func (s PushItemDataResponse) GoString() string {
	return s.String()
}

func (s *PushItemDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushItemDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushItemDataResponse) GetBody() *PushItemDataResponseBody {
	return s.Body
}

func (s *PushItemDataResponse) SetHeaders(v map[string]*string) *PushItemDataResponse {
	s.Headers = v
	return s
}

func (s *PushItemDataResponse) SetStatusCode(v int32) *PushItemDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushItemDataResponse) SetBody(v *PushItemDataResponseBody) *PushItemDataResponse {
	s.Body = v
	return s
}

func (s *PushItemDataResponse) Validate() error {
	return dara.Validate(s)
}
