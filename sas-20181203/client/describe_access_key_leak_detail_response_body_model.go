// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessKeyLeakDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccesskeyId(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetAccesskeyId() *string
	SetAsset(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetAsset() *string
	SetCode(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetCode() *string
	SetDealTime(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetDealTime() *string
	SetDealType(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetDealType() *string
	SetGithubFileName(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGithubFileName() *string
	SetGithubFileType(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGithubFileType() *string
	SetGithubFileUpdateTime(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGithubFileUpdateTime() *string
	SetGithubFileUrl(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGithubFileUrl() *string
	SetGithubRepoName(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGithubRepoName() *string
	SetGithubRepoUrl(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGithubRepoUrl() *string
	SetGithubUser(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGithubUser() *string
	SetGithubUserPicUrl(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGithubUserPicUrl() *string
	SetGmtCreate(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetGmtModified() *string
	SetRemark(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetRequestId() *string
	SetSource(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetSource() *string
	SetTokenValid(v int32) *DescribeAccessKeyLeakDetailResponseBody
	GetTokenValid() *int32
	SetType(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetType() *string
	SetWhitelistStatus(v string) *DescribeAccessKeyLeakDetailResponseBody
	GetWhitelistStatus() *string
	SetWhitelistTime(v int64) *DescribeAccessKeyLeakDetailResponseBody
	GetWhitelistTime() *int64
}

type DescribeAccessKeyLeakDetailResponseBody struct {
	// The ID of the AccessKey pair that is leaked.
	//
	// example:
	//
	// yourAccessKeyID
	AccesskeyId *string `json:"AccesskeyId,omitempty" xml:"AccesskeyId,omitempty"`
	// The platform to which the asset belongs. The value is fixed as **Cloud platform**.
	//
	// example:
	//
	// Cloud platform
	Asset *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	// The code snippet that is leaked.
	//
	// example:
	//
	// \\n1231 \\nak=yourAccessKeyID \\n12311123 \\nsk1999 \\nsk1999sk1999 \\nsk1999sk1999 \\n\\n\\ntest001 ak hht \\nak=yourAccessKeyID \\nsk=yourAccessKeySecret
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the AccessKey pair leak was handled.
	//
	// example:
	//
	// 2022-01-17 15:47:08
	DealTime *string `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// The solution to the AccessKey pair leak. Valid values:
	//
	// 	- **manual**: manually deleted
	//
	// 	- **disable**: manually disabled
	//
	// 	- **add-whitelist**: added to the whitelist
	//
	// 	- **pending**: unhandled
	//
	// example:
	//
	// add-whitelist
	DealType *string `json:"DealType,omitempty" xml:"DealType,omitempty"`
	// The name of the GitHub file.
	//
	// example:
	//
	// testAkLeak
	GithubFileName *string `json:"GithubFileName,omitempty" xml:"GithubFileName,omitempty"`
	// The type of the GitHub file. Valid values:
	//
	// 	- Python
	//
	// 	- XML
	//
	// 	- GO
	//
	// 	- Javascript
	//
	// 	- INI
	//
	// 	- JSON
	//
	// 	- C++
	//
	// example:
	//
	// Python
	GithubFileType *string `json:"GithubFileType,omitempty" xml:"GithubFileType,omitempty"`
	// The time when the GitHub file was updated.
	//
	// example:
	//
	// 2021-07-06T09:49:33
	GithubFileUpdateTime *string `json:"GithubFileUpdateTime,omitempty" xml:"GithubFileUpdateTime,omitempty"`
	// The URL of the GitHub file.
	//
	// example:
	//
	// https://github.com/Blue00Blue/ExamOnline/blob/6c932c10fc3f217783f3937e2b230f79656c18a7/testAk****
	GithubFileUrl *string `json:"GithubFileUrl,omitempty" xml:"GithubFileUrl,omitempty"`
	// The name of the GitHub repository.
	//
	// example:
	//
	// ExamOnline
	GithubRepoName *string `json:"GithubRepoName,omitempty" xml:"GithubRepoName,omitempty"`
	// The URL of the GitHub repository.
	//
	// example:
	//
	// https://github.com/Blue00Blue/ExamOn****
	GithubRepoUrl *string `json:"GithubRepoUrl,omitempty" xml:"GithubRepoUrl,omitempty"`
	// The username of the GitHub user.
	//
	// example:
	//
	// Blue00Blue
	GithubUser *string `json:"GithubUser,omitempty" xml:"GithubUser,omitempty"`
	// The URL of the profile picture for the GitHub user.
	//
	// example:
	//
	// https://avatars.githubusercontent.com/u/26296896?s=48&v=****
	GithubUserPicUrl *string `json:"GithubUserPicUrl,omitempty" xml:"GithubUserPicUrl,omitempty"`
	// The first time when the AccessKey pair leak was detected.
	//
	// example:
	//
	// 2021-07-06 17:49:41
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last time when the AccessKey pair leak was detected.
	//
	// example:
	//
	// 2021-07-06 17:49:39
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The remarks of the AccessKey pair leak.
	//
	// example:
	//
	// 12
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 79CFF74D-E967-5407-8A78-EE03B925FDAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The platform on which the AccessKey pair leak is detected.
	//
	// example:
	//
	// GitHub
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The validity of the key that is associated with the AccessKey pair. Valid values:
	//
	// 	- **0**: to be confirmed.
	//
	// 	- **1**: valid.
	//
	// 	- **2**: invalid.
	//
	// example:
	//
	// 2
	TokenValid *int32 `json:"TokenValid,omitempty" xml:"TokenValid,omitempty"`
	// The type of the leak. The value is fixed as **AccessKey**.
	//
	// example:
	//
	// AccessKey
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Indicates whether the AccessKey pair leak is added to the whitelist. Valid values:
	//
	// 	- **no**: no
	//
	// 	- **yes**: yes
	//
	// example:
	//
	// no
	WhitelistStatus *string `json:"WhitelistStatus,omitempty" xml:"WhitelistStatus,omitempty"`
	// The time when the AccessKey pair was added to the whitelist. Unit: milliseconds.
	//
	// example:
	//
	// 1689172004478
	WhitelistTime *int64 `json:"WhitelistTime,omitempty" xml:"WhitelistTime,omitempty"`
}

func (s DescribeAccessKeyLeakDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessKeyLeakDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetAccesskeyId() *string {
	return s.AccesskeyId
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetAsset() *string {
	return s.Asset
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetDealTime() *string {
	return s.DealTime
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetDealType() *string {
	return s.DealType
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGithubFileName() *string {
	return s.GithubFileName
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGithubFileType() *string {
	return s.GithubFileType
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGithubFileUpdateTime() *string {
	return s.GithubFileUpdateTime
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGithubFileUrl() *string {
	return s.GithubFileUrl
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGithubRepoName() *string {
	return s.GithubRepoName
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGithubRepoUrl() *string {
	return s.GithubRepoUrl
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGithubUser() *string {
	return s.GithubUser
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGithubUserPicUrl() *string {
	return s.GithubUserPicUrl
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetSource() *string {
	return s.Source
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetTokenValid() *int32 {
	return s.TokenValid
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetWhitelistStatus() *string {
	return s.WhitelistStatus
}

func (s *DescribeAccessKeyLeakDetailResponseBody) GetWhitelistTime() *int64 {
	return s.WhitelistTime
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetAccesskeyId(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.AccesskeyId = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetAsset(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.Asset = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetCode(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetDealTime(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.DealTime = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetDealType(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.DealType = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGithubFileName(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GithubFileName = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGithubFileType(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GithubFileType = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGithubFileUpdateTime(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GithubFileUpdateTime = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGithubFileUrl(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GithubFileUrl = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGithubRepoName(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GithubRepoName = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGithubRepoUrl(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GithubRepoUrl = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGithubUser(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GithubUser = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGithubUserPicUrl(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GithubUserPicUrl = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGmtCreate(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetGmtModified(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetRemark(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetRequestId(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetSource(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.Source = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetTokenValid(v int32) *DescribeAccessKeyLeakDetailResponseBody {
	s.TokenValid = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetType(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetWhitelistStatus(v string) *DescribeAccessKeyLeakDetailResponseBody {
	s.WhitelistStatus = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) SetWhitelistTime(v int64) *DescribeAccessKeyLeakDetailResponseBody {
	s.WhitelistTime = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
