// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDrdsAtomUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshDrdsAtomUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshDrdsAtomUrlResponse
	GetStatusCode() *int32
	SetBody(v *RefreshDrdsAtomUrlResponseBody) *RefreshDrdsAtomUrlResponse
	GetBody() *RefreshDrdsAtomUrlResponseBody
}

type RefreshDrdsAtomUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshDrdsAtomUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshDrdsAtomUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshDrdsAtomUrlResponse) GoString() string {
	return s.String()
}

func (s *RefreshDrdsAtomUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshDrdsAtomUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshDrdsAtomUrlResponse) GetBody() *RefreshDrdsAtomUrlResponseBody {
	return s.Body
}

func (s *RefreshDrdsAtomUrlResponse) SetHeaders(v map[string]*string) *RefreshDrdsAtomUrlResponse {
	s.Headers = v
	return s
}

func (s *RefreshDrdsAtomUrlResponse) SetStatusCode(v int32) *RefreshDrdsAtomUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponse) SetBody(v *RefreshDrdsAtomUrlResponseBody) *RefreshDrdsAtomUrlResponse {
	s.Body = v
	return s
}

func (s *RefreshDrdsAtomUrlResponse) Validate() error {
	return dara.Validate(s)
}
