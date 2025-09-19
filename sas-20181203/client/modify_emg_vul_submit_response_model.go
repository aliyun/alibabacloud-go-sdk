// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmgVulSubmitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEmgVulSubmitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEmgVulSubmitResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEmgVulSubmitResponseBody) *ModifyEmgVulSubmitResponse
	GetBody() *ModifyEmgVulSubmitResponseBody
}

type ModifyEmgVulSubmitResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEmgVulSubmitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEmgVulSubmitResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmgVulSubmitResponse) GoString() string {
	return s.String()
}

func (s *ModifyEmgVulSubmitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEmgVulSubmitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEmgVulSubmitResponse) GetBody() *ModifyEmgVulSubmitResponseBody {
	return s.Body
}

func (s *ModifyEmgVulSubmitResponse) SetHeaders(v map[string]*string) *ModifyEmgVulSubmitResponse {
	s.Headers = v
	return s
}

func (s *ModifyEmgVulSubmitResponse) SetStatusCode(v int32) *ModifyEmgVulSubmitResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEmgVulSubmitResponse) SetBody(v *ModifyEmgVulSubmitResponseBody) *ModifyEmgVulSubmitResponse {
	s.Body = v
	return s
}

func (s *ModifyEmgVulSubmitResponse) Validate() error {
	return dara.Validate(s)
}
