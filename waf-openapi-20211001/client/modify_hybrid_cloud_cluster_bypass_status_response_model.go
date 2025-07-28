// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudClusterBypassStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridCloudClusterBypassStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridCloudClusterBypassStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridCloudClusterBypassStatusResponseBody) *ModifyHybridCloudClusterBypassStatusResponse
	GetBody() *ModifyHybridCloudClusterBypassStatusResponseBody
}

type ModifyHybridCloudClusterBypassStatusResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridCloudClusterBypassStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridCloudClusterBypassStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudClusterBypassStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) GetBody() *ModifyHybridCloudClusterBypassStatusResponseBody {
	return s.Body
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) SetHeaders(v map[string]*string) *ModifyHybridCloudClusterBypassStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) SetStatusCode(v int32) *ModifyHybridCloudClusterBypassStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) SetBody(v *ModifyHybridCloudClusterBypassStatusResponseBody) *ModifyHybridCloudClusterBypassStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) Validate() error {
	return dara.Validate(s)
}
