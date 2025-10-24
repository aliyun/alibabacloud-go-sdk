// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTagGroupResponseBody) *DeleteTagGroupResponse
	GetBody() *DeleteTagGroupResponseBody
}

type DeleteTagGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTagGroupResponse) GetBody() *DeleteTagGroupResponseBody {
	return s.Body
}

func (s *DeleteTagGroupResponse) SetHeaders(v map[string]*string) *DeleteTagGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagGroupResponse) SetStatusCode(v int32) *DeleteTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagGroupResponse) SetBody(v *DeleteTagGroupResponseBody) *DeleteTagGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteTagGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
