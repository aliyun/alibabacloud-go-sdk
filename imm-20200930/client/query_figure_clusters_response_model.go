// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFigureClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFigureClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFigureClustersResponse
	GetStatusCode() *int32
	SetBody(v *QueryFigureClustersResponseBody) *QueryFigureClustersResponse
	GetBody() *QueryFigureClustersResponseBody
}

type QueryFigureClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFigureClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFigureClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFigureClustersResponse) GoString() string {
	return s.String()
}

func (s *QueryFigureClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFigureClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFigureClustersResponse) GetBody() *QueryFigureClustersResponseBody {
	return s.Body
}

func (s *QueryFigureClustersResponse) SetHeaders(v map[string]*string) *QueryFigureClustersResponse {
	s.Headers = v
	return s
}

func (s *QueryFigureClustersResponse) SetStatusCode(v int32) *QueryFigureClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFigureClustersResponse) SetBody(v *QueryFigureClustersResponseBody) *QueryFigureClustersResponse {
	s.Body = v
	return s
}

func (s *QueryFigureClustersResponse) Validate() error {
	return dara.Validate(s)
}
