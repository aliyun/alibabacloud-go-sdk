// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcdpCrowdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcdpCrowdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcdpCrowdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcdpCrowdResponseBody) *DeleteMcdpCrowdResponse
	GetBody() *DeleteMcdpCrowdResponseBody
}

type DeleteMcdpCrowdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcdpCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcdpCrowdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpCrowdResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcdpCrowdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcdpCrowdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcdpCrowdResponse) GetBody() *DeleteMcdpCrowdResponseBody {
	return s.Body
}

func (s *DeleteMcdpCrowdResponse) SetHeaders(v map[string]*string) *DeleteMcdpCrowdResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcdpCrowdResponse) SetStatusCode(v int32) *DeleteMcdpCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcdpCrowdResponse) SetBody(v *DeleteMcdpCrowdResponseBody) *DeleteMcdpCrowdResponse {
	s.Body = v
	return s
}

func (s *DeleteMcdpCrowdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
