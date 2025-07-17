// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveResourceResponse
	GetStatusCode() *int32
	SetBody(v *MoveResourceResponseBody) *MoveResourceResponse
	GetBody() *MoveResourceResponseBody
}

type MoveResourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveResourceResponse) GetBody() *MoveResourceResponseBody {
	return s.Body
}

func (s *MoveResourceResponse) SetHeaders(v map[string]*string) *MoveResourceResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceResponse) SetStatusCode(v int32) *MoveResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceResponse) SetBody(v *MoveResourceResponseBody) *MoveResourceResponse {
	s.Body = v
	return s
}

func (s *MoveResourceResponse) Validate() error {
	return dara.Validate(s)
}
