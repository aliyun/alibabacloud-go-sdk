// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelShareLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelShareLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelShareLinkResponse
	GetStatusCode() *int32
}

type CancelShareLinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CancelShareLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelShareLinkResponse) GoString() string {
	return s.String()
}

func (s *CancelShareLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelShareLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelShareLinkResponse) SetHeaders(v map[string]*string) *CancelShareLinkResponse {
	s.Headers = v
	return s
}

func (s *CancelShareLinkResponse) SetStatusCode(v int32) *CancelShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelShareLinkResponse) Validate() error {
	return dara.Validate(s)
}
