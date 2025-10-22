// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIntentionForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitIntentionForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitIntentionForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *SubmitIntentionForPartnerResponseBody) *SubmitIntentionForPartnerResponse
	GetBody() *SubmitIntentionForPartnerResponseBody
}

type SubmitIntentionForPartnerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIntentionForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIntentionForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitIntentionForPartnerResponse) GoString() string {
	return s.String()
}

func (s *SubmitIntentionForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitIntentionForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitIntentionForPartnerResponse) GetBody() *SubmitIntentionForPartnerResponseBody {
	return s.Body
}

func (s *SubmitIntentionForPartnerResponse) SetHeaders(v map[string]*string) *SubmitIntentionForPartnerResponse {
	s.Headers = v
	return s
}

func (s *SubmitIntentionForPartnerResponse) SetStatusCode(v int32) *SubmitIntentionForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIntentionForPartnerResponse) SetBody(v *SubmitIntentionForPartnerResponseBody) *SubmitIntentionForPartnerResponse {
	s.Body = v
	return s
}

func (s *SubmitIntentionForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
