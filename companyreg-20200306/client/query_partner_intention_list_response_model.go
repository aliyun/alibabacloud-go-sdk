// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPartnerIntentionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPartnerIntentionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPartnerIntentionListResponse
	GetStatusCode() *int32
	SetBody(v *QueryPartnerIntentionListResponseBody) *QueryPartnerIntentionListResponse
	GetBody() *QueryPartnerIntentionListResponseBody
}

type QueryPartnerIntentionListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPartnerIntentionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPartnerIntentionListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPartnerIntentionListResponse) GoString() string {
	return s.String()
}

func (s *QueryPartnerIntentionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPartnerIntentionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPartnerIntentionListResponse) GetBody() *QueryPartnerIntentionListResponseBody {
	return s.Body
}

func (s *QueryPartnerIntentionListResponse) SetHeaders(v map[string]*string) *QueryPartnerIntentionListResponse {
	s.Headers = v
	return s
}

func (s *QueryPartnerIntentionListResponse) SetStatusCode(v int32) *QueryPartnerIntentionListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPartnerIntentionListResponse) SetBody(v *QueryPartnerIntentionListResponseBody) *QueryPartnerIntentionListResponse {
	s.Body = v
	return s
}

func (s *QueryPartnerIntentionListResponse) Validate() error {
	return dara.Validate(s)
}
