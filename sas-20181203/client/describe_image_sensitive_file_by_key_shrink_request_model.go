// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSensitiveFileByKeyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeImageSensitiveFileByKeyShrinkRequest
	GetCurrentPage() *int32
	SetImageUuid(v string) *DescribeImageSensitiveFileByKeyShrinkRequest
	GetImageUuid() *string
	SetLang(v string) *DescribeImageSensitiveFileByKeyShrinkRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageSensitiveFileByKeyShrinkRequest
	GetPageSize() *int32
	SetScanRangeShrink(v string) *DescribeImageSensitiveFileByKeyShrinkRequest
	GetScanRangeShrink() *string
	SetSensitiveFileKey(v string) *DescribeImageSensitiveFileByKeyShrinkRequest
	GetSensitiveFileKey() *string
}

type DescribeImageSensitiveFileByKeyShrinkRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// 0083a31ccf7c10367a6e783e8601****
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
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the asset that you want to scan. Valid values:
	//
	// 	- **image**
	//
	// 	- **container**
	ScanRangeShrink *string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	// The type of alerts for the sensitive files. Valid values:
	//
	// 	- **npm_token**: Node Package Manager (NPM) access token
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
	// 	- **aws_cli**: Amazon Web Services (AWS) CLI credentials
	//
	// 	- **cpanel_proftpd**: cPanel ProFTPD credentials
	//
	// 	- **postgresql_passwd**: PostgreSQL password
	//
	// 	- **discord_client_cred**: Discord client credentials
	//
	// 	- **rails_database**: Rails database configuration
	//
	// 	- **aws_access_key**: AWS Access Key
	//
	// 	- **esmtp_cfg*	- :Extended Simple Mail Transfer Protocol (ESMTP) configuration
	//
	// 	- **docker_registry_cfg**: Docker image repository configuration
	//
	// 	- **pem**: Privacy-Enhanced Mail (PEM)
	//
	// 	- **common_cred**: common credentials
	//
	// 	- **sftp_cfg**: Secure File Transfer Protocol (SFTP) connection configuration
	//
	// 	- **grafana_token**: Grafana token
	//
	// 	- **slack_token**: Slack token
	//
	// 	- **ec_private_key**: EC private key
	//
	// 	- **pypi_token**: Python Package Index (PyPI) token
	//
	// 	- **finicity_token**: Finicity token
	//
	// 	- **k8s_client_key**: Kubernetes private key
	//
	// 	- **git_cfg**: Git configuration
	//
	// 	- **django_key**: Django key
	//
	// 	- **jenkins_ssh**: Jenkins SSH configuration file
	//
	// 	- **openssh_private_key**: OpenSSL private key
	//
	// 	- **square_oauth**: OAuth credential for Square
	//
	// 	- **typeform_token**: Typeform token
	//
	// 	- **common_database_cfg**: general database connection configuration
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
	// 	- **robomongo_cred**: Robomongo credentials
	//
	// 	- **github_oauth_token**: OAuth access token for GitHub
	//
	// 	- **pulumi_token**: Pulumi token
	//
	// 	- **ventrilo_voip**: Ventrilo VoIP server configuration
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
	// 	- **kubernetes_dashboard_cred**: user credentials for Kubernetes Dashboard
	//
	// 	- **atlassian_token**: Atlassian token
	//
	// 	- **rdp**: remote desktop protocol (RDP)
	//
	// 	- **mailgun_key**: Mailgun webhook signing key
	//
	// 	- **mailchimp_api_key**: API key for Mailchimp
	//
	// 	- **netrc_cfg**: .netrc configuration file
	//
	// 	- **openvpn_cfg**: OpenVPN configuration
	//
	// 	- **github_refresh_token**: GitHub refresh token
	//
	// 	- **salesforce**: Salesforce credentials
	//
	// 	- **salesforce**: Sendinblue credentials
	//
	// 	- **pkcs_private_key**: PKCS#12 key
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
	// 	- **heroku_api_key**: Heroku API key
	//
	// 	- **messagebird_token**: MessageBird token
	//
	// 	- **messagebird_token**: MessageBird token
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
	// 	- **common_private_key**: private key of a common type
	//
	// 	- **microsoft_mdf**: Microsoft SQL Server database
	//
	// 	- **mediawiki_cfg**: MediaWiki configuration
	//
	// 	- **jenkins_cred**: Jenkins credentials
	//
	// 	- **rubygems_cred**: RubyGems credentials
	//
	// 	- **clojars_token**: Clojars token
	//
	// 	- **phoenix_web_passwd**: Phoenix web credentials
	//
	// 	- **puttygen_private_key**: PuTTYgen private key
	//
	// 	- **google_oauth_token**: Google OAuth access token
	//
	// 	- **rubyonrails_cfg**: Ruby On Rails database configuration
	//
	// 	- **lob_api_key**: Lob API key
	//
	// 	- **pkcs_cred**: PKCS#12 certificate
	//
	// 	- **otr_private_key**: Off-the-Record Messaging (OTR) private key
	//
	// 	- **contentful_delivery_token**: delivery token for Contentful
	//
	// 	- **digital_ocean_tugboat**: DigitalOcean Tugboat configuration
	//
	// 	- **dsa_private_key**: Digital Signature Algorithm (DSA) private key
	//
	// 	- **rails_app_token**: Rails app token
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
}

func (s DescribeImageSensitiveFileByKeyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileByKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) GetScanRangeShrink() *string {
	return s.ScanRangeShrink
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) GetSensitiveFileKey() *string {
	return s.SensitiveFileKey
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) SetCurrentPage(v int32) *DescribeImageSensitiveFileByKeyShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) SetImageUuid(v string) *DescribeImageSensitiveFileByKeyShrinkRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) SetLang(v string) *DescribeImageSensitiveFileByKeyShrinkRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) SetPageSize(v int32) *DescribeImageSensitiveFileByKeyShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) SetScanRangeShrink(v string) *DescribeImageSensitiveFileByKeyShrinkRequest {
	s.ScanRangeShrink = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) SetSensitiveFileKey(v string) *DescribeImageSensitiveFileByKeyShrinkRequest {
	s.SensitiveFileKey = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
