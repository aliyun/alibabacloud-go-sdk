// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotwordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunHotwordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunHotwordResponse
	GetStatusCode() *int32
	SetId(v string) *RunHotwordResponse
	GetId() *string
	SetEvent(v string) *RunHotwordResponse
	GetEvent() *string
	SetBody(v *RunHotwordResponseBody) *RunHotwordResponse
	GetBody() *RunHotwordResponseBody
}

type RunHotwordResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                 `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                 `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunHotwordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunHotwordResponse) String() string {
	return dara.Prettify(s)
}

func (s RunHotwordResponse) GoString() string {
	return s.String()
}

func (s *RunHotwordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunHotwordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunHotwordResponse) GetId() *string {
	return s.Id
}

func (s *RunHotwordResponse) GetEvent() *string {
	return s.Event
}

func (s *RunHotwordResponse) GetBody() *RunHotwordResponseBody {
	return s.Body
}

func (s *RunHotwordResponse) SetHeaders(v map[string]*string) *RunHotwordResponse {
	s.Headers = v
	return s
}

func (s *RunHotwordResponse) SetStatusCode(v int32) *RunHotwordResponse {
	s.StatusCode = &v
	return s
}

func (s *RunHotwordResponse) SetId(v string) *RunHotwordResponse {
	s.Id = &v
	return s
}

func (s *RunHotwordResponse) SetEvent(v string) *RunHotwordResponse {
	s.Event = &v
	return s
}

func (s *RunHotwordResponse) SetBody(v *RunHotwordResponseBody) *RunHotwordResponse {
	s.Body = v
	return s
}

func (s *RunHotwordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
