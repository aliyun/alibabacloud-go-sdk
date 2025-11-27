// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageFigureClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchImageFigureClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchImageFigureClusterResponse
	GetStatusCode() *int32
	SetBody(v *SearchImageFigureClusterResponseBody) *SearchImageFigureClusterResponse
	GetBody() *SearchImageFigureClusterResponseBody
}

type SearchImageFigureClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchImageFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchImageFigureClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchImageFigureClusterResponse) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchImageFigureClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchImageFigureClusterResponse) GetBody() *SearchImageFigureClusterResponseBody {
	return s.Body
}

func (s *SearchImageFigureClusterResponse) SetHeaders(v map[string]*string) *SearchImageFigureClusterResponse {
	s.Headers = v
	return s
}

func (s *SearchImageFigureClusterResponse) SetStatusCode(v int32) *SearchImageFigureClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageFigureClusterResponse) SetBody(v *SearchImageFigureClusterResponseBody) *SearchImageFigureClusterResponse {
	s.Body = v
	return s
}

func (s *SearchImageFigureClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
