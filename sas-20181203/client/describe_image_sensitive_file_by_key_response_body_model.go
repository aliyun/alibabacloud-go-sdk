// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSensitiveFileByKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageSensitiveFileByKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeImageSensitiveFileByKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeImageSensitiveFileByKeyResponseBody
	GetMessage() *string
	SetPageInfo(v *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) *DescribeImageSensitiveFileByKeyResponseBody
	GetPageInfo() *DescribeImageSensitiveFileByKeyResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageSensitiveFileByKeyResponseBody
	GetRequestId() *string
	SetSensitiveFileList(v []*DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) *DescribeImageSensitiveFileByKeyResponseBody
	GetSensitiveFileList() []*DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList
	SetSuccess(v bool) *DescribeImageSensitiveFileByKeyResponseBody
	GetSuccess() *bool
}

type DescribeImageSensitiveFileByKeyResponseBody struct {
	// The status code returned. If the 200 status code is returned, the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *DescribeImageSensitiveFileByKeyResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the sensitive files.
	SensitiveFileList []*DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList `json:"SensitiveFileList,omitempty" xml:"SensitiveFileList,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeImageSensitiveFileByKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) GetPageInfo() *DescribeImageSensitiveFileByKeyResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) GetSensitiveFileList() []*DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	return s.SensitiveFileList
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) SetCode(v string) *DescribeImageSensitiveFileByKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) SetHttpStatusCode(v int32) *DescribeImageSensitiveFileByKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) SetMessage(v string) *DescribeImageSensitiveFileByKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) SetPageInfo(v *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) *DescribeImageSensitiveFileByKeyResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) SetRequestId(v string) *DescribeImageSensitiveFileByKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) SetSensitiveFileList(v []*DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) *DescribeImageSensitiveFileByKeyResponseBody {
	s.SensitiveFileList = v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) SetSuccess(v bool) *DescribeImageSensitiveFileByKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SensitiveFileList != nil {
		for _, item := range s.SensitiveFileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageSensitiveFileByKeyResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The key of the last data entry.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAGYXFWIAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM1Mzc3Njc4MzA2ODY5NmI2YTY1Nzg2NTcxNjE2NDc4NjE=
	LastRowKey *string `json:"LastRowKey,omitempty" xml:"LastRowKey,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageSensitiveFileByKeyResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileByKeyResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) GetLastRowKey() *string {
	return s.LastRowKey
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) SetCount(v int32) *DescribeImageSensitiveFileByKeyResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageSensitiveFileByKeyResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) SetLastRowKey(v string) *DescribeImageSensitiveFileByKeyResponseBodyPageInfo {
	s.LastRowKey = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageSensitiveFileByKeyResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageSensitiveFileByKeyResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList struct {
	// The suggestion.
	//
	// example:
	//
	// Assess risks based on business conditions, remove risky content, and rebuild image
	Advice *string `json:"Advice,omitempty" xml:"Advice,omitempty"`
	// The description of the sensitive file.
	//
	// example:
	//
	// Verify the validity of the leaked AK.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file path.
	//
	// example:
	//
	// /usr/lib/abc.txt
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The timestamp generated when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1663321552000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The timestamp when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1663691592000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The digest of the image.
	//
	// example:
	//
	// 0083a31cc0083a31ccf7c10367a6e783e8601e290f7c10367a6e783e860****
	LayerDigest *string `json:"LayerDigest,omitempty" xml:"LayerDigest,omitempty"`
	// The MD5 value of the sensitive file.
	//
	// example:
	//
	// b484b0dff093f358897486b58266****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The sensitive content.
	//
	// example:
	//
	// AKPIDteow289f9s************
	Promt *string `json:"Promt,omitempty" xml:"Promt,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The type of the alert for the sensitive file. Valid values:
	//
	// 	- **npm_token**: NPM access token
	//
	// 	- **ftp_cfg**: FTP configuration
	//
	// 	- **google_oauth_key**: Google OAuth key
	//
	// 	- **planetscale_passwd**: PlanetScale password
	//
	// 	- **github_ssh_key**: Github SSH key
	//
	// 	- **msbuild_publish_profile**: MSBuild publish profile
	//
	// 	- **fastly_cdn_token**: Fastly CDN token
	//
	// 	- **ssh_private_key**: SSH private key
	//
	// 	- **aws_cli**: Amazon Web Services (AWS) CLI credential
	//
	// 	- **cpanel_proftpd**: cPanel ProFTPD credential
	//
	// 	- **postgresql_passwd**: PostgreSQL password file
	//
	// 	- **discord_client_cred**: Discord client credential
	//
	// 	- **rails_database**: Rails database configuration
	//
	// 	- **aws_access_key**: AWS Access Key
	//
	// 	- **esmtp_cfg**: Extended Simple Mail Transfer Protocol (ESMTP) configuration
	//
	// 	- **docker_registry_cfg**: configuration of a Docker image repository
	//
	// 	- **pem**: Privacy-Enhanced Mail (PEM)
	//
	// 	- **common_cred**: common credential
	//
	// 	- **sftp_cfg**: configuration of connection over Secure File Transfer Protocol (SFTP)
	//
	// 	- **grafana_token**: Grafana token
	//
	// 	- **slack_token**: Slack token
	//
	// 	- **ec_private_key**: Elliptic Curve (EC) private key
	//
	// 	- **pypi_token**: Python Package Index (PyPI) token
	//
	// 	- **finicity_token**: Finicity token
	//
	// 	- **k8s_client_key**: private key for the Kubernetes client
	//
	// 	- **git_cfg**: Git configuration
	//
	// 	- **django_key**: Django key
	//
	// 	- **jenkins_ssh**: SSH configuration file for Jenkins
	//
	// 	- **openssh_private_key**: OpenSSH private key
	//
	// 	- **square_oauth**: Square OAuth credential
	//
	// 	- **typeform_token**: Typeform token
	//
	// 	- **common_database_cfg**: configuration of general database connection
	//
	// 	- **wordpress_database_cfg**: WordPress database configuration
	//
	// 	- **googlecloud_api_key**: API key for Google Cloud
	//
	// 	- **vscode_sftp**: VSCode SFTP configuration
	//
	// 	- **apache_htpasswd**: Apache htpasswd
	//
	// 	- **planetscale_token**: PlanetScale token
	//
	// 	- **contentful_preview_token**: preview token for Contentful
	//
	// 	- **php_database_cfg**: database password for a PHP application
	//
	// 	- **atom_remote_sync**: Atom remote synchronization configuration
	//
	// 	- **aws_session_token**: AWS session token
	//
	// 	- **atom_sftp_cfg**: Atom SFTP configuration
	//
	// 	- **asana_client_private_key**: Asana client key
	//
	// 	- **tencentcloud_ak**: secret ID of a third-party cloud
	//
	// 	- **rsa_private_key**: Rivest-Shamir-Adleman (RSA) private key
	//
	// 	- **github_personal_token**: personal access token for GitHub
	//
	// 	- **pgp**: Pretty Good Privacy (PGP) encrypted file
	//
	// 	- **stripe_skpk**: Stripe secret key
	//
	// 	- **square_token**: Square access token
	//
	// 	- **rails_carrierwave**: Rails Carrierwave credential
	//
	// 	- **dbeaver_database_cfg**: DBeaver database configuration
	//
	// 	- **robomongo_cred**: RoboMongo credential
	//
	// 	- **github_oauth_token**: OAuth access token for GitHub
	//
	// 	- **pulumi_token**: Pulumi token
	//
	// 	- **ventrilo_voip**: configuration of a Ventrilo VoIP server
	//
	// 	- **macos_keychain**: macOS Keychain
	//
	// 	- **amazon_mws_token**: Amazon MWS token
	//
	// 	- **dynatrace_token**: Dynatrace token
	//
	// 	- **java_keystore**: Java KeyStore (JKS)
	//
	// 	- **microsoft_sdf**: Microsoft SQL Server Compact Edition (CE) database
	//
	// 	- **kubernetes_dashboard_cred**: user credential for Kubernetes Dashboard
	//
	// 	- **atlassian_token**: Atlassian token
	//
	// 	- **rdp**: remote desktop protocol (RDP)
	//
	// 	- **mailgun_key**: Mailgun webhook signing key
	//
	// 	- **mailchimp_api_key**: API key for Mailchimp
	//
	// 	- **netrc_cfg**: netrc configuration file
	//
	// 	- **openvpn_cfg**: configuration of the OpenVPN client
	//
	// 	- **github_refresh_token**: GitHub refresh token
	//
	// 	- **salesforce**: Salesforce credential
	//
	// 	- **salesforce**: Sendinblue token
	//
	// 	- **pkcs_private_key**: PKCS#12 private key
	//
	// 	- **rubyonrails_passwd**: Ruby on Rails password file
	//
	// 	- **filezilla_ftp**: FileZilla FTP configuration
	//
	// 	- **databricks_token**: Databricks token
	//
	// 	- **gitLab_personal_toke**: personal access token for GitLab
	//
	// 	- **rails_master_key**: Rails master key
	//
	// 	- **sqlite**: SQLite3 or SQLite database
	//
	// 	- **firefox_logins**: Firefox logon configuration
	//
	// 	- **mailgun_private_token**: Mailgun private token
	//
	// 	- **joomla_cfg**: Joomla configuration
	//
	// 	- **hashicorp_terraform_token**: HashiCorp Terraform token
	//
	// 	- **jetbrains_ides**: JetBrains IDEs configuration
	//
	// 	- **heroku_api_key**: API key for Heroku
	//
	// 	- **messagebird_token**: MessageBird token
	//
	// 	- **github_app_token**: Github app token
	//
	// 	- **hashicorp_vault_token**: HashiCorp Vault token
	//
	// 	- **pgp_private_key**: PGP private key
	//
	// 	- **sshpasswd**: SSH password
	//
	// 	- **huaweicloud_ak**: secret access key of a third-party cloud
	//
	// 	- **aws_s3cmd**: AWS S3cmd configuration
	//
	// 	- **php_config**: PHP configuration
	//
	// 	- **common_private_key**: common private key
	//
	// 	- **microsoft_mdf**: Microsoft SQL Server database
	//
	// 	- **mediawiki_cfg**: MediaWiki configuration
	//
	// 	- **jenkins_cred**: Jenkins credential
	//
	// 	- **rubygems_cred**: RubyGems credential
	//
	// 	- **clojars_token**: Clojars token
	//
	// 	- **phoenix_web_passwd**: Phoenix web credential
	//
	// 	- **puttygen_private_key**: PuTTYgen private key
	//
	// 	- **google_oauth_token**: Google OAuth access token
	//
	// 	- **rubyonrails_cfg**: Ruby On Rails database configuration
	//
	// 	- **lob_api_key**: Lob API key for Lob
	//
	// 	- **pkcs_cred**: PKCS#12 certificate
	//
	// 	- **otr_private_key**: Off-the-Record Messaging (OTR) private key
	//
	// 	- **contentful_delivery_token**: Contentful delivery token
	//
	// 	- **digital_ocean_tugboat**: DigitalOcean Tugboat configuration
	//
	// 	- **dsa_private_key**: Digital Signature Algorithm (DSA) private key
	//
	// 	- **rails_app_token**: app token for Rails
	//
	// 	- **git_cred**: Git user credential
	//
	// 	- **newrelic_api_key**: User API key for New Relic
	//
	// 	- **github_hub**: hub configuration for storing GitHub tokens
	//
	// 	- **rubygem**: Rubygem Token
	//
	// example:
	//
	// google_oauth_key
	SensitiveFileKey *string `json:"SensitiveFileKey,omitempty" xml:"SensitiveFileKey,omitempty"`
	// The name of the alert type for the sensitive file.
	//
	// example:
	//
	// Google OAuth Key
	SensitiveFileName *string `json:"SensitiveFileName,omitempty" xml:"SensitiveFileName,omitempty"`
}

func (s DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetAdvice() *string {
	return s.Advice
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetFilePath() *string {
	return s.FilePath
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetLayerDigest() *string {
	return s.LayerDigest
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetMd5() *string {
	return s.Md5
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetPromt() *string {
	return s.Promt
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetSensitiveFileKey() *string {
	return s.SensitiveFileKey
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) GetSensitiveFileName() *string {
	return s.SensitiveFileName
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetAdvice(v string) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.Advice = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetDescription(v string) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.Description = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetFilePath(v string) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.FilePath = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetFirstScanTime(v int64) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.FirstScanTime = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetLastScanTime(v int64) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetLayerDigest(v string) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.LayerDigest = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetMd5(v string) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.Md5 = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetPromt(v string) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.Promt = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetRiskLevel(v string) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetSensitiveFileKey(v string) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.SensitiveFileKey = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) SetSensitiveFileName(v string) *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList {
	s.SensitiveFileName = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponseBodySensitiveFileList) Validate() error {
	return dara.Validate(s)
}
