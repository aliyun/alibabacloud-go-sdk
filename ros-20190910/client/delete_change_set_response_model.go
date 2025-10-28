// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChangeSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChangeSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChangeSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChangeSetResponseBody) *DeleteChangeSetResponse
	GetBody() *DeleteChangeSetResponseBody
}

type DeleteChangeSetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChangeSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChangeSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteChangeSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChangeSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChangeSetResponse) GetBody() *DeleteChangeSetResponseBody {
	return s.Body
}

func (s *DeleteChangeSetResponse) SetHeaders(v map[string]*string) *DeleteChangeSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteChangeSetResponse) SetStatusCode(v int32) *DeleteChangeSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChangeSetResponse) SetBody(v *DeleteChangeSetResponseBody) *DeleteChangeSetResponse {
	s.Body = v
	return s
}

func (s *DeleteChangeSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
