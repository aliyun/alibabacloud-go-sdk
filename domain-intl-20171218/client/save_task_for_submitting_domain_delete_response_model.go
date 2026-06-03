// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForSubmittingDomainDeleteResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForSubmittingDomainDeleteResponseBody) *SaveTaskForSubmittingDomainDeleteResponse
	GetBody() *SaveTaskForSubmittingDomainDeleteResponseBody
}

type SaveTaskForSubmittingDomainDeleteResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForSubmittingDomainDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForSubmittingDomainDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainDeleteResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) GetBody() *SaveTaskForSubmittingDomainDeleteResponseBody {
	return s.Body
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainDeleteResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) SetStatusCode(v int32) *SaveTaskForSubmittingDomainDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) SetBody(v *SaveTaskForSubmittingDomainDeleteResponseBody) *SaveTaskForSubmittingDomainDeleteResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
