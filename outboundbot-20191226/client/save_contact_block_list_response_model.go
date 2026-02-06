// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactBlockListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveContactBlockListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveContactBlockListResponse
	GetStatusCode() *int32
	SetBody(v *SaveContactBlockListResponseBody) *SaveContactBlockListResponse
	GetBody() *SaveContactBlockListResponseBody
}

type SaveContactBlockListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveContactBlockListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveContactBlockListResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveContactBlockListResponse) GoString() string {
	return s.String()
}

func (s *SaveContactBlockListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveContactBlockListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveContactBlockListResponse) GetBody() *SaveContactBlockListResponseBody {
	return s.Body
}

func (s *SaveContactBlockListResponse) SetHeaders(v map[string]*string) *SaveContactBlockListResponse {
	s.Headers = v
	return s
}

func (s *SaveContactBlockListResponse) SetStatusCode(v int32) *SaveContactBlockListResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveContactBlockListResponse) SetBody(v *SaveContactBlockListResponseBody) *SaveContactBlockListResponse {
	s.Body = v
	return s
}

func (s *SaveContactBlockListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
