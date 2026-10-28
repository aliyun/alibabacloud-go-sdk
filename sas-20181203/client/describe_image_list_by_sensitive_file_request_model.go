// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListBySensitiveFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeImageListBySensitiveFileRequest
	GetCurrentPage() *int32
	SetImageDigest(v string) *DescribeImageListBySensitiveFileRequest
	GetImageDigest() *string
	SetLang(v string) *DescribeImageListBySensitiveFileRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageListBySensitiveFileRequest
	GetPageSize() *int32
	SetRepoInstanceId(v string) *DescribeImageListBySensitiveFileRequest
	GetRepoInstanceId() *string
	SetRepoName(v string) *DescribeImageListBySensitiveFileRequest
	GetRepoName() *string
	SetRepoNamespace(v string) *DescribeImageListBySensitiveFileRequest
	GetRepoNamespace() *string
	SetRiskLevel(v string) *DescribeImageListBySensitiveFileRequest
	GetRiskLevel() *string
	SetScanRange(v []*string) *DescribeImageListBySensitiveFileRequest
	GetScanRange() []*string
	SetSensitiveFileKey(v string) *DescribeImageListBySensitiveFileRequest
	GetSensitiveFileKey() *string
	SetStatus(v string) *DescribeImageListBySensitiveFileRequest
	GetStatus() *string
}

type DescribeImageListBySensitiveFileRequest struct {
	// The page number of the returned page. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The image digest.
	//
	// > Fuzzy match is supported.
	//
	// example:
	//
	// v005
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page during paging in a paged query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the container image instance.
	//
	// > You can call the [ListRepository](https://help.aliyun.com/document_detail/451339.html) operation of Container Registry to obtain the container image instance ID from the **InstanceId*	- response parameter.
	//
	// example:
	//
	// i-qewqrqcsadf****
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	// The name of the image repository.
	//
	// > Fuzzy match is supported.
	//
	// example:
	//
	// harbor-image-v001
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace of the image repository.
	//
	// > Fuzzy match is supported.
	//
	// example:
	//
	// libssh2
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The risk level of the file. Separate multiple risk levels with commas (,). Valid values:
	//
	// - **high**: High.
	//
	// - **medium**: Medium.
	//
	// - **low**: Low.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The scan scope. Valid values:
	//
	// - **image**: Images.
	//
	// - **container**: Containers.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The alert type for sensitive files. Valid values:
	//
	// - **npm_token**: NPM access token
	//
	// - **ftp_cfg**: FTP configuration
	//
	// - **google_oauth_key**: Google OAuth key
	//
	// - **planetscale_passwd**: PlanetScale password
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
	// - **common_cred**: Common credential
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
	// - **common_database_cfg**: Common database connection configuration
	//
	// - **wordpress_database_cfg**: WordPress database configuration
	//
	// - **googlecloud_api_key**: Google Cloud API key
	//
	// - **vscode_sftp**: VSCode SFTP configuration
	//
	// - **apache_htpasswd**: Apache htpasswd
	//
	// - **planetscale_token**: PlanetScale token
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
	// - **tencentcloud_ak**: SecretId of other cloud providers
	//
	// - **rsa_private_key**: RSA private key
	//
	// - **github_personal_token**: GitHub personal access token
	//
	// - **pgp**: PGP encrypt file
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
	// - **huaweicloud_ak**: Secret Access Key of other cloud providers
	//
	// - **aws_s3cmd**: AWS S3cmd configuration
	//
	// - **php_config**: PHP configuration
	//
	// - **common_private_key**: Common private key type
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
	// > This parameter is required. If you do not specify this parameter, the API returns an HTTP 400 error with the error code -101.
	//
	// example:
	//
	// sshpasswd
	SensitiveFileKey *string `json:"SensitiveFileKey,omitempty" xml:"SensitiveFileKey,omitempty"`
	// The status of the sensitive file. Valid values:
	//
	// - **0**: Unhandled.
	//
	// - **1**: Ignored.
	//
	// - **2**: False positive.
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImageListBySensitiveFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListBySensitiveFileRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageListBySensitiveFileRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageListBySensitiveFileRequest) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeImageListBySensitiveFileRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageListBySensitiveFileRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageListBySensitiveFileRequest) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *DescribeImageListBySensitiveFileRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageListBySensitiveFileRequest) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageListBySensitiveFileRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageListBySensitiveFileRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeImageListBySensitiveFileRequest) GetSensitiveFileKey() *string {
	return s.SensitiveFileKey
}

func (s *DescribeImageListBySensitiveFileRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageListBySensitiveFileRequest) SetCurrentPage(v int32) *DescribeImageListBySensitiveFileRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetImageDigest(v string) *DescribeImageListBySensitiveFileRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetLang(v string) *DescribeImageListBySensitiveFileRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetPageSize(v int32) *DescribeImageListBySensitiveFileRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetRepoInstanceId(v string) *DescribeImageListBySensitiveFileRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetRepoName(v string) *DescribeImageListBySensitiveFileRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetRepoNamespace(v string) *DescribeImageListBySensitiveFileRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetRiskLevel(v string) *DescribeImageListBySensitiveFileRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetScanRange(v []*string) *DescribeImageListBySensitiveFileRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetSensitiveFileKey(v string) *DescribeImageListBySensitiveFileRequest {
	s.SensitiveFileKey = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) SetStatus(v string) *DescribeImageListBySensitiveFileRequest {
	s.Status = &v
	return s
}

func (s *DescribeImageListBySensitiveFileRequest) Validate() error {
	return dara.Validate(s)
}
