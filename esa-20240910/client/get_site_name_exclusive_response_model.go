// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteNameExclusiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSiteNameExclusiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSiteNameExclusiveResponse
	GetStatusCode() *int32
	SetBody(v *GetSiteNameExclusiveResponseBody) *GetSiteNameExclusiveResponse
	GetBody() *GetSiteNameExclusiveResponseBody
}

type GetSiteNameExclusiveResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteNameExclusiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteNameExclusiveResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSiteNameExclusiveResponse) GoString() string {
	return s.String()
}

func (s *GetSiteNameExclusiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSiteNameExclusiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSiteNameExclusiveResponse) GetBody() *GetSiteNameExclusiveResponseBody {
	return s.Body
}

func (s *GetSiteNameExclusiveResponse) SetHeaders(v map[string]*string) *GetSiteNameExclusiveResponse {
	s.Headers = v
	return s
}

func (s *GetSiteNameExclusiveResponse) SetStatusCode(v int32) *GetSiteNameExclusiveResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteNameExclusiveResponse) SetBody(v *GetSiteNameExclusiveResponseBody) *GetSiteNameExclusiveResponse {
	s.Body = v
	return s
}

func (s *GetSiteNameExclusiveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
