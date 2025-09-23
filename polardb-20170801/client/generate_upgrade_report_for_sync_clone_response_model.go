// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUpgradeReportForSyncCloneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateUpgradeReportForSyncCloneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateUpgradeReportForSyncCloneResponse
	GetStatusCode() *int32
	SetBody(v *GenerateUpgradeReportForSyncCloneResponseBody) *GenerateUpgradeReportForSyncCloneResponse
	GetBody() *GenerateUpgradeReportForSyncCloneResponseBody
}

type GenerateUpgradeReportForSyncCloneResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUpgradeReportForSyncCloneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUpgradeReportForSyncCloneResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateUpgradeReportForSyncCloneResponse) GoString() string {
	return s.String()
}

func (s *GenerateUpgradeReportForSyncCloneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateUpgradeReportForSyncCloneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateUpgradeReportForSyncCloneResponse) GetBody() *GenerateUpgradeReportForSyncCloneResponseBody {
	return s.Body
}

func (s *GenerateUpgradeReportForSyncCloneResponse) SetHeaders(v map[string]*string) *GenerateUpgradeReportForSyncCloneResponse {
	s.Headers = v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneResponse) SetStatusCode(v int32) *GenerateUpgradeReportForSyncCloneResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneResponse) SetBody(v *GenerateUpgradeReportForSyncCloneResponseBody) *GenerateUpgradeReportForSyncCloneResponse {
	s.Body = v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneResponse) Validate() error {
	return dara.Validate(s)
}
