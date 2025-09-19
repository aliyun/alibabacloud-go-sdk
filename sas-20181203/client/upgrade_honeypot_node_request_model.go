// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeHoneypotNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowHoneypotAccessInternet(v bool) *UpgradeHoneypotNodeRequest
	GetAllowHoneypotAccessInternet() *bool
	SetLang(v string) *UpgradeHoneypotNodeRequest
	GetLang() *string
	SetNodeId(v string) *UpgradeHoneypotNodeRequest
	GetNodeId() *string
}

type UpgradeHoneypotNodeRequest struct {
	// Specifies whether to allow the honeypot to access the Internet. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AllowHoneypotAccessInternet *bool `json:"AllowHoneypotAccessInternet,omitempty" xml:"AllowHoneypotAccessInternet,omitempty"`
	// The language of the content within the request and response.
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the management node that you want to upgrade.
	//
	// >  You can call the [ListHoneypotNode](~~ListHoneypotNode~~) operation to obtain the ID.
	//
	// example:
	//
	// a882e590-b87b-45a6-87b9-d0a3e5a0****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s UpgradeHoneypotNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeHoneypotNodeRequest) GoString() string {
	return s.String()
}

func (s *UpgradeHoneypotNodeRequest) GetAllowHoneypotAccessInternet() *bool {
	return s.AllowHoneypotAccessInternet
}

func (s *UpgradeHoneypotNodeRequest) GetLang() *string {
	return s.Lang
}

func (s *UpgradeHoneypotNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpgradeHoneypotNodeRequest) SetAllowHoneypotAccessInternet(v bool) *UpgradeHoneypotNodeRequest {
	s.AllowHoneypotAccessInternet = &v
	return s
}

func (s *UpgradeHoneypotNodeRequest) SetLang(v string) *UpgradeHoneypotNodeRequest {
	s.Lang = &v
	return s
}

func (s *UpgradeHoneypotNodeRequest) SetNodeId(v string) *UpgradeHoneypotNodeRequest {
	s.NodeId = &v
	return s
}

func (s *UpgradeHoneypotNodeRequest) Validate() error {
	return dara.Validate(s)
}
