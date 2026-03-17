// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagLanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagLanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagLanResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagLanResponseBody) *ModifySagLanResponse
	GetBody() *ModifySagLanResponseBody
}

type ModifySagLanResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagLanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagLanResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagLanResponse) GoString() string {
	return s.String()
}

func (s *ModifySagLanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagLanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagLanResponse) GetBody() *ModifySagLanResponseBody {
	return s.Body
}

func (s *ModifySagLanResponse) SetHeaders(v map[string]*string) *ModifySagLanResponse {
	s.Headers = v
	return s
}

func (s *ModifySagLanResponse) SetStatusCode(v int32) *ModifySagLanResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagLanResponse) SetBody(v *ModifySagLanResponseBody) *ModifySagLanResponse {
	s.Body = v
	return s
}

func (s *ModifySagLanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
