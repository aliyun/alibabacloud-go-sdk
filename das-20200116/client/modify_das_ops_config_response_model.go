// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDasOpsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDasOpsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDasOpsConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDasOpsConfigResponseBody) *ModifyDasOpsConfigResponse
	GetBody() *ModifyDasOpsConfigResponseBody
}

type ModifyDasOpsConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDasOpsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDasOpsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDasOpsConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDasOpsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDasOpsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDasOpsConfigResponse) GetBody() *ModifyDasOpsConfigResponseBody {
	return s.Body
}

func (s *ModifyDasOpsConfigResponse) SetHeaders(v map[string]*string) *ModifyDasOpsConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDasOpsConfigResponse) SetStatusCode(v int32) *ModifyDasOpsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDasOpsConfigResponse) SetBody(v *ModifyDasOpsConfigResponseBody) *ModifyDasOpsConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDasOpsConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
