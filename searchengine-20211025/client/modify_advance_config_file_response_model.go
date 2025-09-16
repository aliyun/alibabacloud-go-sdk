// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvanceConfigFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAdvanceConfigFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAdvanceConfigFileResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAdvanceConfigFileResponseBody) *ModifyAdvanceConfigFileResponse
	GetBody() *ModifyAdvanceConfigFileResponseBody
}

type ModifyAdvanceConfigFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAdvanceConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAdvanceConfigFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvanceConfigFileResponse) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAdvanceConfigFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAdvanceConfigFileResponse) GetBody() *ModifyAdvanceConfigFileResponseBody {
	return s.Body
}

func (s *ModifyAdvanceConfigFileResponse) SetHeaders(v map[string]*string) *ModifyAdvanceConfigFileResponse {
	s.Headers = v
	return s
}

func (s *ModifyAdvanceConfigFileResponse) SetStatusCode(v int32) *ModifyAdvanceConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAdvanceConfigFileResponse) SetBody(v *ModifyAdvanceConfigFileResponseBody) *ModifyAdvanceConfigFileResponse {
	s.Body = v
	return s
}

func (s *ModifyAdvanceConfigFileResponse) Validate() error {
	return dara.Validate(s)
}
