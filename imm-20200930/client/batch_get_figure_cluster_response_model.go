// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFigureClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetFigureClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetFigureClusterResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetFigureClusterResponseBody) *BatchGetFigureClusterResponse
	GetBody() *BatchGetFigureClusterResponseBody
}

type BatchGetFigureClusterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetFigureClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFigureClusterResponse) GoString() string {
	return s.String()
}

func (s *BatchGetFigureClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetFigureClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetFigureClusterResponse) GetBody() *BatchGetFigureClusterResponseBody {
	return s.Body
}

func (s *BatchGetFigureClusterResponse) SetHeaders(v map[string]*string) *BatchGetFigureClusterResponse {
	s.Headers = v
	return s
}

func (s *BatchGetFigureClusterResponse) SetStatusCode(v int32) *BatchGetFigureClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetFigureClusterResponse) SetBody(v *BatchGetFigureClusterResponseBody) *BatchGetFigureClusterResponse {
	s.Body = v
	return s
}

func (s *BatchGetFigureClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
