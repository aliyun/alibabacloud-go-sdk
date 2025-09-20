// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFigureClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFigureClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFigureClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetFigureClusterResponseBody) *GetFigureClusterResponse
	GetBody() *GetFigureClusterResponseBody
}

type GetFigureClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFigureClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFigureClusterResponse) GoString() string {
	return s.String()
}

func (s *GetFigureClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFigureClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFigureClusterResponse) GetBody() *GetFigureClusterResponseBody {
	return s.Body
}

func (s *GetFigureClusterResponse) SetHeaders(v map[string]*string) *GetFigureClusterResponse {
	s.Headers = v
	return s
}

func (s *GetFigureClusterResponse) SetStatusCode(v int32) *GetFigureClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFigureClusterResponse) SetBody(v *GetFigureClusterResponseBody) *GetFigureClusterResponse {
	s.Body = v
	return s
}

func (s *GetFigureClusterResponse) Validate() error {
	return dara.Validate(s)
}
