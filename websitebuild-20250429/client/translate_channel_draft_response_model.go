// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateChannelDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranslateChannelDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranslateChannelDraftResponse
	GetStatusCode() *int32
	SetBody(v *TranslateChannelDraftResponseBody) *TranslateChannelDraftResponse
	GetBody() *TranslateChannelDraftResponseBody
}

type TranslateChannelDraftResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateChannelDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateChannelDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s TranslateChannelDraftResponse) GoString() string {
	return s.String()
}

func (s *TranslateChannelDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranslateChannelDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranslateChannelDraftResponse) GetBody() *TranslateChannelDraftResponseBody {
	return s.Body
}

func (s *TranslateChannelDraftResponse) SetHeaders(v map[string]*string) *TranslateChannelDraftResponse {
	s.Headers = v
	return s
}

func (s *TranslateChannelDraftResponse) SetStatusCode(v int32) *TranslateChannelDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateChannelDraftResponse) SetBody(v *TranslateChannelDraftResponseBody) *TranslateChannelDraftResponse {
	s.Body = v
	return s
}

func (s *TranslateChannelDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
