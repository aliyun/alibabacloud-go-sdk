// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFEClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFEClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFEClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFEClusterResponseBody) *ModifyFEClusterResponse
	GetBody() *ModifyFEClusterResponseBody
}

type ModifyFEClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFEClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFEClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFEClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyFEClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFEClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFEClusterResponse) GetBody() *ModifyFEClusterResponseBody {
	return s.Body
}

func (s *ModifyFEClusterResponse) SetHeaders(v map[string]*string) *ModifyFEClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyFEClusterResponse) SetStatusCode(v int32) *ModifyFEClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFEClusterResponse) SetBody(v *ModifyFEClusterResponseBody) *ModifyFEClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyFEClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
