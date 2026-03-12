// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingContactByTemplateIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForUpdatingContactByTemplateIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForUpdatingContactByTemplateIdResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForUpdatingContactByTemplateIdResponseBody) *SaveTaskForUpdatingContactByTemplateIdResponse
	GetBody() *SaveTaskForUpdatingContactByTemplateIdResponseBody
}

type SaveTaskForUpdatingContactByTemplateIdResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForUpdatingContactByTemplateIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForUpdatingContactByTemplateIdResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingContactByTemplateIdResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponse) GetBody() *SaveTaskForUpdatingContactByTemplateIdResponseBody {
	return s.Body
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponse) SetHeaders(v map[string]*string) *SaveTaskForUpdatingContactByTemplateIdResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponse) SetStatusCode(v int32) *SaveTaskForUpdatingContactByTemplateIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponse) SetBody(v *SaveTaskForUpdatingContactByTemplateIdResponseBody) *SaveTaskForUpdatingContactByTemplateIdResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
