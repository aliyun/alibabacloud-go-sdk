// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunManualDagNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunManualDagNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunManualDagNodesResponse
	GetStatusCode() *int32
	SetBody(v *RunManualDagNodesResponseBody) *RunManualDagNodesResponse
	GetBody() *RunManualDagNodesResponseBody
}

type RunManualDagNodesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunManualDagNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunManualDagNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RunManualDagNodesResponse) GoString() string {
	return s.String()
}

func (s *RunManualDagNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunManualDagNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunManualDagNodesResponse) GetBody() *RunManualDagNodesResponseBody {
	return s.Body
}

func (s *RunManualDagNodesResponse) SetHeaders(v map[string]*string) *RunManualDagNodesResponse {
	s.Headers = v
	return s
}

func (s *RunManualDagNodesResponse) SetStatusCode(v int32) *RunManualDagNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RunManualDagNodesResponse) SetBody(v *RunManualDagNodesResponseBody) *RunManualDagNodesResponse {
	s.Body = v
	return s
}

func (s *RunManualDagNodesResponse) Validate() error {
	return dara.Validate(s)
}
