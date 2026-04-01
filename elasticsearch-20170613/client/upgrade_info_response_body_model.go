// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeInfoResponseBody
	GetRequestId() *string
	SetResult(v *UpgradeInfoResponseBodyResult) *UpgradeInfoResponseBody
	GetResult() *UpgradeInfoResponseBodyResult
}

type UpgradeInfoResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpgradeInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpgradeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeInfoResponseBody) GetResult() *UpgradeInfoResponseBodyResult {
	return s.Result
}

func (s *UpgradeInfoResponseBody) SetRequestId(v string) *UpgradeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeInfoResponseBody) SetResult(v *UpgradeInfoResponseBodyResult) *UpgradeInfoResponseBody {
	s.Result = v
	return s
}

func (s *UpgradeInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeInfoResponseBodyResult struct {
	UpgradeInfo *UpgradeInfoResponseBodyResultUpgradeInfo `json:"UpgradeInfo,omitempty" xml:"UpgradeInfo,omitempty" type:"Struct"`
}

func (s UpgradeInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpgradeInfoResponseBodyResult) GetUpgradeInfo() *UpgradeInfoResponseBodyResultUpgradeInfo {
	return s.UpgradeInfo
}

func (s *UpgradeInfoResponseBodyResult) SetUpgradeInfo(v *UpgradeInfoResponseBodyResultUpgradeInfo) *UpgradeInfoResponseBodyResult {
	s.UpgradeInfo = v
	return s
}

func (s *UpgradeInfoResponseBodyResult) Validate() error {
	if s.UpgradeInfo != nil {
		if err := s.UpgradeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeInfoResponseBodyResultUpgradeInfo struct {
	// example:
	//
	// 1.7.3
	CurRepoVersion *string `json:"CurRepoVersion,omitempty" xml:"CurRepoVersion,omitempty"`
	// example:
	//
	// 1.7.3
	UpdateRepoVersion *string `json:"UpdateRepoVersion,omitempty" xml:"UpdateRepoVersion,omitempty"`
	// example:
	//
	// false
	Upgrade *bool `json:"Upgrade,omitempty" xml:"Upgrade,omitempty"`
	// example:
	//
	// 2.2.4
	CurApackVersion *string `json:"curApackVersion,omitempty" xml:"curApackVersion,omitempty"`
	// example:
	//
	// 8.17.0
	CurEsVersion *string `json:"curEsVersion,omitempty" xml:"curEsVersion,omitempty"`
	// example:
	//
	// 2.2.4
	UpgradeApackVersion *string `json:"upgradeApackVersion,omitempty" xml:"upgradeApackVersion,omitempty"`
	// example:
	//
	// 8.17.0
	UpgradeEsVersion *string `json:"upgradeEsVersion,omitempty" xml:"upgradeEsVersion,omitempty"`
}

func (s UpgradeInfoResponseBodyResultUpgradeInfo) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInfoResponseBodyResultUpgradeInfo) GoString() string {
	return s.String()
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) GetCurRepoVersion() *string {
	return s.CurRepoVersion
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) GetUpdateRepoVersion() *string {
	return s.UpdateRepoVersion
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) GetCurApackVersion() *string {
	return s.CurApackVersion
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) GetCurEsVersion() *string {
	return s.CurEsVersion
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) GetUpgradeApackVersion() *string {
	return s.UpgradeApackVersion
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) GetUpgradeEsVersion() *string {
	return s.UpgradeEsVersion
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) SetCurRepoVersion(v string) *UpgradeInfoResponseBodyResultUpgradeInfo {
	s.CurRepoVersion = &v
	return s
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) SetUpdateRepoVersion(v string) *UpgradeInfoResponseBodyResultUpgradeInfo {
	s.UpdateRepoVersion = &v
	return s
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) SetUpgrade(v bool) *UpgradeInfoResponseBodyResultUpgradeInfo {
	s.Upgrade = &v
	return s
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) SetCurApackVersion(v string) *UpgradeInfoResponseBodyResultUpgradeInfo {
	s.CurApackVersion = &v
	return s
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) SetCurEsVersion(v string) *UpgradeInfoResponseBodyResultUpgradeInfo {
	s.CurEsVersion = &v
	return s
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) SetUpgradeApackVersion(v string) *UpgradeInfoResponseBodyResultUpgradeInfo {
	s.UpgradeApackVersion = &v
	return s
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) SetUpgradeEsVersion(v string) *UpgradeInfoResponseBodyResultUpgradeInfo {
	s.UpgradeEsVersion = &v
	return s
}

func (s *UpgradeInfoResponseBodyResultUpgradeInfo) Validate() error {
	return dara.Validate(s)
}
