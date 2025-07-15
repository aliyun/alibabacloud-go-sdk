// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDeployApplicationGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContinueDeployApplicationGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContinueDeployApplicationGroupResponse
	GetStatusCode() *int32
	SetBody(v *ContinueDeployApplicationGroupResponseBody) *ContinueDeployApplicationGroupResponse
	GetBody() *ContinueDeployApplicationGroupResponseBody
}

type ContinueDeployApplicationGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueDeployApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueDeployApplicationGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ContinueDeployApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *ContinueDeployApplicationGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContinueDeployApplicationGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContinueDeployApplicationGroupResponse) GetBody() *ContinueDeployApplicationGroupResponseBody {
	return s.Body
}

func (s *ContinueDeployApplicationGroupResponse) SetHeaders(v map[string]*string) *ContinueDeployApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *ContinueDeployApplicationGroupResponse) SetStatusCode(v int32) *ContinueDeployApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueDeployApplicationGroupResponse) SetBody(v *ContinueDeployApplicationGroupResponseBody) *ContinueDeployApplicationGroupResponse {
	s.Body = v
	return s
}

func (s *ContinueDeployApplicationGroupResponse) Validate() error {
	return dara.Validate(s)
}
