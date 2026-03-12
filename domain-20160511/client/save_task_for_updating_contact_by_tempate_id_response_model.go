// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingContactByTempateIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForUpdatingContactByTempateIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForUpdatingContactByTempateIdResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForUpdatingContactByTempateIdResponseBody) *SaveTaskForUpdatingContactByTempateIdResponse
	GetBody() *SaveTaskForUpdatingContactByTempateIdResponseBody
}

type SaveTaskForUpdatingContactByTempateIdResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForUpdatingContactByTempateIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForUpdatingContactByTempateIdResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingContactByTempateIdResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingContactByTempateIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForUpdatingContactByTempateIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForUpdatingContactByTempateIdResponse) GetBody() *SaveTaskForUpdatingContactByTempateIdResponseBody {
	return s.Body
}

func (s *SaveTaskForUpdatingContactByTempateIdResponse) SetHeaders(v map[string]*string) *SaveTaskForUpdatingContactByTempateIdResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdResponse) SetStatusCode(v int32) *SaveTaskForUpdatingContactByTempateIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdResponse) SetBody(v *SaveTaskForUpdatingContactByTempateIdResponseBody) *SaveTaskForUpdatingContactByTempateIdResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
