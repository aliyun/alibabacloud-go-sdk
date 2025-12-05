// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsScenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePtsScenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePtsScenesResponse
	GetStatusCode() *int32
	SetBody(v *DeletePtsScenesResponseBody) *DeletePtsScenesResponse
	GetBody() *DeletePtsScenesResponseBody
}

type DeletePtsScenesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePtsScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePtsScenesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsScenesResponse) GoString() string {
	return s.String()
}

func (s *DeletePtsScenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePtsScenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePtsScenesResponse) GetBody() *DeletePtsScenesResponseBody {
	return s.Body
}

func (s *DeletePtsScenesResponse) SetHeaders(v map[string]*string) *DeletePtsScenesResponse {
	s.Headers = v
	return s
}

func (s *DeletePtsScenesResponse) SetStatusCode(v int32) *DeletePtsScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePtsScenesResponse) SetBody(v *DeletePtsScenesResponseBody) *DeletePtsScenesResponse {
	s.Body = v
	return s
}

func (s *DeletePtsScenesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
