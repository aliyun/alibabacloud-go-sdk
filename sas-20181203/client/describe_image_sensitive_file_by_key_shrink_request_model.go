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
	// The page number of the returned page. Default value: **1**.
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
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page in a paging query. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scan scope. Valid values:
	//
	// - **image**: image
	//
	// - **container**: container
	ScanRangeShrink *string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	// The alerting type of the sensitive file. Valid values:
	//
	// - **npm_token**: NPM access token
	//
	// - **ftp_cfg**: FTP configuration
	//
	// - **google_oauth_key**: Google OAuth key
	//
	// - **planetscale_passwd**: Planetscale password
	//
	// - **github_ssh_key**: GitHub SSH key
	//
	// - **msbuild_publish_profile**: MSBuild publish profile
	//
	// - **fastly_cdn_token**: Fastly CDN token
	//
	// - **ssh_private_key**: SSH private key
	//
	// - **aws_cli**: AWS CLI credential
	//
	// - **cpanel_proftpd**: cPanel ProFTPd credential
	//
	// - **postgresql_passwd**: PostgreSQL password file
	//
	// - **discord_client_cred**: Discord client credential
	//
	// - **rails_database**: Rails database configuration
	//
	// - **aws_access_key**: AWS access key
	//
	// - **esmtp_cfg**: ESMTP mail server configuration
	//
	// - **docker_registry_cfg**: Docker image repository configuration
	//
	// - **pem**: PEM
	//
	// - **common_cred**: common credential
	//
	// - **sftp_cfg**: SFTP connection configuration
	//
	// - **grafana_token**: Grafana token
	//
	// - **slack_token**: Slack token
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
	// - **openssh_private_key**: OpenSSH private key
	//
	// - **square_oauth**: Square OAuth credential
	//
	// - **typeform_token**: Typeform token
	//
	// - **common_database_cfg**: common database connection configuration
	//
	// - **wordpress_database_cfg**: WordPress database configuration
	//
	// - **googlecloud_api_key**: Google Cloud API key
	//
	// - **vscode_sftp**: VSCode SFTP configuration
	//
	// - **apache_htpasswd**: Apache htpasswd
	//
	// - **planetscale_token**: Planetscale token
	//
	// - **contentful_preview_token**: Contentful preview token
	//
	// - **php_database_cfg**: PHP application database password
	//
	// - **atom_remote_sync**: Atom remote synchronization configuration
	//
	// - **aws_session_token**: AWS session token
	//
	// - **atom_sftp_cfg**: Atom SFTP configuration
	//
	// - **asana_client_private_key**: Asana project management platform client key
	//
	// - **tencentcloud_ak**: SecretId of other cloud service providers
	//
	// - **rsa_private_key**: RSA private key
	//
	// - **github_personal_token**: GitHub personal access token
	//
	// - **pgp**: PGP encrypted file
	//
	// - **stripe_skpk**: Stripe secret key
	//
	// - **square_token**: Square access token
	//
	// - **rails_carrierwave**: Rails Carrierwave file upload credential
	//
	// - **dbeaver_database_cfg**: DBeaver database configuration
	//
	// - **robomongo_cred**: Robomongo credential
	//
	// - **github_oauth_token**: GitHub OAuth access token
	//
	// - **pulumi_token**: Pulumi token
	//
	// - **ventrilo_voip**: Ventrilo VoIP server configuration
	//
	// - **macos_keychain**: macOS keychain
	//
	// - **amazon_mws_token**: Amazon MWS token
	//
	// - **dynatrace_token**: Dynatrace token
	//
	// - **java_keystore**: Java keystore
	//
	// - **microsoft_sdf**: Microsoft SQL CE database
	//
	// - **kubernetes_dashboard_cred**: Kubernetes Dashboard user credential
	//
	// - **atlassian_token**: Atlassian token
	//
	// - **rdp**: Remote Desktop Protocol (RDP) connection
	//
	// - **mailgun_key**: Mailgun webhook signing key
	//
	// - **mailchimp_api_key**: Mailchimp API key
	//
	// - **netrc_cfg**: .netrc configuration file
	//
	// - **openvpn_cfg**: OpenVPN client configuration
	//
	// - **github_refresh_token**: GitHub refresh token
	//
	// - **salesforce**: Salesforce credential
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
	// - **gitLab_personal_token**: GitLab personal access token
	//
	// - **rails_master_key**: Rails master key
	//
	// - **sqlite**: SQLite3/SQLite database
	//
	// - **firefox_logins**: Firefox logon configuration
	//
	// - **mailgun_private_token**: Mailgun private token
	//
	// - **joomla_cfg**: Joomla configuration
	//
	// - **hashicorp_terraform_token**: HashiCorp Terraform token
	//
	// - **jetbrains_ides**: JetBrains IDEs configuration
	//
	// - **heroku_api_key**: Heroku API key
	//
	// - **messagebird_token**: MessageBird token
	//
	// - **github_app_token**: GitHub app token
	//
	// - **hashicorp_vault_token**: HashiCorp Vault token
	//
	// - **pgp_private_key**: PGP private key
	//
	// - **sshpasswd**: SSH password
	//
	// - **huaweicloud_ak**: Secret access key of other cloud service providers
	//
	// - **aws_s3cmd**: AWS S3cmd configuration
	//
	// - **php_config**: PHP configuration
	//
	// - **common_private_key**: common private key type
	//
	// - **microsoft_mdf**: Microsoft SQL database
	//
	// - **mediawiki_cfg**: MediaWiki configuration
	//
	// - **jenkins_cred**: Jenkins credential
	//
	// - **rubygems_cred**: RubyGems credential
	//
	// - **clojars_token**: Clojars token
	//
	// - **phoenix_web_passwd**: Phoenix web credential
	//
	// - **puttygen_private_key**: PuTTYgen private key
	//
	// - **google_oauth_token**: Google OAuth access token
	//
	// - **rubyonrails_cfg**: Ruby on Rails database configuration
	//
	// - **lob_api_key**: Lob API key
	//
	// - **pkcs_cred**: PKCS#12 certificate
	//
	// - **otr_private_key**: OTR private key
	//
	// - **contentful_delivery_token**: Contentful delivery token
	//
	// - **digital_ocean_tugboat**: DigitalOcean Tugboat configuration
	//
	// - **dsa_private_key**: DSA private key
	//
	// - **rails_app_token**: Rails app token
	//
	// - **git_cred**: Git user credential
	//
	// - **newrelic_api_key**: New Relic user API key
	//
	// - **github_hub**: Hub configuration that stores GitHub tokens
	//
	// - **rubygem**: RubyGems token
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
