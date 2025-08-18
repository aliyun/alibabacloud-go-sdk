// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppLogRiverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerAppLogRiverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerAppLogRiverResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerAppLogRiverResponseBody) *GetEdgeContainerAppLogRiverResponse
	GetBody() *GetEdgeContainerAppLogRiverResponseBody
}

type GetEdgeContainerAppLogRiverResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerAppLogRiverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerAppLogRiverResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppLogRiverResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppLogRiverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerAppLogRiverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerAppLogRiverResponse) GetBody() *GetEdgeContainerAppLogRiverResponseBody {
	return s.Body
}

func (s *GetEdgeContainerAppLogRiverResponse) SetHeaders(v map[string]*string) *GetEdgeContainerAppLogRiverResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerAppLogRiverResponse) SetStatusCode(v int32) *GetEdgeContainerAppLogRiverResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerAppLogRiverResponse) SetBody(v *GetEdgeContainerAppLogRiverResponseBody) *GetEdgeContainerAppLogRiverResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerAppLogRiverResponse) Validate() error {
	return dara.Validate(s)
}
