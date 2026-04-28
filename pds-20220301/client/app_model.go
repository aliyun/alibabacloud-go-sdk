// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApp interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *App
	GetAppId() *string
	SetAppName(v string) *App
	GetAppName() *string
	SetAppSecret(v string) *App
	GetAppSecret() *string
	SetCreatedAt(v string) *App
	GetCreatedAt() *string
	SetDescription(v string) *App
	GetDescription() *string
	SetLogo(v string) *App
	GetLogo() *string
	SetProvider(v string) *App
	GetProvider() *string
	SetRedirectUri(v string) *App
	GetRedirectUri() *string
	SetScope(v []*string) *App
	GetScope() []*string
	SetStage(v string) *App
	GetStage() *string
	SetType(v string) *App
	GetType() *string
	SetUpdatedAt(v string) *App
	GetUpdatedAt() *string
}

type App struct {
	AppId       *string   `json:"app_id,omitempty" xml:"app_id,omitempty"`
	AppName     *string   `json:"app_name,omitempty" xml:"app_name,omitempty"`
	AppSecret   *string   `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
	CreatedAt   *string   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	Logo        *string   `json:"logo,omitempty" xml:"logo,omitempty"`
	Provider    *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	RedirectUri *string   `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	Scope       []*string `json:"scope,omitempty" xml:"scope,omitempty" type:"Repeated"`
	Stage       *string   `json:"stage,omitempty" xml:"stage,omitempty"`
	Type        *string   `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt   *string   `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s App) String() string {
	return dara.Prettify(s)
}

func (s App) GoString() string {
	return s.String()
}

func (s *App) GetAppId() *string {
	return s.AppId
}

func (s *App) GetAppName() *string {
	return s.AppName
}

func (s *App) GetAppSecret() *string {
	return s.AppSecret
}

func (s *App) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *App) GetDescription() *string {
	return s.Description
}

func (s *App) GetLogo() *string {
	return s.Logo
}

func (s *App) GetProvider() *string {
	return s.Provider
}

func (s *App) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *App) GetScope() []*string {
	return s.Scope
}

func (s *App) GetStage() *string {
	return s.Stage
}

func (s *App) GetType() *string {
	return s.Type
}

func (s *App) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *App) SetAppId(v string) *App {
	s.AppId = &v
	return s
}

func (s *App) SetAppName(v string) *App {
	s.AppName = &v
	return s
}

func (s *App) SetAppSecret(v string) *App {
	s.AppSecret = &v
	return s
}

func (s *App) SetCreatedAt(v string) *App {
	s.CreatedAt = &v
	return s
}

func (s *App) SetDescription(v string) *App {
	s.Description = &v
	return s
}

func (s *App) SetLogo(v string) *App {
	s.Logo = &v
	return s
}

func (s *App) SetProvider(v string) *App {
	s.Provider = &v
	return s
}

func (s *App) SetRedirectUri(v string) *App {
	s.RedirectUri = &v
	return s
}

func (s *App) SetScope(v []*string) *App {
	s.Scope = v
	return s
}

func (s *App) SetStage(v string) *App {
	s.Stage = &v
	return s
}

func (s *App) SetType(v string) *App {
	s.Type = &v
	return s
}

func (s *App) SetUpdatedAt(v string) *App {
	s.UpdatedAt = &v
	return s
}

func (s *App) Validate() error {
	return dara.Validate(s)
}
