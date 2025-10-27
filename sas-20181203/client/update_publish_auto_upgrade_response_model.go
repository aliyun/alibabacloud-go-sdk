// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishAutoUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePublishAutoUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePublishAutoUpgradeResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePublishAutoUpgradeResponseBody) *UpdatePublishAutoUpgradeResponse
	GetBody() *UpdatePublishAutoUpgradeResponseBody
}

type UpdatePublishAutoUpgradeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublishAutoUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublishAutoUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishAutoUpgradeResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublishAutoUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePublishAutoUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePublishAutoUpgradeResponse) GetBody() *UpdatePublishAutoUpgradeResponseBody {
	return s.Body
}

func (s *UpdatePublishAutoUpgradeResponse) SetHeaders(v map[string]*string) *UpdatePublishAutoUpgradeResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublishAutoUpgradeResponse) SetStatusCode(v int32) *UpdatePublishAutoUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublishAutoUpgradeResponse) SetBody(v *UpdatePublishAutoUpgradeResponseBody) *UpdatePublishAutoUpgradeResponse {
	s.Body = v
	return s
}

func (s *UpdatePublishAutoUpgradeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
