// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainMultiStreamConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveDomainMultiStreamConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveDomainMultiStreamConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveDomainMultiStreamConfigResponseBody) *SetLiveDomainMultiStreamConfigResponse
	GetBody() *SetLiveDomainMultiStreamConfigResponseBody
}

type SetLiveDomainMultiStreamConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveDomainMultiStreamConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveDomainMultiStreamConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainMultiStreamConfigResponse) GoString() string {
	return s.String()
}

func (s *SetLiveDomainMultiStreamConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveDomainMultiStreamConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveDomainMultiStreamConfigResponse) GetBody() *SetLiveDomainMultiStreamConfigResponseBody {
	return s.Body
}

func (s *SetLiveDomainMultiStreamConfigResponse) SetHeaders(v map[string]*string) *SetLiveDomainMultiStreamConfigResponse {
	s.Headers = v
	return s
}

func (s *SetLiveDomainMultiStreamConfigResponse) SetStatusCode(v int32) *SetLiveDomainMultiStreamConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveDomainMultiStreamConfigResponse) SetBody(v *SetLiveDomainMultiStreamConfigResponseBody) *SetLiveDomainMultiStreamConfigResponse {
	s.Body = v
	return s
}

func (s *SetLiveDomainMultiStreamConfigResponse) Validate() error {
	return dara.Validate(s)
}
