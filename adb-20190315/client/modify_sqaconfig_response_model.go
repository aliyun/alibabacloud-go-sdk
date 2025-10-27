// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQAConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySQAConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySQAConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifySQAConfigResponseBody) *ModifySQAConfigResponse
	GetBody() *ModifySQAConfigResponseBody
}

type ModifySQAConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySQAConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySQAConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySQAConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifySQAConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySQAConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySQAConfigResponse) GetBody() *ModifySQAConfigResponseBody {
	return s.Body
}

func (s *ModifySQAConfigResponse) SetHeaders(v map[string]*string) *ModifySQAConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifySQAConfigResponse) SetStatusCode(v int32) *ModifySQAConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySQAConfigResponse) SetBody(v *ModifySQAConfigResponseBody) *ModifySQAConfigResponse {
	s.Body = v
	return s
}

func (s *ModifySQAConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
