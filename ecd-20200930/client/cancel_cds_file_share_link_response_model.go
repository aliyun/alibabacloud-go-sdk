// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCdsFileShareLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCdsFileShareLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCdsFileShareLinkResponse
	GetStatusCode() *int32
	SetBody(v *CancelCdsFileShareLinkResponseBody) *CancelCdsFileShareLinkResponse
	GetBody() *CancelCdsFileShareLinkResponseBody
}

type CancelCdsFileShareLinkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCdsFileShareLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCdsFileShareLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCdsFileShareLinkResponse) GoString() string {
	return s.String()
}

func (s *CancelCdsFileShareLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCdsFileShareLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCdsFileShareLinkResponse) GetBody() *CancelCdsFileShareLinkResponseBody {
	return s.Body
}

func (s *CancelCdsFileShareLinkResponse) SetHeaders(v map[string]*string) *CancelCdsFileShareLinkResponse {
	s.Headers = v
	return s
}

func (s *CancelCdsFileShareLinkResponse) SetStatusCode(v int32) *CancelCdsFileShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCdsFileShareLinkResponse) SetBody(v *CancelCdsFileShareLinkResponseBody) *CancelCdsFileShareLinkResponse {
	s.Body = v
	return s
}

func (s *CancelCdsFileShareLinkResponse) Validate() error {
	return dara.Validate(s)
}
