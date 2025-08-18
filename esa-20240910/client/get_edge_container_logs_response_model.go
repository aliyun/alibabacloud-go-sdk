// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerLogsResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerLogsResponseBody) *GetEdgeContainerLogsResponse
	GetBody() *GetEdgeContainerLogsResponseBody
}

type GetEdgeContainerLogsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerLogsResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerLogsResponse) GetBody() *GetEdgeContainerLogsResponseBody {
	return s.Body
}

func (s *GetEdgeContainerLogsResponse) SetHeaders(v map[string]*string) *GetEdgeContainerLogsResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerLogsResponse) SetStatusCode(v int32) *GetEdgeContainerLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerLogsResponse) SetBody(v *GetEdgeContainerLogsResponseBody) *GetEdgeContainerLogsResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerLogsResponse) Validate() error {
	return dara.Validate(s)
}
