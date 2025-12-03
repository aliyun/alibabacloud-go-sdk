// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMultiZoneClusterNodeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMultiZoneClusterNodeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMultiZoneClusterNodeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMultiZoneClusterNodeTypeResponseBody) *ModifyMultiZoneClusterNodeTypeResponse
	GetBody() *ModifyMultiZoneClusterNodeTypeResponseBody
}

type ModifyMultiZoneClusterNodeTypeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMultiZoneClusterNodeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMultiZoneClusterNodeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMultiZoneClusterNodeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) GetBody() *ModifyMultiZoneClusterNodeTypeResponseBody {
	return s.Body
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) SetHeaders(v map[string]*string) *ModifyMultiZoneClusterNodeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) SetStatusCode(v int32) *ModifyMultiZoneClusterNodeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) SetBody(v *ModifyMultiZoneClusterNodeTypeResponseBody) *ModifyMultiZoneClusterNodeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
