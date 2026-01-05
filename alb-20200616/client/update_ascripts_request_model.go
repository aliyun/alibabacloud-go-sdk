// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAScripts(v []*UpdateAScriptsRequestAScripts) *UpdateAScriptsRequest
	GetAScripts() []*UpdateAScriptsRequestAScripts
	SetClientToken(v string) *UpdateAScriptsRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateAScriptsRequest
	GetDryRun() *bool
}

type UpdateAScriptsRequest struct {
	// The information about the AScript rule.
	AScripts []*UpdateAScriptsRequestAScripts `json:"AScripts,omitempty" xml:"AScripts,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// b1f642ac-5558-4a36-b7d9-cf53f40ea5c8
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
}

func (s UpdateAScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAScriptsRequest) GoString() string {
	return s.String()
}

func (s *UpdateAScriptsRequest) GetAScripts() []*UpdateAScriptsRequestAScripts {
	return s.AScripts
}

func (s *UpdateAScriptsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAScriptsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateAScriptsRequest) SetAScripts(v []*UpdateAScriptsRequestAScripts) *UpdateAScriptsRequest {
	s.AScripts = v
	return s
}

func (s *UpdateAScriptsRequest) SetClientToken(v string) *UpdateAScriptsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAScriptsRequest) SetDryRun(v bool) *UpdateAScriptsRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAScriptsRequest) Validate() error {
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

type UpdateAScriptsRequestAScripts struct {
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// as-mhqxcanmivn4g5****
	AScriptId *string `json:"AScriptId,omitempty" xml:"AScriptId,omitempty"`
	// The name of the AScript rule.
	//
	// The name must be 2 to 128 character in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// Group1
	AScriptName *string `json:"AScriptName,omitempty" xml:"AScriptName,omitempty"`
	// Specifies whether to enable the AScript rule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Specifies whether to enable the extended attributes of the Ascript rule. Valid values:
	//
	// 	- true
	//
	// 	- false (false)
	//
	// example:
	//
	// true
	ExtAttributeEnabled *bool `json:"ExtAttributeEnabled,omitempty" xml:"ExtAttributeEnabled,omitempty"`
	// The extended attribute.
	ExtAttributes []*UpdateAScriptsRequestAScriptsExtAttributes `json:"ExtAttributes,omitempty" xml:"ExtAttributes,omitempty" type:"Repeated"`
	// The content of the AScript rule.
	//
	// example:
	//
	// if and(match_re($uri, \\"^/1.txt$\\"), $arg_type) { rewrite(concat(\\"/1.\\", $arg_type), \\"break\\") }
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
}

func (s UpdateAScriptsRequestAScripts) String() string {
	return dara.Prettify(s)
}

func (s UpdateAScriptsRequestAScripts) GoString() string {
	return s.String()
}

func (s *UpdateAScriptsRequestAScripts) GetAScriptId() *string {
	return s.AScriptId
}

func (s *UpdateAScriptsRequestAScripts) GetAScriptName() *string {
	return s.AScriptName
}

func (s *UpdateAScriptsRequestAScripts) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAScriptsRequestAScripts) GetExtAttributeEnabled() *bool {
	return s.ExtAttributeEnabled
}

func (s *UpdateAScriptsRequestAScripts) GetExtAttributes() []*UpdateAScriptsRequestAScriptsExtAttributes {
	return s.ExtAttributes
}

func (s *UpdateAScriptsRequestAScripts) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *UpdateAScriptsRequestAScripts) SetAScriptId(v string) *UpdateAScriptsRequestAScripts {
	s.AScriptId = &v
	return s
}

func (s *UpdateAScriptsRequestAScripts) SetAScriptName(v string) *UpdateAScriptsRequestAScripts {
	s.AScriptName = &v
	return s
}

func (s *UpdateAScriptsRequestAScripts) SetEnabled(v bool) *UpdateAScriptsRequestAScripts {
	s.Enabled = &v
	return s
}

func (s *UpdateAScriptsRequestAScripts) SetExtAttributeEnabled(v bool) *UpdateAScriptsRequestAScripts {
	s.ExtAttributeEnabled = &v
	return s
}

func (s *UpdateAScriptsRequestAScripts) SetExtAttributes(v []*UpdateAScriptsRequestAScriptsExtAttributes) *UpdateAScriptsRequestAScripts {
	s.ExtAttributes = v
	return s
}

func (s *UpdateAScriptsRequestAScripts) SetScriptContent(v string) *UpdateAScriptsRequestAScripts {
	s.ScriptContent = &v
	return s
}

func (s *UpdateAScriptsRequestAScripts) Validate() error {
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

type UpdateAScriptsRequestAScriptsExtAttributes struct {
	// The attribute name.
	//
	// Set the value to **EsDebug**, which specifies that when requests carry the _es_dbg parameter whose value is the specified key, the debugging header is enabled to output the execution result.
	//
	// This parameter is required.
	//
	// example:
	//
	// EsDebug
	AttributeKey *string `json:"AttributeKey,omitempty" xml:"AttributeKey,omitempty"`
	// The attribute value, which must be 1 to 128 characters in length, and can contain letters and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// test123
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s UpdateAScriptsRequestAScriptsExtAttributes) String() string {
	return dara.Prettify(s)
}

func (s UpdateAScriptsRequestAScriptsExtAttributes) GoString() string {
	return s.String()
}

func (s *UpdateAScriptsRequestAScriptsExtAttributes) GetAttributeKey() *string {
	return s.AttributeKey
}

func (s *UpdateAScriptsRequestAScriptsExtAttributes) GetAttributeValue() *string {
	return s.AttributeValue
}

func (s *UpdateAScriptsRequestAScriptsExtAttributes) SetAttributeKey(v string) *UpdateAScriptsRequestAScriptsExtAttributes {
	s.AttributeKey = &v
	return s
}

func (s *UpdateAScriptsRequestAScriptsExtAttributes) SetAttributeValue(v string) *UpdateAScriptsRequestAScriptsExtAttributes {
	s.AttributeValue = &v
	return s
}

func (s *UpdateAScriptsRequestAScriptsExtAttributes) Validate() error {
	return dara.Validate(s)
}
