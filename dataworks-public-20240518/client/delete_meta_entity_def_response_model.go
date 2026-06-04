// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaEntityDefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMetaEntityDefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMetaEntityDefResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMetaEntityDefResponseBody) *DeleteMetaEntityDefResponse
	GetBody() *DeleteMetaEntityDefResponseBody
}

type DeleteMetaEntityDefResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetaEntityDefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetaEntityDefResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaEntityDefResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetaEntityDefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMetaEntityDefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMetaEntityDefResponse) GetBody() *DeleteMetaEntityDefResponseBody {
	return s.Body
}

func (s *DeleteMetaEntityDefResponse) SetHeaders(v map[string]*string) *DeleteMetaEntityDefResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetaEntityDefResponse) SetStatusCode(v int32) *DeleteMetaEntityDefResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetaEntityDefResponse) SetBody(v *DeleteMetaEntityDefResponseBody) *DeleteMetaEntityDefResponse {
	s.Body = v
	return s
}

func (s *DeleteMetaEntityDefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
