// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSensitiveFileListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeImageSensitiveFileListRequest
	GetCriteria() *string
	SetCriteriaType(v string) *DescribeImageSensitiveFileListRequest
	GetCriteriaType() *string
	SetCurrentPage(v int32) *DescribeImageSensitiveFileListRequest
	GetCurrentPage() *int32
	SetImageUuid(v string) *DescribeImageSensitiveFileListRequest
	GetImageUuid() *string
	SetLang(v string) *DescribeImageSensitiveFileListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageSensitiveFileListRequest
	GetPageSize() *int32
	SetRiskLevel(v string) *DescribeImageSensitiveFileListRequest
	GetRiskLevel() *string
	SetScanRange(v []*string) *DescribeImageSensitiveFileListRequest
	GetScanRange() []*string
	SetSensitiveKeyList(v []*string) *DescribeImageSensitiveFileListRequest
	GetSensitiveKeyList() []*string
}

type DescribeImageSensitiveFileListRequest struct {
	// The value that corresponds to the search type.
	//
	// example:
	//
	// Rails Master Key
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The search type for sensitive files. Valid values:
	//
	// - **SensitiveFileKey**: the sensitive file alerting categorization. Valid values:
	//
	//     - **npm_token**: NPM access token
	//
	//     - **ftp_cfg**: FTP configuration
	//
	//     - **google_oauth_key**: Google OAuth Key
	//
	//     - **planetscale_passwd**: Planetscale password
	//
	//     - **github_ssh_key**: Github SSH key
	//
	//     - **msbuild_publish_profile**: MSBuild publish profile
	//
	//     - **fastly_cdn_token**: Fastly CDN token
	//
	//     - **ssh_private_key**: SSH private key
	//
	//     - **aws_cli**: AWS CLI credentials
	//
	//     - **cpanel_proftpd**: cPanel ProFTPd credentials
	//
	//     - **postgresql_passwd**: PostgreSQL password file
	//
	//     - **discord_client_cred**: Discord client credentials
	//
	//     - **rails_database**: Rails database configuration
	//
	//     - **aws_access_key**: AWS Access Key
	//
	//     - **esmtp_cfg**: ESMTP mail server configuration
	//
	//     - **docker_registry_cfg**: Docker image repository configuration
	//
	//     - **pem**: PEM
	//
	//     - **common_cred**: common credentials
	//
	//     - **sftp_cfg**: SFTP connection configuration
	//
	//     - **grafana_token**: Grafana token
	//
	//     - **slack_token**: Slack Token
	//
	//     - **ec_private_key**: EC private key
	//
	//     - **pypi_token**: PyPI upload token
	//
	//     - **finicity_token**: Finicity platform token
	//
	//     - **k8s_client_key**: Kubernetes client private key
	//
	//     - **git_cfg**: Git configuration
	//
	//     - **django_key**: Django key
	//
	//     - **jenkins_ssh**: Jenkins SSH configuration file
	//
	//     - **openssh_private_key**: OPENSSH private key
	//
	//     - **square_oauth**: Square OAuth credentials
	//
	//     - **typeform_token**: Typeform token
	//
	//     - **common_database_cfg**: common database connection configuration
	//
	//     - **wordpress_database_cfg**: WordPress database configuration
	//
	//     - **googlecloud_api_key**: Google Cloud API Key
	//
	//     - **vscode_sftp**: VSCode SFTP configuration
	//
	//     - **apache_htpasswd**: Apache htpasswd
	//
	//     - **planetscale_token**: Planetscale token
	//
	//     - **contentful_preview_token**: Contentful Preview token
	//
	//     - **php_database_cfg**: PHP application database password
	//
	//     - **atom_remote_sync**: Atom remote synchronization configuration
	//
	//     - **aws_session_token**: AWS session token
	//
	//     - **atom_sftp_cfg**: Atom SFTP configuration
	//
	//     - **asana_client_private_key**: Asana client private key
	//
	//     - **tencentcloud_ak**: third-party cloud SecretId
	//
	//     - **rsa_private_key**: RSA private key
	//
	//     - **github_personal_token**: Github Personal access token
	//
	//     - **pgp**: PGP encrypt file
	//
	//     - **stripe_skpk**: Stripe Secret Key
	//
	//     - **square_token**: Square access token
	//
	//     - **rails_carrierwave**: Rails Carrierwave file upload credentials
	//
	//     - **dbeaver_database_cfg**: DBeaver database configuration
	//
	//     - **robomongo_cred**: Robomongo credentials
	//
	//     - **github_oauth_token**: Github OAuth access token
	//
	//     - **pulumi_token**: Pulumi token
	//
	//     - **ventrilo_voip**: Ventrilo VoIP Server configuration
	//
	//     - **macos_keychain**: macOS Keychain
	//
	//     - **amazon_mws_token**: Amazon MWS Token
	//
	//     - **dynatrace_token**: Dynatrace token
	//
	//     - **java_keystore**: Java KeyStore
	//
	//     - **microsoft_sdf**: Microsoft SQL CE database
	//
	//     - **kubernetes_dashboard_cred**: Kubernetes Dashboard user credentials
	//
	//     - **atlassian_token**: Atlassian token
	//
	//     - **rdp**: Remote Desktop Protocol (RDP) connection
	//
	//     - **mailgun_key**: Mailgun Webhook Signing Key
	//
	//     - **mailchimp_api_key**: Mailchimp API Key
	//
	//     - **netrc_cfg**: .netrc configuration file
	//
	//     - **openvpn_cfg**: OpenVPN client configuration
	//
	//     - **github_refresh_token**: Github Refresh Token
	//
	//     - **salesforce**: Salesforce credentials
	//
	//     - **sendinblue**: Sendinblue token
	//
	//     - **pkcs_private_key**: PKCS#12 key
	//
	//     - **rubyonrails_passwd**: Ruby on Rails password file
	//
	//     - **filezilla_ftp**: FileZilla FTP configuration
	//
	//     - **databricks_token**: Databricks token
	//
	//     - **gitLab_personal_token**: GitLab Personal access token
	//
	//     - **rails_master_key**: Rails Master Key
	//
	//     - **sqlite**: SQLite3/SQLite database
	//
	//     - **firefox_logins**: Firefox logon configuration
	//
	//     - **mailgun_private_token**: Mailgun Private token
	//
	//     - **joomla_cfg**: Joomla configuration
	//
	//     - **hashicorp_terraform_token**: Hashicorp Terraform Token
	//
	//     - **jetbrains_ides**: Jetbrains IDEs configuration
	//
	//     - **heroku_api_key**: Heroku API key
	//
	//     - **messagebird_token**: MessageBird token
	//
	//     - **github_app_token**: Github App Token
	//
	//     - **hashicorp_vault_token**: Hashicorp Vault Token
	//
	//     - **pgp_private_key**: PGP private key
	//
	//     - **sshpasswd**: SSH password
	//
	//     - **huaweicloud_ak**: third-party cloud Secret Access Key
	//
	//     - **aws_s3cmd**: AWS S3cmd configuration
	//
	//     - **php_config**: PHP configuration
	//
	//     - **common_private_key**: common private key types
	//
	//     - **microsoft_mdf**: Microsoft SQL database
	//
	//     - **mediawiki_cfg**: MediaWiki configuration
	//
	//     - **jenkins_cred**: Jenkins credentials
	//
	//     - **rubygems_cred**: Rubygems credentials
	//
	//     - **clojars_token**: Clojars token
	//
	//     - **phoenix_web_passwd**: Phoenix Web credentials
	//
	//     - **puttygen_private_key**: PuTTYgen private key
	//
	//     - **google_oauth_token**: Google OAuth access token
	//
	//     - **rubyonrails_cfg**: Ruby On Rails database configuration
	//
	//     - **lob_api_key**: Lob API Key
	//
	//     - **pkcs_cred**: PKCS#12 certificate
	//
	//     - **otr_private_key**: OTR private key
	//
	//     - **contentful_delivery_token**: Contentful Delivery token
	//
	//     - **digital_ocean_tugboat**: Digital Ocean Tugboat configuration
	//
	//     - **dsa_private_key**: DSA private key
	//
	//     - **rails_app_token**: Rails App token
	//
	//     - **git_cred**: Git user credentials
	//
	//     - **newrelic_api_key**: New Relic User API Key
	//
	//     - **github_hub**: hub configuration that stores Github tokens
	//
	//     - **rubygem**: Rubygem token
	//
	// - **SensitiveFileName**: the sensitive file alerting type.
	//
	// example:
	//
	// SensitiveFileKey
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The page number of the page to return. Minimum value: **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The unique identifier of the image.
	//
	// > Call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation of Container Registry. You can obtain the unique identifier of the container image from the **ImageUuid*	- response parameter.
	//
	// example:
	//
	// 850613a48999900f48417c7e6e9d****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The language of the content in the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page in a paging query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The collection of scan scopes. Valid values:
	//
	// - **image**: image.
	//
	// - **container**: container.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The list of sensitive file keys.
	SensitiveKeyList []*string `json:"SensitiveKeyList,omitempty" xml:"SensitiveKeyList,omitempty" type:"Repeated"`
}

func (s DescribeImageSensitiveFileListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileListRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeImageSensitiveFileListRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *DescribeImageSensitiveFileListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageSensitiveFileListRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeImageSensitiveFileListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageSensitiveFileListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageSensitiveFileListRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageSensitiveFileListRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeImageSensitiveFileListRequest) GetSensitiveKeyList() []*string {
	return s.SensitiveKeyList
}

func (s *DescribeImageSensitiveFileListRequest) SetCriteria(v string) *DescribeImageSensitiveFileListRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeImageSensitiveFileListRequest) SetCriteriaType(v string) *DescribeImageSensitiveFileListRequest {
	s.CriteriaType = &v
	return s
}

func (s *DescribeImageSensitiveFileListRequest) SetCurrentPage(v int32) *DescribeImageSensitiveFileListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageSensitiveFileListRequest) SetImageUuid(v string) *DescribeImageSensitiveFileListRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeImageSensitiveFileListRequest) SetLang(v string) *DescribeImageSensitiveFileListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageSensitiveFileListRequest) SetPageSize(v int32) *DescribeImageSensitiveFileListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSensitiveFileListRequest) SetRiskLevel(v string) *DescribeImageSensitiveFileListRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageSensitiveFileListRequest) SetScanRange(v []*string) *DescribeImageSensitiveFileListRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeImageSensitiveFileListRequest) SetSensitiveKeyList(v []*string) *DescribeImageSensitiveFileListRequest {
	s.SensitiveKeyList = v
	return s
}

func (s *DescribeImageSensitiveFileListRequest) Validate() error {
	return dara.Validate(s)
}
