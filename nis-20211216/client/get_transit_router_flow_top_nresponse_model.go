// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransitRouterFlowTopNResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTransitRouterFlowTopNResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTransitRouterFlowTopNResponse
	GetStatusCode() *int32
	SetBody(v *GetTransitRouterFlowTopNResponseBody) *GetTransitRouterFlowTopNResponse
	GetBody() *GetTransitRouterFlowTopNResponseBody
}

type GetTransitRouterFlowTopNResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTransitRouterFlowTopNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTransitRouterFlowTopNResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTransitRouterFlowTopNResponse) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTransitRouterFlowTopNResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTransitRouterFlowTopNResponse) GetBody() *GetTransitRouterFlowTopNResponseBody {
	return s.Body
}

func (s *GetTransitRouterFlowTopNResponse) SetHeaders(v map[string]*string) *GetTransitRouterFlowTopNResponse {
	s.Headers = v
	return s
}

func (s *GetTransitRouterFlowTopNResponse) SetStatusCode(v int32) *GetTransitRouterFlowTopNResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponse) SetBody(v *GetTransitRouterFlowTopNResponseBody) *GetTransitRouterFlowTopNResponse {
	s.Body = v
	return s
}

func (s *GetTransitRouterFlowTopNResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
