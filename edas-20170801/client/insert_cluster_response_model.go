// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertClusterResponse
	GetStatusCode() *int32
	SetBody(v *InsertClusterResponseBody) *InsertClusterResponse
	GetBody() *InsertClusterResponseBody
}

type InsertClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertClusterResponse) GoString() string {
	return s.String()
}

func (s *InsertClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertClusterResponse) GetBody() *InsertClusterResponseBody {
	return s.Body
}

func (s *InsertClusterResponse) SetHeaders(v map[string]*string) *InsertClusterResponse {
	s.Headers = v
	return s
}

func (s *InsertClusterResponse) SetStatusCode(v int32) *InsertClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertClusterResponse) SetBody(v *InsertClusterResponseBody) *InsertClusterResponse {
	s.Body = v
	return s
}

func (s *InsertClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
