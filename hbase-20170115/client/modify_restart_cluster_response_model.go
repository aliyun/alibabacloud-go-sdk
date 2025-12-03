// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRestartClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRestartClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRestartClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRestartClusterResponseBody) *ModifyRestartClusterResponse
	GetBody() *ModifyRestartClusterResponseBody
}

type ModifyRestartClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRestartClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRestartClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRestartClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyRestartClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRestartClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRestartClusterResponse) GetBody() *ModifyRestartClusterResponseBody {
	return s.Body
}

func (s *ModifyRestartClusterResponse) SetHeaders(v map[string]*string) *ModifyRestartClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyRestartClusterResponse) SetStatusCode(v int32) *ModifyRestartClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRestartClusterResponse) SetBody(v *ModifyRestartClusterResponseBody) *ModifyRestartClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyRestartClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
