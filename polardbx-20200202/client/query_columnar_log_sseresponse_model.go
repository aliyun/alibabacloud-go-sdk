// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryColumnarLogSSEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryColumnarLogSSEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryColumnarLogSSEResponse
	GetStatusCode() *int32
	SetId(v string) *QueryColumnarLogSSEResponse
	GetId() *string
	SetEvent(v string) *QueryColumnarLogSSEResponse
	GetEvent() *string
	SetBody(v *QueryColumnarLogSSEResponseBody) *QueryColumnarLogSSEResponse
	GetBody() *QueryColumnarLogSSEResponseBody
}

type QueryColumnarLogSSEResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                          `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                          `json:"event,omitempty" xml:"event,omitempty"`
	Body       *QueryColumnarLogSSEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryColumnarLogSSEResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryColumnarLogSSEResponse) GoString() string {
	return s.String()
}

func (s *QueryColumnarLogSSEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryColumnarLogSSEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryColumnarLogSSEResponse) GetId() *string {
	return s.Id
}

func (s *QueryColumnarLogSSEResponse) GetEvent() *string {
	return s.Event
}

func (s *QueryColumnarLogSSEResponse) GetBody() *QueryColumnarLogSSEResponseBody {
	return s.Body
}

func (s *QueryColumnarLogSSEResponse) SetHeaders(v map[string]*string) *QueryColumnarLogSSEResponse {
	s.Headers = v
	return s
}

func (s *QueryColumnarLogSSEResponse) SetStatusCode(v int32) *QueryColumnarLogSSEResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryColumnarLogSSEResponse) SetId(v string) *QueryColumnarLogSSEResponse {
	s.Id = &v
	return s
}

func (s *QueryColumnarLogSSEResponse) SetEvent(v string) *QueryColumnarLogSSEResponse {
	s.Event = &v
	return s
}

func (s *QueryColumnarLogSSEResponse) SetBody(v *QueryColumnarLogSSEResponseBody) *QueryColumnarLogSSEResponse {
	s.Body = v
	return s
}

func (s *QueryColumnarLogSSEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
