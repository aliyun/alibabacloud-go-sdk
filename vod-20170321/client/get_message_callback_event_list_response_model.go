// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCallbackEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageCallbackEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageCallbackEventListResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageCallbackEventListResponseBody) *GetMessageCallbackEventListResponse
	GetBody() *GetMessageCallbackEventListResponseBody
}

type GetMessageCallbackEventListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageCallbackEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageCallbackEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCallbackEventListResponse) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageCallbackEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageCallbackEventListResponse) GetBody() *GetMessageCallbackEventListResponseBody {
	return s.Body
}

func (s *GetMessageCallbackEventListResponse) SetHeaders(v map[string]*string) *GetMessageCallbackEventListResponse {
	s.Headers = v
	return s
}

func (s *GetMessageCallbackEventListResponse) SetStatusCode(v int32) *GetMessageCallbackEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageCallbackEventListResponse) SetBody(v *GetMessageCallbackEventListResponseBody) *GetMessageCallbackEventListResponse {
	s.Body = v
	return s
}

func (s *GetMessageCallbackEventListResponse) Validate() error {
	return dara.Validate(s)
}
