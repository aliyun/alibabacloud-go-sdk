// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnlineAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnlineAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *OnlineAppInstanceResponseBody) *OnlineAppInstanceResponse
	GetBody() *OnlineAppInstanceResponseBody
}

type OnlineAppInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s OnlineAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *OnlineAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnlineAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnlineAppInstanceResponse) GetBody() *OnlineAppInstanceResponseBody {
	return s.Body
}

func (s *OnlineAppInstanceResponse) SetHeaders(v map[string]*string) *OnlineAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *OnlineAppInstanceResponse) SetStatusCode(v int32) *OnlineAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineAppInstanceResponse) SetBody(v *OnlineAppInstanceResponseBody) *OnlineAppInstanceResponse {
	s.Body = v
	return s
}

func (s *OnlineAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
