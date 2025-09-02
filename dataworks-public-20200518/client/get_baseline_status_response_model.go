// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBaselineStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBaselineStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetBaselineStatusResponseBody) *GetBaselineStatusResponse
	GetBody() *GetBaselineStatusResponseBody
}

type GetBaselineStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBaselineStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBaselineStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineStatusResponse) GoString() string {
	return s.String()
}

func (s *GetBaselineStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBaselineStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBaselineStatusResponse) GetBody() *GetBaselineStatusResponseBody {
	return s.Body
}

func (s *GetBaselineStatusResponse) SetHeaders(v map[string]*string) *GetBaselineStatusResponse {
	s.Headers = v
	return s
}

func (s *GetBaselineStatusResponse) SetStatusCode(v int32) *GetBaselineStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBaselineStatusResponse) SetBody(v *GetBaselineStatusResponseBody) *GetBaselineStatusResponse {
	s.Body = v
	return s
}

func (s *GetBaselineStatusResponse) Validate() error {
	return dara.Validate(s)
}
