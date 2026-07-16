// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteTrafficSequenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSiteTrafficSequenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSiteTrafficSequenceResponse
	GetStatusCode() *int32
	SetBody(v *GetSiteTrafficSequenceResponseBody) *GetSiteTrafficSequenceResponse
	GetBody() *GetSiteTrafficSequenceResponseBody
}

type GetSiteTrafficSequenceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteTrafficSequenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteTrafficSequenceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSiteTrafficSequenceResponse) GoString() string {
	return s.String()
}

func (s *GetSiteTrafficSequenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSiteTrafficSequenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSiteTrafficSequenceResponse) GetBody() *GetSiteTrafficSequenceResponseBody {
	return s.Body
}

func (s *GetSiteTrafficSequenceResponse) SetHeaders(v map[string]*string) *GetSiteTrafficSequenceResponse {
	s.Headers = v
	return s
}

func (s *GetSiteTrafficSequenceResponse) SetStatusCode(v int32) *GetSiteTrafficSequenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteTrafficSequenceResponse) SetBody(v *GetSiteTrafficSequenceResponseBody) *GetSiteTrafficSequenceResponse {
	s.Body = v
	return s
}

func (s *GetSiteTrafficSequenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
