// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainDetail(v *DescribeVodDomainDetailResponseBodyDomainDetail) *DescribeVodDomainDetailResponseBody
	GetDomainDetail() *DescribeVodDomainDetailResponseBodyDomainDetail
	SetRequestId(v string) *DescribeVodDomainDetailResponseBody
	GetRequestId() *string
}

type DescribeVodDomainDetailResponseBody struct {
	// The basic configuration information of the domain name.
	DomainDetail *DescribeVodDomainDetailResponseBodyDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 09ABE829-6CD3-4FE0-556113E2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBody) GetDomainDetail() *DescribeVodDomainDetailResponseBodyDomainDetail {
	return s.DomainDetail
}

func (s *DescribeVodDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainDetailResponseBody) SetDomainDetail(v *DescribeVodDomainDetailResponseBodyDomainDetail) *DescribeVodDomainDetailResponseBody {
	s.DomainDetail = v
	return s
}

func (s *DescribeVodDomainDetailResponseBody) SetRequestId(v string) *DescribeVodDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBody) Validate() error {
	if s.DomainDetail != nil {
		if err := s.DomainDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainDetailResponseBodyDomainDetail struct {
	// The name of the HTTPS certificate. This parameter is returned only if HTTPS secure acceleration is enabled.
	//
	// example:
	//
	// testCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The CNAME assigned to the accelerated domain name. You must add a CNAME record with your DNS provider to map the accelerated domain name to this CNAME. For more information, see [Configure a CNAME record](https://help.aliyun.com/document_detail/86075.html).
	//
	// example:
	//
	// example.com.w.alikunlun.net
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The description of the VOD acceleration domain name.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The VOD acceleration domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the accelerated domain name. Valid values:
	//
	// - **online**: enabled.
	//
	// - **offline**: disabled.
	//
	// - **configuring**: being configured.
	//
	// - **configure_failed**: configuration failed.
	//
	// - **checking**: being reviewed.
	//
	// - **check_failed**: review failed.
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the domain name was created. The time follows the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-27T06:51:26Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the domain name was last modified. The time follows the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-27T06:55:26Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the SSL certificate is enabled. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The public key of the HTTPS certificate. This parameter is returned only if HTTPS secure acceleration is enabled.
	//
	// example:
	//
	// yourSSLPub
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// The acceleration region. Valid values:
	//
	// - **domestic*	- (default): the Chinese mainland only.
	//
	// - **overseas**: global (excluding the Chinese mainland).
	//
	// - **global**: global.
	//
	// example:
	//
	// domestic
	Scope   *string                                                 `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Sources *DescribeVodDomainDetailResponseBodyDomainDetailSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	// The back-to-origin weight.
	//
	// example:
	//
	// 1
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVodDomainDetailResponseBodyDomainDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBodyDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetCname() *string {
	return s.Cname
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetDescription() *string {
	return s.Description
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetScope() *string {
	return s.Scope
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetSources() *DescribeVodDomainDetailResponseBodyDomainDetailSources {
	return s.Sources
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetWeight() *string {
	return s.Weight
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetCertName(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.CertName = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetCname(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetDescription(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetDomainName(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetDomainStatus(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetGmtCreated(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetGmtModified(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetSSLProtocol(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetSSLPub(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetScope(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetSources(v *DescribeVodDomainDetailResponseBodyDomainDetailSources) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Sources = v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetWeight(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Weight = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) Validate() error {
	if s.Sources != nil {
		if err := s.Sources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainDetailResponseBodyDomainDetailSources struct {
	Source []*DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSources) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSources) GetSource() []*DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	return s.Source
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSources) SetSource(v []*DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) *DescribeVodDomainDetailResponseBodyDomainDetailSources {
	s.Source = v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSources) Validate() error {
	if s.Source != nil {
		for _, item := range s.Source {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Enabled  *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetEnabled() *string {
	return s.Enabled
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetContent(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetEnabled(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Enabled = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetPort(v int32) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetPriority(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetType(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetWeight(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) Validate() error {
	return dara.Validate(s)
}
