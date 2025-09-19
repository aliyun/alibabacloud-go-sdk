// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSensitiveFileListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageSensitiveFileListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeImageSensitiveFileListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeImageSensitiveFileListResponseBody
	GetMessage() *string
	SetPageInfo(v *DescribeImageSensitiveFileListResponseBodyPageInfo) *DescribeImageSensitiveFileListResponseBody
	GetPageInfo() *DescribeImageSensitiveFileListResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageSensitiveFileListResponseBody
	GetRequestId() *string
	SetSensitiveFileList(v []*DescribeImageSensitiveFileListResponseBodySensitiveFileList) *DescribeImageSensitiveFileListResponseBody
	GetSensitiveFileList() []*DescribeImageSensitiveFileListResponseBodySensitiveFileList
	SetSuccess(v bool) *DescribeImageSensitiveFileListResponseBody
	GetSuccess() *bool
}

type DescribeImageSensitiveFileListResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
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
	PageInfo *DescribeImageSensitiveFileListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8D19A089-E6BC-5244-800C-7E590D50487F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the information about the sensitive files.
	SensitiveFileList []*DescribeImageSensitiveFileListResponseBodySensitiveFileList `json:"SensitiveFileList,omitempty" xml:"SensitiveFileList,omitempty" type:"Repeated"`
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

func (s DescribeImageSensitiveFileListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageSensitiveFileListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeImageSensitiveFileListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeImageSensitiveFileListResponseBody) GetPageInfo() *DescribeImageSensitiveFileListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageSensitiveFileListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageSensitiveFileListResponseBody) GetSensitiveFileList() []*DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	return s.SensitiveFileList
}

func (s *DescribeImageSensitiveFileListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImageSensitiveFileListResponseBody) SetCode(v string) *DescribeImageSensitiveFileListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBody) SetHttpStatusCode(v int32) *DescribeImageSensitiveFileListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBody) SetMessage(v string) *DescribeImageSensitiveFileListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBody) SetPageInfo(v *DescribeImageSensitiveFileListResponseBodyPageInfo) *DescribeImageSensitiveFileListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBody) SetRequestId(v string) *DescribeImageSensitiveFileListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBody) SetSensitiveFileList(v []*DescribeImageSensitiveFileListResponseBodySensitiveFileList) *DescribeImageSensitiveFileListResponseBody {
	s.SensitiveFileList = v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBody) SetSuccess(v bool) *DescribeImageSensitiveFileListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageSensitiveFileListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
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
	// The number of entries returned per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageSensitiveFileListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) GetLastRowKey() *string {
	return s.LastRowKey
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) SetCount(v int32) *DescribeImageSensitiveFileListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageSensitiveFileListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) SetLastRowKey(v string) *DescribeImageSensitiveFileListResponseBodyPageInfo {
	s.LastRowKey = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageSensitiveFileListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageSensitiveFileListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeImageSensitiveFileListResponseBodySensitiveFileList struct {
	// The suggestion.
	//
	// example:
	//
	// Assess risks based on business conditions, remove risky content.
	Advice *string `json:"Advice,omitempty" xml:"Advice,omitempty"`
	// The key of the sensitive file type.
	//
	// example:
	//
	// password
	ClassKey *string `json:"ClassKey,omitempty" xml:"ClassKey,omitempty"`
	// The name of the sensitive file type.
	//
	// example:
	//
	// password
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The number of scans that are performed on the sensitive file.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The description of the sensitive file.
	//
	// example:
	//
	// Verify the validity of the leaked AK.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp generated when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1663321552000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The timestamp generated when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1663321552000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
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
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The type of the alert for the sensitive file. Valid values:
	//
	// 	- **npm_token**: Node Package Manager (NPM) access token.
	//
	// 	- **ftp_cfg**: FTP configuration.
	//
	// 	- **google_oauth_key**: Google OAuth key.
	//
	// 	- **planetscale_passwd**: PlanetScale password.
	//
	// 	- **github_ssh_key**: GitHub SSH key.
	//
	// 	- **msbuild_publish_profile**: MSBuild publish profile.
	//
	// 	- **fastly_cdn_token**: Fastly CDN token.
	//
	// 	- **ssh_private_key**: SSH private key.
	//
	// 	- **aws_cli**: Amazon Web Services (AWS) CLI credential.
	//
	// 	- **cpanel_proftpd**: cPanel ProFTPD credential.
	//
	// 	- **postgresql_passwd**: PostgreSQL password file.
	//
	// 	- **discord_client_cred**: Discord client credential.
	//
	// 	- **rails_database**: Rails database configuration.
	//
	// 	- **aws_access_key**: AWS access key.
	//
	// 	- **esmtp_cfg**: Extended Simple Mail Transfer Protocol (ESMTP) configuration.
	//
	// 	- **docker_registry_cfg**: Docker image repository configuration.
	//
	// 	- **pem**: Privacy-Enhanced Mail (PEM).
	//
	// 	- **common_cred**: common credential.
	//
	// 	- **sftp_cfg**: Secure File Transfer Protocol (SFTP) connection configuration.
	//
	// 	- **grafana_token**: Grafana token.
	//
	// 	- **slack_token**: Slack token.
	//
	// 	- **ec_private_key**: EC private key.
	//
	// 	- **pypi_token**: upload token for the Python Package Index (PyPI).
	//
	// 	- **finicity_token**: Finicity token.
	//
	// 	- **k8s_client_key**: Kubernetes client private key.
	//
	// 	- **git_cfg**: Git configuration.
	//
	// 	- **django_key**: Django key.
	//
	// 	- **jenkins_ssh**: Jenkins SSH configuration file.
	//
	// 	- **openssh_private_key**: OpenSSH private key.
	//
	// 	- **square_oauth**: OAuth credential for Square.
	//
	// 	- **typeform_token**: Typeform token.
	//
	// 	- **common_database_cfg**: general database connection configuration.
	//
	// 	- **wordpress_database_cfg**: WordPress database configuration.
	//
	// 	- **googlecloud_api_key**: API key for Google Cloud.
	//
	// 	- **vscode_sftp**: VS Code SFTP configuration.
	//
	// 	- **apache_htpasswd**: Apache htpasswd.
	//
	// 	- **planetscale_token**: PlanetScale token.
	//
	// 	- **contentful_preview_token**: preview token for Contentful.
	//
	// 	- **php_database_cfg**: database password for a PHP application.
	//
	// 	- **atom_remote_sync**: Atom remote synchronization configuration.
	//
	// 	- **aws_session_token**: AWS session token.
	//
	// 	- **atom_sftp_cfg**: Atom SFTP configuration.
	//
	// 	- **asana_client_private_key**: Asana client key.
	//
	// 	- **tencentcloud_ak**: secret ID of a third-party cloud.
	//
	// 	- **rsa_private_key**: Rivest-Shamir-Adleman (RSA) private key.
	//
	// 	- **github_personal_token**: personal access token for GitHub.
	//
	// 	- **pgp**: Pretty Good Privacy (PGP) encrypted file.
	//
	// 	- **stripe_skpk**: Stripe secret key.
	//
	// 	- **square_token**: Square access token.
	//
	// 	- **rails_carrierwave**: file upload credential for Rails Carrierwave.
	//
	// 	- **dbeaver_database_cfg**: DBeaver database configuration.
	//
	// 	- **robomongo_cred**: Robomongo credential.
	//
	// 	- **github_oauth_token**: OAuth access token for GitHub.
	//
	// 	- **pulumi_token**: Pulumi token.
	//
	// 	- **ventrilo_voip**: Ventrilo VoIP server configuration.
	//
	// 	- **macos_keychain**: macOS keychain.
	//
	// 	- **amazon_mws_token**: Amazon MWS token.
	//
	// 	- **dynatrace_token**: Dynatrace token.
	//
	// 	- **java_keystore**: Java KeyStore (JKS).
	//
	// 	- **microsoft_sdf**: Microsoft SQL Server Compact Edition (CE) database.
	//
	// 	- **kubernetes_dashboard_cred**: user credential for Kubernetes Dashboard.
	//
	// 	- **atlassian_token**: Atlassian token.
	//
	// 	- **rdp**: remote desktop protocol (RDP).
	//
	// 	- **mailgun_key**: Mailgun webhook signing key.
	//
	// 	- **mailchimp_api_key**: API key for Mailchimp.
	//
	// 	- **netrc_cfg**: .netrc configuration file.
	//
	// 	- **openvpn_cfg**: OpenVPN client configuration.
	//
	// 	- **github_refresh_token**: GitHub refresh token.
	//
	// 	- **salesforce**: Salesforce credential.
	//
	// 	- **sendinblue**: Sendinblue token.
	//
	// 	- **pkcs_private_key**: PKCS#12 key.
	//
	// 	- **rubyonrails_passwd**: Ruby on Rails password file.
	//
	// 	- **filezilla_ftp**: FileZilla FTP configuration.
	//
	// 	- **databricks_token**: Databricks token.
	//
	// 	- **gitLab_personal_token**: personal access token for GitLab.
	//
	// 	- **rails_master_key**: Rails master key.
	//
	// 	- **sqlite**: SQLite3 or SQLite database.
	//
	// 	- **firefox_logins**: Firefox logon configuration.
	//
	// 	- **mailgun_private_token**: Mailgun private token.
	//
	// 	- **joomla_cfg**: Joomla configuration.
	//
	// 	- **hashicorp_terraform_token**: HashiCorp Terraform token.
	//
	// 	- **jetbrains_ides**: JetBrains IDEs configuration.
	//
	// 	- **heroku_api_key**: Heroku API key.
	//
	// 	- **messagebird_token**: MessageBird token.
	//
	// 	- **github_app_token**: GitHub app token.
	//
	// 	- **hashicorp_vault_token**: HashiCorp Vault token.
	//
	// 	- **pgp_private_key**: PGP private key.
	//
	// 	- **sshpasswd**: SSH password.
	//
	// 	- **huaweicloud_ak**: secret access key of a third-party cloud.
	//
	// 	- **aws_s3cmd**: AWS S3cmd configuration.
	//
	// 	- **php_config**: PHP configuration.
	//
	// 	- **common_private_key**: private key of a common type.
	//
	// 	- **microsoft_mdf**: Microsoft SQL Server database.
	//
	// 	- **mediawiki_cfg**: MediaWiki configuration.
	//
	// 	- **jenkins_cred**: Jenkins credential.
	//
	// 	- **rubygems_cred**: RubyGems credential.
	//
	// 	- **clojars_token**: Clojars token.
	//
	// 	- **phoenix_web_passwd**: Phoenix web credential.
	//
	// 	- **puttygen_private_key**: PuTTYgen private key.
	//
	// 	- **google_oauth_token**: Google OAuth access token.
	//
	// 	- **rubyonrails_cfg**: Ruby on Rails database configuration.
	//
	// 	- **lob_api_key**: Lob API key.
	//
	// 	- **pkcs_cred**: PKCS#12 certificate.
	//
	// 	- **otr_private_key**: Off-the-Record Messaging (OTR) private key.
	//
	// 	- **contentful_delivery_token**: delivery token for Contentful.
	//
	// 	- **digital_ocean_tugboat**: DigitalOcean Tugboat configuration.
	//
	// 	- **dsa_private_key**: Digital Signature Algorithm (DSA) private key.
	//
	// 	- **rails_app_token**: Rails app token.
	//
	// 	- **git_cred**: Git user credential.
	//
	// 	- **newrelic_api_key**: user API key for New Relic.
	//
	// 	- **github_hub**: hub configuration for storing GitHub tokens.
	//
	// 	- **rubygem**: RubyGems token.
	//
	// example:
	//
	// google_oauth_key
	SensitiveFileKey *string `json:"SensitiveFileKey,omitempty" xml:"SensitiveFileKey,omitempty"`
	// The name of the alert type for the sensitive file.
	//
	// example:
	//
	// AccessKeyLeak
	SensitiveFileName *string `json:"SensitiveFileName,omitempty" xml:"SensitiveFileName,omitempty"`
	// The status of the sensitive file. Valid values:
	//
	// 	- **0**: unhandled.
	//
	// 	- **1**: handled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unprocessed mirrors.
	//
	// example:
	//
	// 2
	UnprocessedNum *int32 `json:"UnprocessedNum,omitempty" xml:"UnprocessedNum,omitempty"`
}

func (s DescribeImageSensitiveFileListResponseBodySensitiveFileList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileListResponseBodySensitiveFileList) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetAdvice() *string {
	return s.Advice
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetClassKey() *string {
	return s.ClassKey
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetClassName() *string {
	return s.ClassName
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetSensitiveFileKey() *string {
	return s.SensitiveFileKey
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetSensitiveFileName() *string {
	return s.SensitiveFileName
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) GetUnprocessedNum() *int32 {
	return s.UnprocessedNum
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetAdvice(v string) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.Advice = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetClassKey(v string) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.ClassKey = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetClassName(v string) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.ClassName = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetCount(v int32) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.Count = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetDescription(v string) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.Description = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetFirstScanTime(v int64) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.FirstScanTime = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetLastScanTime(v int64) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetRiskLevel(v string) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetSensitiveFileKey(v string) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.SensitiveFileKey = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetSensitiveFileName(v string) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.SensitiveFileName = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetStatus(v int32) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.Status = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) SetUnprocessedNum(v int32) *DescribeImageSensitiveFileListResponseBodySensitiveFileList {
	s.UnprocessedNum = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponseBodySensitiveFileList) Validate() error {
	return dara.Validate(s)
}
