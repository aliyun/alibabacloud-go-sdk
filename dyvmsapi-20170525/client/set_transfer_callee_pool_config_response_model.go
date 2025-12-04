// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTransferCalleePoolConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetTransferCalleePoolConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetTransferCalleePoolConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetTransferCalleePoolConfigResponseBody) *SetTransferCalleePoolConfigResponse
	GetBody() *SetTransferCalleePoolConfigResponseBody
}

type SetTransferCalleePoolConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTransferCalleePoolConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTransferCalleePoolConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetTransferCalleePoolConfigResponse) GoString() string {
	return s.String()
}

func (s *SetTransferCalleePoolConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetTransferCalleePoolConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetTransferCalleePoolConfigResponse) GetBody() *SetTransferCalleePoolConfigResponseBody {
	return s.Body
}

func (s *SetTransferCalleePoolConfigResponse) SetHeaders(v map[string]*string) *SetTransferCalleePoolConfigResponse {
	s.Headers = v
	return s
}

func (s *SetTransferCalleePoolConfigResponse) SetStatusCode(v int32) *SetTransferCalleePoolConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTransferCalleePoolConfigResponse) SetBody(v *SetTransferCalleePoolConfigResponseBody) *SetTransferCalleePoolConfigResponse {
	s.Body = v
	return s
}

func (s *SetTransferCalleePoolConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
