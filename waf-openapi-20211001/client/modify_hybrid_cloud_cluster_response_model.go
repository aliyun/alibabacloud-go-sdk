// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridCloudClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridCloudClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridCloudClusterResponseBody) *ModifyHybridCloudClusterResponse
	GetBody() *ModifyHybridCloudClusterResponseBody
}

type ModifyHybridCloudClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridCloudClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridCloudClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridCloudClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridCloudClusterResponse) GetBody() *ModifyHybridCloudClusterResponseBody {
	return s.Body
}

func (s *ModifyHybridCloudClusterResponse) SetHeaders(v map[string]*string) *ModifyHybridCloudClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridCloudClusterResponse) SetStatusCode(v int32) *ModifyHybridCloudClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridCloudClusterResponse) SetBody(v *ModifyHybridCloudClusterResponseBody) *ModifyHybridCloudClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridCloudClusterResponse) Validate() error {
	return dara.Validate(s)
}
