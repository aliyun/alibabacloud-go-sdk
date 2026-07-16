// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageShieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *GetPageShieldResponseBody
	GetEnable() *string
	SetReportUri(v string) *GetPageShieldResponseBody
	GetReportUri() *string
	SetRequestId(v string) *GetPageShieldResponseBody
	GetRequestId() *string
	SetSiteVersion(v int32) *GetPageShieldResponseBody
	GetSiteVersion() *int32
}

type GetPageShieldResponseBody struct {
	// The switch status. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The report URI.
	//
	// example:
	//
	// /test
	ReportUri *string `json:"ReportUri,omitempty" xml:"ReportUri,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version number of the site. For sites with version management enabled, you can use this parameter to specify the site version on which the configuration takes effect. The default value is 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetPageShieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPageShieldResponseBody) GoString() string {
	return s.String()
}

func (s *GetPageShieldResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetPageShieldResponseBody) GetReportUri() *string {
	return s.ReportUri
}

func (s *GetPageShieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPageShieldResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetPageShieldResponseBody) SetEnable(v string) *GetPageShieldResponseBody {
	s.Enable = &v
	return s
}

func (s *GetPageShieldResponseBody) SetReportUri(v string) *GetPageShieldResponseBody {
	s.ReportUri = &v
	return s
}

func (s *GetPageShieldResponseBody) SetRequestId(v string) *GetPageShieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPageShieldResponseBody) SetSiteVersion(v int32) *GetPageShieldResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetPageShieldResponseBody) Validate() error {
	return dara.Validate(s)
}
