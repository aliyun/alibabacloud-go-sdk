// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCfwInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCfwInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCfwInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCfwInstanceResponseBody) *ModifyCfwInstanceResponse
	GetBody() *ModifyCfwInstanceResponseBody
}

type ModifyCfwInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCfwInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCfwInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCfwInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyCfwInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCfwInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCfwInstanceResponse) GetBody() *ModifyCfwInstanceResponseBody {
	return s.Body
}

func (s *ModifyCfwInstanceResponse) SetHeaders(v map[string]*string) *ModifyCfwInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyCfwInstanceResponse) SetStatusCode(v int32) *ModifyCfwInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCfwInstanceResponse) SetBody(v *ModifyCfwInstanceResponseBody) *ModifyCfwInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyCfwInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
