// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnLinkAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnLinkAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnLinkAccountResponse
	GetStatusCode() *int32
}

type UnLinkAccountResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UnLinkAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UnLinkAccountResponse) GoString() string {
	return s.String()
}

func (s *UnLinkAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnLinkAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnLinkAccountResponse) SetHeaders(v map[string]*string) *UnLinkAccountResponse {
	s.Headers = v
	return s
}

func (s *UnLinkAccountResponse) SetStatusCode(v int32) *UnLinkAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UnLinkAccountResponse) Validate() error {
	return dara.Validate(s)
}
