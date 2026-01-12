// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileRecycledMetaVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MobileRecycledMetaVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MobileRecycledMetaVerifyResponse
	GetStatusCode() *int32
	SetBody(v *MobileRecycledMetaVerifyResponseBody) *MobileRecycledMetaVerifyResponse
	GetBody() *MobileRecycledMetaVerifyResponseBody
}

type MobileRecycledMetaVerifyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MobileRecycledMetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MobileRecycledMetaVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s MobileRecycledMetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *MobileRecycledMetaVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MobileRecycledMetaVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MobileRecycledMetaVerifyResponse) GetBody() *MobileRecycledMetaVerifyResponseBody {
	return s.Body
}

func (s *MobileRecycledMetaVerifyResponse) SetHeaders(v map[string]*string) *MobileRecycledMetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *MobileRecycledMetaVerifyResponse) SetStatusCode(v int32) *MobileRecycledMetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *MobileRecycledMetaVerifyResponse) SetBody(v *MobileRecycledMetaVerifyResponseBody) *MobileRecycledMetaVerifyResponse {
	s.Body = v
	return s
}

func (s *MobileRecycledMetaVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
