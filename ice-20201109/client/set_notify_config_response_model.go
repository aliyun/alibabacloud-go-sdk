// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetNotifyConfigResponseBody) *SetNotifyConfigResponse
	GetBody() *SetNotifyConfigResponseBody
}

type SetNotifyConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *SetNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetNotifyConfigResponse) GetBody() *SetNotifyConfigResponseBody {
	return s.Body
}

func (s *SetNotifyConfigResponse) SetHeaders(v map[string]*string) *SetNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *SetNotifyConfigResponse) SetStatusCode(v int32) *SetNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetNotifyConfigResponse) SetBody(v *SetNotifyConfigResponseBody) *SetNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *SetNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}
