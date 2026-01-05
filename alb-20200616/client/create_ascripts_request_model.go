// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAScripts(v []*CreateAScriptsRequestAScripts) *CreateAScriptsRequest
	GetAScripts() []*CreateAScriptsRequestAScripts
	SetClientToken(v string) *CreateAScriptsRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateAScriptsRequest
	GetDryRun() *bool
	SetListenerId(v string) *CreateAScriptsRequest
	GetListenerId() *string
}

type CreateAScriptsRequest struct {
	// The information about the AScript rules.
	AScripts []*CreateAScriptsRequestAScripts `json:"AScripts,omitempty" xml:"AScripts,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-6hfq3zs0x04ibn****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s CreateAScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAScriptsRequest) GoString() string {
	return s.String()
}

func (s *CreateAScriptsRequest) GetAScripts() []*CreateAScriptsRequestAScripts {
	return s.AScripts
}

func (s *CreateAScriptsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAScriptsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAScriptsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateAScriptsRequest) SetAScripts(v []*CreateAScriptsRequestAScripts) *CreateAScriptsRequest {
	s.AScripts = v
	return s
}

func (s *CreateAScriptsRequest) SetClientToken(v string) *CreateAScriptsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAScriptsRequest) SetDryRun(v bool) *CreateAScriptsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAScriptsRequest) SetListenerId(v string) *CreateAScriptsRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateAScriptsRequest) Validate() error {
	if s.AScripts != nil {
		for _, item := range s.AScripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAScriptsRequestAScripts struct {
	// The name of the AScript rule.
	//
	// The length must be between 2 and 128 characters. This name must start with a letter and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AScriptName *string `json:"AScriptName,omitempty" xml:"AScriptName,omitempty"`
	// Enables the AScript rule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Enables the extended attribute of the Ascript rule. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// true
	ExtAttributeEnabled *bool `json:"ExtAttributeEnabled,omitempty" xml:"ExtAttributeEnabled,omitempty"`
	// The extended attribute of the AScript rule.
	ExtAttributes []*CreateAScriptsRequestAScriptsExtAttributes `json:"ExtAttributes,omitempty" xml:"ExtAttributes,omitempty" type:"Repeated"`
	// The position where the Ascript rule is evaluated. Valid values are:
	//
	// 	- RequestHead (default): before inbound rules are evaluated
	//
	// 	- RequestFoot: after inbound rules are evaluated
	//
	// 	- ResponseHead: before outbound rules are evaluated
	//
	// example:
	//
	// RequestFoot
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The content of the AScript rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// if and(match_re($uri, \\"^/1.txt$\\"), $arg_type) {   rewrite(concat(\\"/1.\\", $arg_type), \\"break\\") }
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
}

func (s CreateAScriptsRequestAScripts) String() string {
	return dara.Prettify(s)
}

func (s CreateAScriptsRequestAScripts) GoString() string {
	return s.String()
}

func (s *CreateAScriptsRequestAScripts) GetAScriptName() *string {
	return s.AScriptName
}

func (s *CreateAScriptsRequestAScripts) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAScriptsRequestAScripts) GetExtAttributeEnabled() *bool {
	return s.ExtAttributeEnabled
}

func (s *CreateAScriptsRequestAScripts) GetExtAttributes() []*CreateAScriptsRequestAScriptsExtAttributes {
	return s.ExtAttributes
}

func (s *CreateAScriptsRequestAScripts) GetPosition() *string {
	return s.Position
}

func (s *CreateAScriptsRequestAScripts) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *CreateAScriptsRequestAScripts) SetAScriptName(v string) *CreateAScriptsRequestAScripts {
	s.AScriptName = &v
	return s
}

func (s *CreateAScriptsRequestAScripts) SetEnabled(v bool) *CreateAScriptsRequestAScripts {
	s.Enabled = &v
	return s
}

func (s *CreateAScriptsRequestAScripts) SetExtAttributeEnabled(v bool) *CreateAScriptsRequestAScripts {
	s.ExtAttributeEnabled = &v
	return s
}

func (s *CreateAScriptsRequestAScripts) SetExtAttributes(v []*CreateAScriptsRequestAScriptsExtAttributes) *CreateAScriptsRequestAScripts {
	s.ExtAttributes = v
	return s
}

func (s *CreateAScriptsRequestAScripts) SetPosition(v string) *CreateAScriptsRequestAScripts {
	s.Position = &v
	return s
}

func (s *CreateAScriptsRequestAScripts) SetScriptContent(v string) *CreateAScriptsRequestAScripts {
	s.ScriptContent = &v
	return s
}

func (s *CreateAScriptsRequestAScripts) Validate() error {
	if s.ExtAttributes != nil {
		for _, item := range s.ExtAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAScriptsRequestAScriptsExtAttributes struct {
	// The key of the extended attribute.
	//
	// You can only set the key to **EsDebug**. This extended attribute adds a debug response header to record the execution of the AScript rule if the client request includes the _es_dbg parameter and its value matches the specified value of the extended attribute.
	//
	// example:
	//
	// EsDebug
	AttributeKey *string `json:"AttributeKey,omitempty" xml:"AttributeKey,omitempty"`
	// The value of the extended attribute, which can contain a maximum of 128 characters, including letters and digits.
	//
	// example:
	//
	// test123
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s CreateAScriptsRequestAScriptsExtAttributes) String() string {
	return dara.Prettify(s)
}

func (s CreateAScriptsRequestAScriptsExtAttributes) GoString() string {
	return s.String()
}

func (s *CreateAScriptsRequestAScriptsExtAttributes) GetAttributeKey() *string {
	return s.AttributeKey
}

func (s *CreateAScriptsRequestAScriptsExtAttributes) GetAttributeValue() *string {
	return s.AttributeValue
}

func (s *CreateAScriptsRequestAScriptsExtAttributes) SetAttributeKey(v string) *CreateAScriptsRequestAScriptsExtAttributes {
	s.AttributeKey = &v
	return s
}

func (s *CreateAScriptsRequestAScriptsExtAttributes) SetAttributeValue(v string) *CreateAScriptsRequestAScriptsExtAttributes {
	s.AttributeValue = &v
	return s
}

func (s *CreateAScriptsRequestAScriptsExtAttributes) Validate() error {
	return dara.Validate(s)
}
