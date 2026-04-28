// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLinkInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLinkInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLinkInfoResponse
	GetStatusCode() *int32
	SetBody(v *AccountLinkInfo) *GetLinkInfoResponse
	GetBody() *AccountLinkInfo
}

type GetLinkInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountLinkInfo   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLinkInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLinkInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLinkInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLinkInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLinkInfoResponse) GetBody() *AccountLinkInfo {
	return s.Body
}

func (s *GetLinkInfoResponse) SetHeaders(v map[string]*string) *GetLinkInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLinkInfoResponse) SetStatusCode(v int32) *GetLinkInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLinkInfoResponse) SetBody(v *AccountLinkInfo) *GetLinkInfoResponse {
	s.Body = v
	return s
}

func (s *GetLinkInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
