// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityAlertOfAllRuleScopeByWatchIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityAlertOfAllRuleScopeByWatchIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityAlertOfAllRuleScopeByWatchIdResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) *GetQualityAlertOfAllRuleScopeByWatchIdResponse
	GetBody() *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody
}

type GetQualityAlertOfAllRuleScopeByWatchIdResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdResponse) GoString() string {
	return s.String()
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponse) GetBody() *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody {
	return s.Body
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponse) SetHeaders(v map[string]*string) *GetQualityAlertOfAllRuleScopeByWatchIdResponse {
	s.Headers = v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponse) SetStatusCode(v int32) *GetQualityAlertOfAllRuleScopeByWatchIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponse) SetBody(v *GetQualityAlertOfAllRuleScopeByWatchIdResponseBody) *GetQualityAlertOfAllRuleScopeByWatchIdResponse {
	s.Body = v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
