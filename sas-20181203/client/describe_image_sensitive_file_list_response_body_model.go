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
	// The result code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
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
	// The detailed information about the error code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *DescribeImageSensitiveFileListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates a unique ID for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 8D19A089-E6BC-5244-800C-7E590D50487F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of sensitive files.
	SensitiveFileList []*DescribeImageSensitiveFileListResponseBodySensitiveFileList `json:"SensitiveFileList,omitempty" xml:"SensitiveFileList,omitempty" type:"Repeated"`
	// Indicates whether the query was successful. Valid values:
	//
	// - **true**: successful
	//
	// - **false**: failed.
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

type DescribeImageSensitiveFileListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The key of the last entry.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAGYXFWIAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM1Mzc3Njc4MzA2ODY5NmI2YTY1Nzg2NTcxNjE2N******
	LastRowKey *string `json:"LastRowKey,omitempty" xml:"LastRowKey,omitempty"`
	// The number of entries returned per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
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
	// The hardening suggestion for the sensitive file check item.
	//
	// example:
	//
	// PEM (Privacy Enhanced Mail) format is a common format for digital certificates. PEM files can contain certificates, public keys, private keys, and other sensitive information. When a PEM file is either unencrypted or protected with a weak password, or if the password has been compromise, it poses a significantly higher security risk. This detection rule aims to identify such PEM files.
	Advice *string `json:"Advice,omitempty" xml:"Advice,omitempty"`
	// The classification key of the sensitive file.
	//
	// example:
	//
	// password
	ClassKey *string `json:"ClassKey,omitempty" xml:"ClassKey,omitempty"`
	// The classification name of the sensitive file.
	//
	// example:
	//
	// password
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The number of times the sensitive file was detected by scans.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The description of the sensitive file check item.
	//
	// example:
	//
	// Assess the risk based on business context and promptly remove any risky content.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp of the first scan. Unit: milliseconds.
	//
	// example:
	//
	// 1663321552000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The timestamp of the most recent scan. Unit: milliseconds.
	//
	// example:
	//
	// 1663321552000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The risk level. Valid values:
	//
	// - **high**: high
	//
	// - **medium**: medium
	//
	// - **low**: low.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The sensitive file alerting type. Valid values:
	//
	// - **npm_token**: NPM access token
	//
	// - **ftp_cfg**: FTP configuration
	//
	// - **google_oauth_key**: Google OAuth Key
	//
	// - **planetscale_passwd**: Planetscale password
	//
	// - **github_ssh_key**: Github SSH key
	//
	// - **msbuild_publish_profile**: MSBuild publish profile
	//
	// - **fastly_cdn_token**: Fastly CDN token
	//
	// - **ssh_private_key**: SSH private key
	//
	// - **aws_cli**: AWS CLI credentials
	//
	// - **cpanel_proftpd**: cPanel ProFTPd credentials
	//
	// - **postgresql_passwd**: PostgreSQL password file
	//
	// - **discord_client_cred**: Discord client credentials
	//
	// - **rails_database**: Rails database configuration
	//
	// - **aws_access_key**: AWS Access Key
	//
	// - **esmtp_cfg**: ESMTP mail server configuration
	//
	// - **docker_registry_cfg**: Docker image repository configuration
	//
	// - **pem**: PEM
	//
	// - **common_cred**: common credentials
	//
	// - **sftp_cfg**: SFTP connection configuration
	//
	// - **grafana_token**: Grafana token
	//
	// - **slack_token**: Slack Token
	//
	// - **ec_private_key**: EC private key
	//
	// - **pypi_token**: PyPI upload token
	//
	// - **finicity_token**: Finicity platform token
	//
	// - **k8s_client_key**: Kubernetes client private key
	//
	// - **git_cfg**: Git configuration
	//
	// - **django_key**: Django key
	//
	// - **jenkins_ssh**: Jenkins SSH configuration file
	//
	// - **openssh_private_key**: OPENSSH private key
	//
	// - **square_oauth**: Square OAuth credentials
	//
	// - **typeform_token**: Typeform token
	//
	// - **common_database_cfg**: common database connection configuration
	//
	// - **wordpress_database_cfg**: WordPress database configuration
	//
	// - **googlecloud_api_key**: Google Cloud API Key
	//
	// - **vscode_sftp**: VSCode SFTP configuration
	//
	// - **apache_htpasswd**: Apache htpasswd
	//
	// - **planetscale_token**: Planetscale token
	//
	// - **contentful_preview_token**: Contentful Preview token
	//
	// - **php_database_cfg**: PHP application database password
	//
	// - **atom_remote_sync**: Atom remote synchronization configuration
	//
	// - **aws_session_token**: AWS session token
	//
	// - **atom_sftp_cfg**: Atom SFTP configuration
	//
	// - **asana_client_private_key**: Asana client private key
	//
	// - **tencentcloud_ak**: third-party cloud SecretId
	//
	// - **rsa_private_key**: RSA private key
	//
	// - **github_personal_token**: Github Personal access token
	//
	// - **pgp**: PGP encrypt file
	//
	// - **stripe_skpk**: Stripe Secret Key
	//
	// - **square_token**: Square access token
	//
	// - **rails_carrierwave**: Rails Carrierwave file upload credentials
	//
	// - **dbeaver_database_cfg**: DBeaver database configuration
	//
	// - **robomongo_cred**: Robomongo credentials
	//
	// - **github_oauth_token**: Github OAuth access token
	//
	// - **pulumi_token**: Pulumi token
	//
	// - **ventrilo_voip**: Ventrilo VoIP Server configuration
	//
	// - **macos_keychain**: macOS Keychain
	//
	// - **amazon_mws_token**: Amazon MWS Token
	//
	// - **dynatrace_token**: Dynatrace token
	//
	// - **java_keystore**: Java KeyStore
	//
	// - **microsoft_sdf**: Microsoft SQL CE database
	//
	// - **kubernetes_dashboard_cred**: Kubernetes Dashboard user credentials
	//
	// - **atlassian_token**: Atlassian token
	//
	// - **rdp**: Remote Desktop Protocol (RDP) connection
	//
	// - **mailgun_key**: Mailgun Webhook Signing Key
	//
	// - **mailchimp_api_key**: Mailchimp API Key
	//
	// - **netrc_cfg**: .netrc configuration file
	//
	// - **openvpn_cfg**: OpenVPN client configuration
	//
	// - **github_refresh_token**: Github Refresh Token
	//
	// - **salesforce**: Salesforce credentials
	//
	// - **sendinblue**: Sendinblue token
	//
	// - **pkcs_private_key**: PKCS#12 key
	//
	// - **rubyonrails_passwd**: Ruby on Rails password file
	//
	// - **filezilla_ftp**: FileZilla FTP configuration
	//
	// - **databricks_token**: Databricks token
	//
	// - **gitLab_personal_token**: GitLab Personal access token
	//
	// - **rails_master_key**: Rails Master Key
	//
	// - **sqlite**: SQLite3/SQLite database
	//
	// - **firefox_logins**: Firefox logon configuration
	//
	// - **mailgun_private_token**: Mailgun Private token
	//
	// - **joomla_cfg**: Joomla configuration
	//
	// - **hashicorp_terraform_token**: Hashicorp Terraform Token
	//
	// - **jetbrains_ides**: Jetbrains IDEs configuration
	//
	// - **heroku_api_key**: Heroku API key
	//
	// - **messagebird_token**: MessageBird token
	//
	// - **github_app_token**: Github App Token
	//
	// - **hashicorp_vault_token**: Hashicorp Vault Token
	//
	// - **pgp_private_key**: PGP private key
	//
	// - **sshpasswd**: SSH password
	//
	// - **huaweicloud_ak**: third-party cloud Secret Access Key
	//
	// - **aws_s3cmd**: AWS S3cmd configuration
	//
	// - **php_config**: PHP configuration
	//
	// - **common_private_key**: common private key types
	//
	// - **microsoft_mdf**: Microsoft SQL database
	//
	// - **mediawiki_cfg**: MediaWiki configuration
	//
	// - **jenkins_cred**: Jenkins credentials
	//
	// - **rubygems_cred**: Rubygems credentials
	//
	// - **clojars_token**: Clojars token
	//
	// - **phoenix_web_passwd**: Phoenix Web credentials
	//
	// - **puttygen_private_key**: PuTTYgen private key
	//
	// - **google_oauth_token**: Google OAuth access token
	//
	// - **rubyonrails_cfg**: Ruby On Rails database configuration
	//
	// - **lob_api_key**: Lob API Key
	//
	// - **pkcs_cred**: PKCS#12 certificate
	//
	// - **otr_private_key**: OTR private key
	//
	// - **contentful_delivery_token**: Contentful Delivery token
	//
	// - **digital_ocean_tugboat**: Digital Ocean Tugboat configuration
	//
	// - **dsa_private_key**: DSA private key
	//
	// - **rails_app_token**: Rails App token
	//
	// - **git_cred**: Git user credentials
	//
	// - **newrelic_api_key**: New Relic User API Key
	//
	// - **github_hub**: hub configuration that stores Github tokens
	//
	// - **rubygem**: Rubygem token.
	//
	// example:
	//
	// google_oauth_key
	SensitiveFileKey *string `json:"SensitiveFileKey,omitempty" xml:"SensitiveFileKey,omitempty"`
	// The name of the sensitive file alerting type.
	//
	// example:
	//
	// AccessKeyLeak
	SensitiveFileName *string `json:"SensitiveFileName,omitempty" xml:"SensitiveFileName,omitempty"`
	// The status of the sensitive file check item. Valid values:
	//
	// - **0**: Unprocessed.
	//
	// - **1**: Processed.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unprocessed images.
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
