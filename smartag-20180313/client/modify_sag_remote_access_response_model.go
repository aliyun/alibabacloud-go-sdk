// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagRemoteAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagRemoteAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagRemoteAccessResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagRemoteAccessResponseBody) *ModifySagRemoteAccessResponse
	GetBody() *ModifySagRemoteAccessResponseBody
}

type ModifySagRemoteAccessResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagRemoteAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagRemoteAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagRemoteAccessResponse) GoString() string {
	return s.String()
}

func (s *ModifySagRemoteAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagRemoteAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagRemoteAccessResponse) GetBody() *ModifySagRemoteAccessResponseBody {
	return s.Body
}

func (s *ModifySagRemoteAccessResponse) SetHeaders(v map[string]*string) *ModifySagRemoteAccessResponse {
	s.Headers = v
	return s
}

func (s *ModifySagRemoteAccessResponse) SetStatusCode(v int32) *ModifySagRemoteAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagRemoteAccessResponse) SetBody(v *ModifySagRemoteAccessResponseBody) *ModifySagRemoteAccessResponse {
	s.Body = v
	return s
}

func (s *ModifySagRemoteAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
