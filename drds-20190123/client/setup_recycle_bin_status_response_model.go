// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupRecycleBinStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetupRecycleBinStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetupRecycleBinStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetupRecycleBinStatusResponseBody) *SetupRecycleBinStatusResponse
	GetBody() *SetupRecycleBinStatusResponseBody
}

type SetupRecycleBinStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetupRecycleBinStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetupRecycleBinStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetupRecycleBinStatusResponse) GoString() string {
	return s.String()
}

func (s *SetupRecycleBinStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetupRecycleBinStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetupRecycleBinStatusResponse) GetBody() *SetupRecycleBinStatusResponseBody {
	return s.Body
}

func (s *SetupRecycleBinStatusResponse) SetHeaders(v map[string]*string) *SetupRecycleBinStatusResponse {
	s.Headers = v
	return s
}

func (s *SetupRecycleBinStatusResponse) SetStatusCode(v int32) *SetupRecycleBinStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupRecycleBinStatusResponse) SetBody(v *SetupRecycleBinStatusResponseBody) *SetupRecycleBinStatusResponse {
	s.Body = v
	return s
}

func (s *SetupRecycleBinStatusResponse) Validate() error {
	return dara.Validate(s)
}
