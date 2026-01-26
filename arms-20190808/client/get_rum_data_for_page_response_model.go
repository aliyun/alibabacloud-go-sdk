// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumDataForPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRumDataForPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRumDataForPageResponse
	GetStatusCode() *int32
	SetBody(v *GetRumDataForPageResponseBody) *GetRumDataForPageResponse
	GetBody() *GetRumDataForPageResponseBody
}

type GetRumDataForPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRumDataForPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRumDataForPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRumDataForPageResponse) GoString() string {
	return s.String()
}

func (s *GetRumDataForPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRumDataForPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRumDataForPageResponse) GetBody() *GetRumDataForPageResponseBody {
	return s.Body
}

func (s *GetRumDataForPageResponse) SetHeaders(v map[string]*string) *GetRumDataForPageResponse {
	s.Headers = v
	return s
}

func (s *GetRumDataForPageResponse) SetStatusCode(v int32) *GetRumDataForPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRumDataForPageResponse) SetBody(v *GetRumDataForPageResponseBody) *GetRumDataForPageResponse {
	s.Body = v
	return s
}

func (s *GetRumDataForPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
