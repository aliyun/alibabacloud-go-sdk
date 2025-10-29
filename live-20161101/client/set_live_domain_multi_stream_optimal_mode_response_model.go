// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainMultiStreamOptimalModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveDomainMultiStreamOptimalModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveDomainMultiStreamOptimalModeResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveDomainMultiStreamOptimalModeResponseBody) *SetLiveDomainMultiStreamOptimalModeResponse
	GetBody() *SetLiveDomainMultiStreamOptimalModeResponseBody
}

type SetLiveDomainMultiStreamOptimalModeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveDomainMultiStreamOptimalModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveDomainMultiStreamOptimalModeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainMultiStreamOptimalModeResponse) GoString() string {
	return s.String()
}

func (s *SetLiveDomainMultiStreamOptimalModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveDomainMultiStreamOptimalModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveDomainMultiStreamOptimalModeResponse) GetBody() *SetLiveDomainMultiStreamOptimalModeResponseBody {
	return s.Body
}

func (s *SetLiveDomainMultiStreamOptimalModeResponse) SetHeaders(v map[string]*string) *SetLiveDomainMultiStreamOptimalModeResponse {
	s.Headers = v
	return s
}

func (s *SetLiveDomainMultiStreamOptimalModeResponse) SetStatusCode(v int32) *SetLiveDomainMultiStreamOptimalModeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveDomainMultiStreamOptimalModeResponse) SetBody(v *SetLiveDomainMultiStreamOptimalModeResponseBody) *SetLiveDomainMultiStreamOptimalModeResponse {
	s.Body = v
	return s
}

func (s *SetLiveDomainMultiStreamOptimalModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
