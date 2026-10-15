// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveChannelDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveChannelDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveChannelDraftResponse
	GetStatusCode() *int32
	SetBody(v *SaveChannelDraftResponseBody) *SaveChannelDraftResponse
	GetBody() *SaveChannelDraftResponseBody
}

type SaveChannelDraftResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveChannelDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveChannelDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveChannelDraftResponse) GoString() string {
	return s.String()
}

func (s *SaveChannelDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveChannelDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveChannelDraftResponse) GetBody() *SaveChannelDraftResponseBody {
	return s.Body
}

func (s *SaveChannelDraftResponse) SetHeaders(v map[string]*string) *SaveChannelDraftResponse {
	s.Headers = v
	return s
}

func (s *SaveChannelDraftResponse) SetStatusCode(v int32) *SaveChannelDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveChannelDraftResponse) SetBody(v *SaveChannelDraftResponseBody) *SaveChannelDraftResponse {
	s.Body = v
	return s
}

func (s *SaveChannelDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
