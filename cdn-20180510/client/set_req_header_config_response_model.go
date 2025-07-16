// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetReqHeaderConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetReqHeaderConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetReqHeaderConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetReqHeaderConfigResponseBody) *SetReqHeaderConfigResponse
	GetBody() *SetReqHeaderConfigResponseBody
}

type SetReqHeaderConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetReqHeaderConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetReqHeaderConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetReqHeaderConfigResponse) GoString() string {
	return s.String()
}

func (s *SetReqHeaderConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetReqHeaderConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetReqHeaderConfigResponse) GetBody() *SetReqHeaderConfigResponseBody {
	return s.Body
}

func (s *SetReqHeaderConfigResponse) SetHeaders(v map[string]*string) *SetReqHeaderConfigResponse {
	s.Headers = v
	return s
}

func (s *SetReqHeaderConfigResponse) SetStatusCode(v int32) *SetReqHeaderConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetReqHeaderConfigResponse) SetBody(v *SetReqHeaderConfigResponseBody) *SetReqHeaderConfigResponse {
	s.Body = v
	return s
}

func (s *SetReqHeaderConfigResponse) Validate() error {
	return dara.Validate(s)
}
