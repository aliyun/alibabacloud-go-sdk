// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCampaignsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCampaignsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCampaignsResponse
	GetStatusCode() *int32
	SetBody(v *ListCampaignsResponseBody) *ListCampaignsResponse
	GetBody() *ListCampaignsResponseBody
}

type ListCampaignsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCampaignsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCampaignsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignsResponse) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCampaignsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCampaignsResponse) GetBody() *ListCampaignsResponseBody {
	return s.Body
}

func (s *ListCampaignsResponse) SetHeaders(v map[string]*string) *ListCampaignsResponse {
	s.Headers = v
	return s
}

func (s *ListCampaignsResponse) SetStatusCode(v int32) *ListCampaignsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCampaignsResponse) SetBody(v *ListCampaignsResponseBody) *ListCampaignsResponse {
	s.Body = v
	return s
}

func (s *ListCampaignsResponse) Validate() error {
	return dara.Validate(s)
}
