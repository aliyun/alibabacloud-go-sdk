// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSensitiveFileListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeImageSensitiveFileListShrinkRequest
	GetCriteria() *string
	SetCriteriaType(v string) *DescribeImageSensitiveFileListShrinkRequest
	GetCriteriaType() *string
	SetCurrentPage(v int32) *DescribeImageSensitiveFileListShrinkRequest
	GetCurrentPage() *int32
	SetImageUuid(v string) *DescribeImageSensitiveFileListShrinkRequest
	GetImageUuid() *string
	SetLang(v string) *DescribeImageSensitiveFileListShrinkRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageSensitiveFileListShrinkRequest
	GetPageSize() *int32
	SetRiskLevel(v string) *DescribeImageSensitiveFileListShrinkRequest
	GetRiskLevel() *string
	SetScanRangeShrink(v string) *DescribeImageSensitiveFileListShrinkRequest
	GetScanRangeShrink() *string
	SetSensitiveKeyList(v []*string) *DescribeImageSensitiveFileListShrinkRequest
	GetSensitiveKeyList() []*string
}

type DescribeImageSensitiveFileListShrinkRequest struct {
	// The value of the sensitive file type.
	//
	// example:
	//
	// Rails Master Key
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The type of the sensitive files that you want to query. Valid values:
	//
	// 	- **SensitiveFileKey**: the type of alerts for sensitive files. Valid values:
	//
	//     	- **npm_token**: Node Package Manager (NPM) access token
	//
	//     	- **ftp_cfg**: FTP configuration
	//
	//     	- **google_oauth_key**: Google OAuth key
	//
	//     	- **planetscale_passwd**: PlanetScale password
	//
	//     	- **github_ssh_key**: Github SSH key
	//
	//     	- **msbuild_publish_profile**: MSBuild publish profile
	//
	//     	- **fastly_cdn_token**: Fastly CDN token
	//
	//     	- **ssh_private_key**: SSH private key
	//
	//     	- **aws_cli**: Amazon Web Services (AWS) CLI credentials
	//
	//     	- **cpanel_proftpd**: cPanel ProFTPD credentials
	//
	//     	- **postgresql_passwd**: PostgreSQL password
	//
	//     	- **discord_client_cred**: Discord client credentials
	//
	//     	- **rails_database**: Rails database configuration
	//
	//     	- **aws_access_key**: AWS access key
	//
	//     	- **esmtp_cfg**: Extended Simple Mail Transfer Protocol (ESMTP) configuration
	//
	//     	- **docker_registry_cfg**: Docker image repository configuration
	//
	//     	- **pem**: Privacy-Enhanced Mail (PEM)
	//
	//     	- **common_cred**: common credentials
	//
	//     	- **sftp_cfg**: Secure File Transfer Protocol (SFTP) connection configuration
	//
	//     	- **grafana_token**: Grafana token
	//
	//     	- **slack_token**: Slack token
	//
	//     	- **ec_private_key**: EC private key
	//
	//     	- **pypi_token**: upload token for the Python Package Index (PyPI)
	//
	//     	- **finicity_token**: Finicity token
	//
	//     	- **k8s_client_key**: Kubernetes private key
	//
	//     	- **git_cfg**: Git configuration
	//
	//     	- **django_key**: Django key
	//
	//     	- **jenkins_ssh**: Jenkins SSH configuration file
	//
	//     	- **openssh_private_key**: OpenSSL private key
	//
	//     	- **square_oauth**: OAuth credential for Square
	//
	//     	- **typeform_token**: Typeform token
	//
	//     	- **common_database_cfg**: general database connection configuration
	//
	//     	- **wordpress_database_cfg**: WordPress database configuration
	//
	//     	- **googlecloud_api_key**: API key for Google Cloud
	//
	//     	- **vscode_sftp**: VSCode SFTP configuration
	//
	//     	- **apache_htpasswd**: Apache htpasswd
	//
	//     	- **planetscale_token**: PlanetScale token
	//
	//     	- **contentful_preview_token**: preview token for Contentful
	//
	//     	- **php_database_cfg**: database password for a PHP application
	//
	//     	- **atom_remote_sync**: Atom remote synchronization configuration
	//
	//     	- **aws_session_token**: AWS session token
	//
	//     	- **atom_sftp_cfg**: Atom SFTP configuration
	//
	//     	- **asana_client_private_key**: Asana client key
	//
	//     	- **tencentcloud_ak**: secret ID of a third-party cloud
	//
	//     	- **rsa_private_key**: Rivest-Shamir-Adleman (RSA) private key
	//
	//     	- **github_personal_token**: personal access token for GitHub
	//
	//     	- **pgp**: Pretty Good Privacy (PGP) encrypted file
	//
	//     	- **stripe_skpk**: Stripe secret key
	//
	//     	- **square_token**: Square access token
	//
	//     	- **rails_carrierwave**: file upload credentials for Rails Carrierwave
	//
	//     	- **dbeaver_database_cfg**: DBeaver database configuration
	//
	//     	- **robomongo_cred**: Robomongo credentials
	//
	//     	- **github_oauth_token**: OAuth access token for GitHub
	//
	//     	- **pulumi_token**: Pulumi token
	//
	//     	- **ventrilo_voip**: Ventrilo VoIP server configuration
	//
	//     	- **macos_keychain*	- :macOS keychain
	//
	//     	- **amazon_mws_token**: Amazon MWS token
	//
	//     	- **dynatrace_token**: Dynatrace token
	//
	//     	- **java_keystore**: Java KeyStore (JKS)
	//
	//     	- **microsoft_sdf**: Microsoft SQL Server Compact Edition (CE) database
	//
	//     	- **kubernetes_dashboard_cred**: user credentials for Kubernetes Dashboard
	//
	//     	- **atlassian_token**: Atlassian token
	//
	//     	- **rdp**: remote desktop protocol (RDP)
	//
	//     	- **mailgun_key**: Mailgun webhook signing key
	//
	//     	- **mailchimp_api_key**: API key for Mailchimp
	//
	//     	- **netrc_cfg**: .netrc configuration file
	//
	//     	- **openvpn_cfg**: OpenVPN configuration
	//
	//     	- **github_refresh_token**: GitHub refresh token
	//
	//     	- **salesforce**: Salesforce credentials
	//
	//     	- **salesforce**: Sendinblue credentials
	//
	//     	- **pkcs_private_key**: PKCS#12 key
	//
	//     	- **rubyonrails_passwd**: Ruby on Rails password file
	//
	//     	- **filezilla_ftp**: FileZilla FTP configuration
	//
	//     	- **databricks_token**: Databricks token
	//
	//     	- **gitLab_personal_toke**: personal access token for GitLab
	//
	//     	- **rails_master_key**: Rails master key
	//
	//     	- **sqlite**: SQLite3 or SQLite database
	//
	//     	- **firefox_logins**: Firefox logon configuration
	//
	//     	- **mailgun_private_token**: Mailgun private token
	//
	//     	- **joomla_cfg**: Joomla configuration
	//
	//     	- **hashicorp_terraform_token**: HashiCorp Terraform token
	//
	//     	- **jetbrains_ides**: JetBrains IDEs configuration
	//
	//     	- **heroku_api_key**: Heroku API key
	//
	//     	- **messagebird_token**: MessageBird token
	//
	//     	- **messagebird_token**: MessageBird token
	//
	//     	- **hashicorp_vault_token**: HashiCorp Vault token
	//
	//     	- **pgp_private_key**: PGP private key
	//
	//     	- **sshpasswd**: SSH password
	//
	//     	- **huaweicloud_ak**: secret access key of a third-party cloud
	//
	//     	- **aws_s3cmd**: AWS S3cmd configuration
	//
	//     	- **php_config**: PHP configuration
	//
	//     	- **common_private_key**: private key of a common type
	//
	//     	- **microsoft_mdf**: Microsoft SQL Server database
	//
	//     	- **mediawiki_cfg**: MediaWiki configuration
	//
	//     	- **jenkins_cred**: Jenkins credentials
	//
	//     	- **rubygems_cred**: RubyGems credentials
	//
	//     	- **clojars_token**: Clojars token
	//
	//     	- **phoenix_web_passwd**: Phoenix web credentials
	//
	//     	- **puttygen_private_key**: PuTTYgen private key
	//
	//     	- **google_oauth_token**: Google OAuth access token
	//
	//     	- **rubyonrails_cfg**: Ruby On Rails database configuration
	//
	//     	- **lob_api_key**: Lob API key
	//
	//     	- **pkcs_cred**: PKCS#12 certificate
	//
	//     	- **otr_private_key**: Off-the-Record Messaging (OTR) private key
	//
	//     	- **contentful_delivery_token**: delivery token for Contentful
	//
	//     	- **digital_ocean_tugboat**: DigitalOcean Tugboat configuration
	//
	//     	- **dsa_private_key**: Digital Signature Algorithm (DSA) private key
	//
	//     	- **rails_app_token**: Rails app token
	//
	//     	- **git_cred**: Git user credential
	//
	//     	- **newrelic_api_key**: User API key for New Relic
	//
	//     	- **github_hub**: hub configuration for storing GitHub tokens
	//
	//     	- **rubygem**: RubyGem token
	//
	// 	- **SensitiveFileName**: the name of the alert type for sensitive files.
	//
	// example:
	//
	// SensitiveFileKey
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the image.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation of Container Registry to query the image UUID from the value of the **ImageUuid*	- response parameter.
	//
	// example:
	//
	// 850613a48999900f48417c7e6e9dcfdd
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// An array that consists of the types of the assets that you want to scan. Valid values:
	//
	// 	- **image**
	//
	// 	- **container**
	ScanRangeShrink  *string   `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	SensitiveKeyList []*string `json:"SensitiveKeyList,omitempty" xml:"SensitiveKeyList,omitempty" type:"Repeated"`
}

func (s DescribeImageSensitiveFileListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileListShrinkRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeImageSensitiveFileListShrinkRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *DescribeImageSensitiveFileListShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageSensitiveFileListShrinkRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeImageSensitiveFileListShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageSensitiveFileListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageSensitiveFileListShrinkRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageSensitiveFileListShrinkRequest) GetScanRangeShrink() *string {
	return s.ScanRangeShrink
}

func (s *DescribeImageSensitiveFileListShrinkRequest) GetSensitiveKeyList() []*string {
	return s.SensitiveKeyList
}

func (s *DescribeImageSensitiveFileListShrinkRequest) SetCriteria(v string) *DescribeImageSensitiveFileListShrinkRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeImageSensitiveFileListShrinkRequest) SetCriteriaType(v string) *DescribeImageSensitiveFileListShrinkRequest {
	s.CriteriaType = &v
	return s
}

func (s *DescribeImageSensitiveFileListShrinkRequest) SetCurrentPage(v int32) *DescribeImageSensitiveFileListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageSensitiveFileListShrinkRequest) SetImageUuid(v string) *DescribeImageSensitiveFileListShrinkRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeImageSensitiveFileListShrinkRequest) SetLang(v string) *DescribeImageSensitiveFileListShrinkRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageSensitiveFileListShrinkRequest) SetPageSize(v int32) *DescribeImageSensitiveFileListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSensitiveFileListShrinkRequest) SetRiskLevel(v string) *DescribeImageSensitiveFileListShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageSensitiveFileListShrinkRequest) SetScanRangeShrink(v string) *DescribeImageSensitiveFileListShrinkRequest {
	s.ScanRangeShrink = &v
	return s
}

func (s *DescribeImageSensitiveFileListShrinkRequest) SetSensitiveKeyList(v []*string) *DescribeImageSensitiveFileListShrinkRequest {
	s.SensitiveKeyList = v
	return s
}

func (s *DescribeImageSensitiveFileListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
