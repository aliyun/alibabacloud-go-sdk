// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceArmsAlertConfig interface {
	dara.Model
	String() string
	GoString() string
	SetArmsAlertsEnable(v string) *GrafanaWorkspaceArmsAlertConfig
	GetArmsAlertsEnable() *string
	SetArmsAlertsWebhookUrl(v string) *GrafanaWorkspaceArmsAlertConfig
	GetArmsAlertsWebhookUrl() *string
}

type GrafanaWorkspaceArmsAlertConfig struct {
	// example:
	//
	// true
	ArmsAlertsEnable     *string `json:"armsAlertsEnable,omitempty" xml:"armsAlertsEnable,omitempty"`
	ArmsAlertsWebhookUrl *string `json:"armsAlertsWebhookUrl,omitempty" xml:"armsAlertsWebhookUrl,omitempty"`
}

func (s GrafanaWorkspaceArmsAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceArmsAlertConfig) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceArmsAlertConfig) GetArmsAlertsEnable() *string {
	return s.ArmsAlertsEnable
}

func (s *GrafanaWorkspaceArmsAlertConfig) GetArmsAlertsWebhookUrl() *string {
	return s.ArmsAlertsWebhookUrl
}

func (s *GrafanaWorkspaceArmsAlertConfig) SetArmsAlertsEnable(v string) *GrafanaWorkspaceArmsAlertConfig {
	s.ArmsAlertsEnable = &v
	return s
}

func (s *GrafanaWorkspaceArmsAlertConfig) SetArmsAlertsWebhookUrl(v string) *GrafanaWorkspaceArmsAlertConfig {
	s.ArmsAlertsWebhookUrl = &v
	return s
}

func (s *GrafanaWorkspaceArmsAlertConfig) Validate() error {
	return dara.Validate(s)
}
