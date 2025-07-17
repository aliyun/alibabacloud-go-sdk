// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagToFlinkClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTagToFlinkClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTagToFlinkClusterResponse
	GetStatusCode() *int32
	SetBody(v *AddTagToFlinkClusterResponseBody) *AddTagToFlinkClusterResponse
	GetBody() *AddTagToFlinkClusterResponseBody
}

type AddTagToFlinkClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTagToFlinkClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTagToFlinkClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTagToFlinkClusterResponse) GoString() string {
	return s.String()
}

func (s *AddTagToFlinkClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTagToFlinkClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTagToFlinkClusterResponse) GetBody() *AddTagToFlinkClusterResponseBody {
	return s.Body
}

func (s *AddTagToFlinkClusterResponse) SetHeaders(v map[string]*string) *AddTagToFlinkClusterResponse {
	s.Headers = v
	return s
}

func (s *AddTagToFlinkClusterResponse) SetStatusCode(v int32) *AddTagToFlinkClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTagToFlinkClusterResponse) SetBody(v *AddTagToFlinkClusterResponseBody) *AddTagToFlinkClusterResponse {
	s.Body = v
	return s
}

func (s *AddTagToFlinkClusterResponse) Validate() error {
	return dara.Validate(s)
}
