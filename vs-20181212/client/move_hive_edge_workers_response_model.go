// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveHiveEdgeWorkersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveHiveEdgeWorkersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveHiveEdgeWorkersResponse
	GetStatusCode() *int32
	SetBody(v *MoveHiveEdgeWorkersResponseBody) *MoveHiveEdgeWorkersResponse
	GetBody() *MoveHiveEdgeWorkersResponseBody
}

type MoveHiveEdgeWorkersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveHiveEdgeWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveHiveEdgeWorkersResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveHiveEdgeWorkersResponse) GoString() string {
	return s.String()
}

func (s *MoveHiveEdgeWorkersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveHiveEdgeWorkersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveHiveEdgeWorkersResponse) GetBody() *MoveHiveEdgeWorkersResponseBody {
	return s.Body
}

func (s *MoveHiveEdgeWorkersResponse) SetHeaders(v map[string]*string) *MoveHiveEdgeWorkersResponse {
	s.Headers = v
	return s
}

func (s *MoveHiveEdgeWorkersResponse) SetStatusCode(v int32) *MoveHiveEdgeWorkersResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveHiveEdgeWorkersResponse) SetBody(v *MoveHiveEdgeWorkersResponseBody) *MoveHiveEdgeWorkersResponse {
	s.Body = v
	return s
}

func (s *MoveHiveEdgeWorkersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
