// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainMultiStreamMasterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveDomainMultiStreamMasterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveDomainMultiStreamMasterResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveDomainMultiStreamMasterResponseBody) *SetLiveDomainMultiStreamMasterResponse
	GetBody() *SetLiveDomainMultiStreamMasterResponseBody
}

type SetLiveDomainMultiStreamMasterResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveDomainMultiStreamMasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveDomainMultiStreamMasterResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainMultiStreamMasterResponse) GoString() string {
	return s.String()
}

func (s *SetLiveDomainMultiStreamMasterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveDomainMultiStreamMasterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveDomainMultiStreamMasterResponse) GetBody() *SetLiveDomainMultiStreamMasterResponseBody {
	return s.Body
}

func (s *SetLiveDomainMultiStreamMasterResponse) SetHeaders(v map[string]*string) *SetLiveDomainMultiStreamMasterResponse {
	s.Headers = v
	return s
}

func (s *SetLiveDomainMultiStreamMasterResponse) SetStatusCode(v int32) *SetLiveDomainMultiStreamMasterResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveDomainMultiStreamMasterResponse) SetBody(v *SetLiveDomainMultiStreamMasterResponseBody) *SetLiveDomainMultiStreamMasterResponse {
	s.Body = v
	return s
}

func (s *SetLiveDomainMultiStreamMasterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
