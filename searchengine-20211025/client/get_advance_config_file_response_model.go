// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvanceConfigFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdvanceConfigFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdvanceConfigFileResponse
	GetStatusCode() *int32
	SetBody(v *GetAdvanceConfigFileResponseBody) *GetAdvanceConfigFileResponse
	GetBody() *GetAdvanceConfigFileResponseBody
}

type GetAdvanceConfigFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdvanceConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdvanceConfigFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdvanceConfigFileResponse) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdvanceConfigFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdvanceConfigFileResponse) GetBody() *GetAdvanceConfigFileResponseBody {
	return s.Body
}

func (s *GetAdvanceConfigFileResponse) SetHeaders(v map[string]*string) *GetAdvanceConfigFileResponse {
	s.Headers = v
	return s
}

func (s *GetAdvanceConfigFileResponse) SetStatusCode(v int32) *GetAdvanceConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdvanceConfigFileResponse) SetBody(v *GetAdvanceConfigFileResponseBody) *GetAdvanceConfigFileResponse {
	s.Body = v
	return s
}

func (s *GetAdvanceConfigFileResponse) Validate() error {
	return dara.Validate(s)
}
