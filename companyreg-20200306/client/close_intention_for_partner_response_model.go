// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseIntentionForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseIntentionForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseIntentionForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *CloseIntentionForPartnerResponseBody) *CloseIntentionForPartnerResponse
	GetBody() *CloseIntentionForPartnerResponseBody
}

type CloseIntentionForPartnerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseIntentionForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseIntentionForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseIntentionForPartnerResponse) GoString() string {
	return s.String()
}

func (s *CloseIntentionForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseIntentionForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseIntentionForPartnerResponse) GetBody() *CloseIntentionForPartnerResponseBody {
	return s.Body
}

func (s *CloseIntentionForPartnerResponse) SetHeaders(v map[string]*string) *CloseIntentionForPartnerResponse {
	s.Headers = v
	return s
}

func (s *CloseIntentionForPartnerResponse) SetStatusCode(v int32) *CloseIntentionForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseIntentionForPartnerResponse) SetBody(v *CloseIntentionForPartnerResponseBody) *CloseIntentionForPartnerResponse {
	s.Body = v
	return s
}

func (s *CloseIntentionForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
