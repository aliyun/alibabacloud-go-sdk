// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCVClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCVClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCVClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCVClusterResponseBody) *ModifyRCVClusterResponse
	GetBody() *ModifyRCVClusterResponseBody
}

type ModifyRCVClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCVClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCVClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCVClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCVClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCVClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCVClusterResponse) GetBody() *ModifyRCVClusterResponseBody {
	return s.Body
}

func (s *ModifyRCVClusterResponse) SetHeaders(v map[string]*string) *ModifyRCVClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCVClusterResponse) SetStatusCode(v int32) *ModifyRCVClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCVClusterResponse) SetBody(v *ModifyRCVClusterResponseBody) *ModifyRCVClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyRCVClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
