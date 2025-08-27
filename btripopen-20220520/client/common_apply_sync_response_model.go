// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonApplySyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommonApplySyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommonApplySyncResponse
	GetStatusCode() *int32
	SetBody(v *CommonApplySyncResponseBody) *CommonApplySyncResponse
	GetBody() *CommonApplySyncResponseBody
}

type CommonApplySyncResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommonApplySyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommonApplySyncResponse) String() string {
	return dara.Prettify(s)
}

func (s CommonApplySyncResponse) GoString() string {
	return s.String()
}

func (s *CommonApplySyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommonApplySyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommonApplySyncResponse) GetBody() *CommonApplySyncResponseBody {
	return s.Body
}

func (s *CommonApplySyncResponse) SetHeaders(v map[string]*string) *CommonApplySyncResponse {
	s.Headers = v
	return s
}

func (s *CommonApplySyncResponse) SetStatusCode(v int32) *CommonApplySyncResponse {
	s.StatusCode = &v
	return s
}

func (s *CommonApplySyncResponse) SetBody(v *CommonApplySyncResponseBody) *CommonApplySyncResponse {
	s.Body = v
	return s
}

func (s *CommonApplySyncResponse) Validate() error {
	return dara.Validate(s)
}
