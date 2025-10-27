// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerhubImageRiskRankInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDockerhubImageRiskRankInfoResponseBody
	GetRequestId() *string
	SetRiskRankInfo(v *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo) *GetDockerhubImageRiskRankInfoResponseBody
	GetRiskRankInfo() *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo
}

type GetDockerhubImageRiskRankInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30CBF632-109F-596F-97F2-451C8B2A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The risk information.
	RiskRankInfo *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo `json:"RiskRankInfo,omitempty" xml:"RiskRankInfo,omitempty" type:"Struct"`
}

func (s GetDockerhubImageRiskRankInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskRankInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskRankInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDockerhubImageRiskRankInfoResponseBody) GetRiskRankInfo() *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo {
	return s.RiskRankInfo
}

func (s *GetDockerhubImageRiskRankInfoResponseBody) SetRequestId(v string) *GetDockerhubImageRiskRankInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBody) SetRiskRankInfo(v *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo) *GetDockerhubImageRiskRankInfoResponseBody {
	s.RiskRankInfo = v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBody) Validate() error {
	if s.RiskRankInfo != nil {
		if err := s.RiskRankInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo struct {
	// The baseline risks.
	Baseline []*GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline `json:"Baseline,omitempty" xml:"Baseline,omitempty" type:"Repeated"`
	// The risk information of high-risk vulnerabilities.
	VulAsap []*GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap `json:"VulAsap,omitempty" xml:"VulAsap,omitempty" type:"Repeated"`
}

func (s GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo) GetBaseline() []*GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	return s.Baseline
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo) GetVulAsap() []*GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	return s.VulAsap
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo) SetBaseline(v []*GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo {
	s.Baseline = v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo) SetVulAsap(v []*GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo {
	s.VulAsap = v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfo) Validate() error {
	if s.Baseline != nil {
		for _, item := range s.Baseline {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VulAsap != nil {
		for _, item := range s.VulAsap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline struct {
	// The digest value of the image.
	//
	// example:
	//
	// f28ecca63bfaf22ead4b28b63d752a21e4d2c1de90b8549fbde880d619f3****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The number of baseline risks detected on the image repository.
	//
	// example:
	//
	// 1
	HcCount *int32 `json:"HcCount,omitempty" xml:"HcCount,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// 7c5ad02865aef575387a05bb3c81b27e0d8ed1f2e3f722ea05523b72882f****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image size.
	//
	// example:
	//
	// 1024
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The name of the Container Registry repository.
	//
	// example:
	//
	// glz123
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the repository belongs.
	//
	// example:
	//
	// namespace-01
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The risk statistics of all hosts, images, and containers.
	//
	// example:
	//
	// {\\"account\\":0,\\"agentlessAll\\":0,\\"agentlessBaseline\\":0,\\"agentlessMalicious\\":0,\\"agentlessSensitiveFile\\":0,\\"agentlessVulCve\\":0,\\"agentlessVulSca\\":0,\\"agentlessVulSys\\":0,\\"appNum\\":0,\\"asapVulCount\\":0,\\"baselineHigh\\":0,\\"baselineLow\\":0,\\"baselineMedium\\":0,\\"baselineNum\\":0,\\"cmsNum\\":0,\\"containerAsap\\":0,\\"containerLater\\":0,\\"containerNntf\\":0,\\"containerRemind\\":0,\\"containerSerious\\":0,\\"containerSuspicious\\":0,\\"cveNum\\":0,\\"emgNum\\":0,\\"health\\":0,\\"imageBaselineHigh\\":1,\\"imageBaselineLow\\":0,\\"imageBaselineMedium\\":0,\\"imageBaselineNum\\":1,\\"imageMaliciousFileRemind\\":0,\\"imageMaliciousFileSerious\\":0,\\"imageMaliciousFileSuspicious\\":0,\\"imageVulAsap\\":0,\\"imageVulLater\\":0,\\"imageVulNntf\\":0,\\"laterVulCount\\":0,\\"newSuspicious\\":0,\\"nntfVulCount\\":0,\\"remindNum\\":0,\\"scaNum\\":0,\\"seriousNum\\":0,\\"suspNum\\":0,\\"suspicious\\":0,\\"sysNum\\":0,\\"trojan\\":0,\\"uuid\\":\\"009635bf00c0585e3122ab92f5449919\\",\\"vul\\":0,\\"weakPWNum\\":0}
	RiskLevelDetail *string `json:"RiskLevelDetail,omitempty" xml:"RiskLevelDetail,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// machineResource
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the record.
	//
	// example:
	//
	// 5583aa03-922e-4709-a888-389f2489****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The number of detected vulnerabilities.
	//
	// example:
	//
	// 1
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetDigest() *string {
	return s.Digest
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetHcCount() *int32 {
	return s.HcCount
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetImageId() *string {
	return s.ImageId
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetImageSize() *int64 {
	return s.ImageSize
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetRepoName() *string {
	return s.RepoName
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetRiskLevelDetail() *string {
	return s.RiskLevelDetail
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetTag() *string {
	return s.Tag
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetUuid() *string {
	return s.Uuid
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) GetVulCount() *int32 {
	return s.VulCount
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetDigest(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.Digest = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetHcCount(v int32) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.HcCount = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetImageId(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.ImageId = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetImageSize(v int64) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.ImageSize = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetRepoName(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.RepoName = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetRepoNamespace(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.RepoNamespace = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetRiskLevelDetail(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.RiskLevelDetail = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetTag(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.Tag = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetUuid(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.Uuid = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) SetVulCount(v int32) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline {
	s.VulCount = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoBaseline) Validate() error {
	return dara.Validate(s)
}

type GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap struct {
	// The digest value of the image.
	//
	// example:
	//
	// d97c1348e56eb52902888e6e5673623321b1f19ac45ed532c3114dc0f989****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The number of risks detected on the image repository.
	//
	// example:
	//
	// 0
	HcCount *int32 `json:"HcCount,omitempty" xml:"HcCount,omitempty"`
	// The image ID.
	//
	// example:
	//
	// fabe4203a89765a2c99554040bda51eac7885a18216f4ac0be82710cff60****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image size.
	//
	// example:
	//
	// 1024
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// lkl-zf-ss-ordapi-labs
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// namespace-01
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The risk statistics of all hosts, images, and containers.
	//
	// example:
	//
	// {\\"account\\":0,\\"agentlessAll\\":0,\\"agentlessBaseline\\":0,\\"agentlessMalicious\\":0,\\"agentlessSensitiveFile\\":0,\\"agentlessVulCve\\":0,\\"agentlessVulSca\\":0,\\"agentlessVulSys\\":0,\\"appNum\\":0,\\"asapVulCount\\":0,\\"baselineHigh\\":0,\\"baselineLow\\":0,\\"baselineMedium\\":0,\\"baselineNum\\":0,\\"cmsNum\\":0,\\"containerAsap\\":0,\\"containerLater\\":0,\\"containerNntf\\":0,\\"containerRemind\\":0,\\"containerSerious\\":0,\\"containerSuspicious\\":0,\\"cveNum\\":513,\\"emgNum\\":0,\\"health\\":0,\\"imageBaselineHigh\\":0,\\"imageBaselineLow\\":0,\\"imageBaselineMedium\\":0,\\"imageBaselineNum\\":0,\\"imageMaliciousFileRemind\\":0,\\"imageMaliciousFileSerious\\":0,\\"imageMaliciousFileSuspicious\\":0,\\"imageVulAsap\\":3,\\"imageVulLater\\":182,\\"imageVulNntf\\":328,\\"laterVulCount\\":0,\\"newSuspicious\\":0,\\"nntfVulCount\\":0,\\"remindNum\\":0,\\"scaNum\\":0,\\"seriousNum\\":0,\\"suspNum\\":0,\\"suspicious\\":0,\\"sysNum\\":0,\\"trojan\\":0,\\"uuid\\":\\"44ffb29d6f66d8509598bcdfa660a21d\\",\\"vul\\":513,\\"weakPWNum\\":0}
	RiskLevelDetail *string `json:"RiskLevelDetail,omitempty" xml:"RiskLevelDetail,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// app:app01-ubuntu
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the record.
	//
	// example:
	//
	// 6636c286-8063-4c97-8508-6aaf16a8****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The number of vulnerabilities.
	//
	// example:
	//
	// 1
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetDigest() *string {
	return s.Digest
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetHcCount() *int32 {
	return s.HcCount
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetImageId() *string {
	return s.ImageId
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetImageSize() *int64 {
	return s.ImageSize
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetRepoName() *string {
	return s.RepoName
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetRiskLevelDetail() *string {
	return s.RiskLevelDetail
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetTag() *string {
	return s.Tag
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetUuid() *string {
	return s.Uuid
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) GetVulCount() *int32 {
	return s.VulCount
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetDigest(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.Digest = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetHcCount(v int32) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.HcCount = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetImageId(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.ImageId = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetImageSize(v int64) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.ImageSize = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetRepoName(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.RepoName = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetRepoNamespace(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.RepoNamespace = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetRiskLevelDetail(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.RiskLevelDetail = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetTag(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.Tag = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetUuid(v string) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.Uuid = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) SetVulCount(v int32) *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap {
	s.VulCount = &v
	return s
}

func (s *GetDockerhubImageRiskRankInfoResponseBodyRiskRankInfoVulAsap) Validate() error {
	return dara.Validate(s)
}
