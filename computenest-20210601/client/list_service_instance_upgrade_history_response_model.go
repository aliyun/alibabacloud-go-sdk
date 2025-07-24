// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceUpgradeHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceInstanceUpgradeHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceInstanceUpgradeHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceInstanceUpgradeHistoryResponseBody) *ListServiceInstanceUpgradeHistoryResponse
	GetBody() *ListServiceInstanceUpgradeHistoryResponseBody
}

type ListServiceInstanceUpgradeHistoryResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceUpgradeHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceUpgradeHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceInstanceUpgradeHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceInstanceUpgradeHistoryResponse) GetBody() *ListServiceInstanceUpgradeHistoryResponseBody {
	return s.Body
}

func (s *ListServiceInstanceUpgradeHistoryResponse) SetHeaders(v map[string]*string) *ListServiceInstanceUpgradeHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponse) SetStatusCode(v int32) *ListServiceInstanceUpgradeHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponse) SetBody(v *ListServiceInstanceUpgradeHistoryResponseBody) *ListServiceInstanceUpgradeHistoryResponse {
	s.Body = v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponse) Validate() error {
	return dara.Validate(s)
}
