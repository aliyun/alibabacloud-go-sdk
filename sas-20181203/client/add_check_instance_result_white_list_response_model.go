// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCheckInstanceResultWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCheckInstanceResultWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCheckInstanceResultWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *AddCheckInstanceResultWhiteListResponseBody) *AddCheckInstanceResultWhiteListResponse
	GetBody() *AddCheckInstanceResultWhiteListResponseBody
}

type AddCheckInstanceResultWhiteListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCheckInstanceResultWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCheckInstanceResultWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCheckInstanceResultWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddCheckInstanceResultWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCheckInstanceResultWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCheckInstanceResultWhiteListResponse) GetBody() *AddCheckInstanceResultWhiteListResponseBody {
	return s.Body
}

func (s *AddCheckInstanceResultWhiteListResponse) SetHeaders(v map[string]*string) *AddCheckInstanceResultWhiteListResponse {
	s.Headers = v
	return s
}

func (s *AddCheckInstanceResultWhiteListResponse) SetStatusCode(v int32) *AddCheckInstanceResultWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCheckInstanceResultWhiteListResponse) SetBody(v *AddCheckInstanceResultWhiteListResponseBody) *AddCheckInstanceResultWhiteListResponse {
	s.Body = v
	return s
}

func (s *AddCheckInstanceResultWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
