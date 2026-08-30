// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSaseUserTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSaseUserTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSaseUserTagResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSaseUserTagResponseBody) *DeleteSaseUserTagResponse
	GetBody() *DeleteSaseUserTagResponseBody
}

type DeleteSaseUserTagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSaseUserTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSaseUserTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSaseUserTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteSaseUserTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSaseUserTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSaseUserTagResponse) GetBody() *DeleteSaseUserTagResponseBody {
	return s.Body
}

func (s *DeleteSaseUserTagResponse) SetHeaders(v map[string]*string) *DeleteSaseUserTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteSaseUserTagResponse) SetStatusCode(v int32) *DeleteSaseUserTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSaseUserTagResponse) SetBody(v *DeleteSaseUserTagResponseBody) *DeleteSaseUserTagResponse {
	s.Body = v
	return s
}

func (s *DeleteSaseUserTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
