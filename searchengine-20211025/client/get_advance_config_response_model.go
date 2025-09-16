// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdvanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdvanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAdvanceConfigResponseBody) *GetAdvanceConfigResponse
	GetBody() *GetAdvanceConfigResponseBody
}

type GetAdvanceConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdvanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdvanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdvanceConfigResponse) GetBody() *GetAdvanceConfigResponseBody {
	return s.Body
}

func (s *GetAdvanceConfigResponse) SetHeaders(v map[string]*string) *GetAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAdvanceConfigResponse) SetStatusCode(v int32) *GetAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdvanceConfigResponse) SetBody(v *GetAdvanceConfigResponseBody) *GetAdvanceConfigResponse {
	s.Body = v
	return s
}

func (s *GetAdvanceConfigResponse) Validate() error {
	return dara.Validate(s)
}
