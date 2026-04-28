// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyModelApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyModelApiResponse
	GetStatusCode() *int32
	SetBody(v *ModifyModelApiResponseBody) *ModifyModelApiResponse
	GetBody() *ModifyModelApiResponseBody
}

type ModifyModelApiResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyModelApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyModelApiResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelApiResponse) GoString() string {
	return s.String()
}

func (s *ModifyModelApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyModelApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyModelApiResponse) GetBody() *ModifyModelApiResponseBody {
	return s.Body
}

func (s *ModifyModelApiResponse) SetHeaders(v map[string]*string) *ModifyModelApiResponse {
	s.Headers = v
	return s
}

func (s *ModifyModelApiResponse) SetStatusCode(v int32) *ModifyModelApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyModelApiResponse) SetBody(v *ModifyModelApiResponseBody) *ModifyModelApiResponse {
	s.Body = v
	return s
}

func (s *ModifyModelApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
