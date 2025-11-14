// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCampaignNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCampaignNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCampaignNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCampaignNumbersResponseBody) *ModifyCampaignNumbersResponse
	GetBody() *ModifyCampaignNumbersResponseBody
}

type ModifyCampaignNumbersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCampaignNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCampaignNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCampaignNumbersResponse) GoString() string {
	return s.String()
}

func (s *ModifyCampaignNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCampaignNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCampaignNumbersResponse) GetBody() *ModifyCampaignNumbersResponseBody {
	return s.Body
}

func (s *ModifyCampaignNumbersResponse) SetHeaders(v map[string]*string) *ModifyCampaignNumbersResponse {
	s.Headers = v
	return s
}

func (s *ModifyCampaignNumbersResponse) SetStatusCode(v int32) *ModifyCampaignNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCampaignNumbersResponse) SetBody(v *ModifyCampaignNumbersResponseBody) *ModifyCampaignNumbersResponse {
	s.Body = v
	return s
}

func (s *ModifyCampaignNumbersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
