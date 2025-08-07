// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAndNodesParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterAndNodesParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterAndNodesParametersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterAndNodesParametersResponseBody) *ModifyDBClusterAndNodesParametersResponse
	GetBody() *ModifyDBClusterAndNodesParametersResponseBody
}

type ModifyDBClusterAndNodesParametersResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterAndNodesParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterAndNodesParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAndNodesParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAndNodesParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterAndNodesParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterAndNodesParametersResponse) GetBody() *ModifyDBClusterAndNodesParametersResponseBody {
	return s.Body
}

func (s *ModifyDBClusterAndNodesParametersResponse) SetHeaders(v map[string]*string) *ModifyDBClusterAndNodesParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterAndNodesParametersResponse) SetStatusCode(v int32) *ModifyDBClusterAndNodesParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersResponse) SetBody(v *ModifyDBClusterAndNodesParametersResponseBody) *ModifyDBClusterAndNodesParametersResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterAndNodesParametersResponse) Validate() error {
	return dara.Validate(s)
}
