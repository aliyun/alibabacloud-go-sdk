// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddInstanceWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddInstanceWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *AddInstanceWhiteListResponseBody) *AddInstanceWhiteListResponse
	GetBody() *AddInstanceWhiteListResponseBody
}

type AddInstanceWhiteListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddInstanceWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddInstanceWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddInstanceWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddInstanceWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddInstanceWhiteListResponse) GetBody() *AddInstanceWhiteListResponseBody {
	return s.Body
}

func (s *AddInstanceWhiteListResponse) SetHeaders(v map[string]*string) *AddInstanceWhiteListResponse {
	s.Headers = v
	return s
}

func (s *AddInstanceWhiteListResponse) SetStatusCode(v int32) *AddInstanceWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddInstanceWhiteListResponse) SetBody(v *AddInstanceWhiteListResponseBody) *AddInstanceWhiteListResponse {
	s.Body = v
	return s
}

func (s *AddInstanceWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
