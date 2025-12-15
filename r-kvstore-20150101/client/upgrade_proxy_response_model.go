// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeProxyResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeProxyResponseBody) *UpgradeProxyResponse
	GetBody() *UpgradeProxyResponseBody
}

type UpgradeProxyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeProxyResponse) GoString() string {
	return s.String()
}

func (s *UpgradeProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeProxyResponse) GetBody() *UpgradeProxyResponseBody {
	return s.Body
}

func (s *UpgradeProxyResponse) SetHeaders(v map[string]*string) *UpgradeProxyResponse {
	s.Headers = v
	return s
}

func (s *UpgradeProxyResponse) SetStatusCode(v int32) *UpgradeProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeProxyResponse) SetBody(v *UpgradeProxyResponseBody) *UpgradeProxyResponse {
	s.Body = v
	return s
}

func (s *UpgradeProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
