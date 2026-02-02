// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthAndroidId(v []*string) *CreateQosRuleRequest
	GetAuthAndroidId() []*string
	SetAuthDesktopId(v []*string) *CreateQosRuleRequest
	GetAuthDesktopId() []*string
	SetDownload(v int32) *CreateQosRuleRequest
	GetDownload() *int32
	SetNetworkPackageId(v string) *CreateQosRuleRequest
	GetNetworkPackageId() *string
	SetQosRuleName(v string) *CreateQosRuleRequest
	GetQosRuleName() *string
	SetUpload(v int32) *CreateQosRuleRequest
	GetUpload() *int32
}

type CreateQosRuleRequest struct {
	AuthAndroidId []*string `json:"AuthAndroidId,omitempty" xml:"AuthAndroidId,omitempty" type:"Repeated"`
	AuthDesktopId []*string `json:"AuthDesktopId,omitempty" xml:"AuthDesktopId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	Download *int32 `json:"Download,omitempty" xml:"Download,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// np-cfedn7r2pe48g****
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// This parameter is required.
	QosRuleName *string `json:"QosRuleName,omitempty" xml:"QosRuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	Upload *int32 `json:"Upload,omitempty" xml:"Upload,omitempty"`
}

func (s CreateQosRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQosRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateQosRuleRequest) GetAuthAndroidId() []*string {
	return s.AuthAndroidId
}

func (s *CreateQosRuleRequest) GetAuthDesktopId() []*string {
	return s.AuthDesktopId
}

func (s *CreateQosRuleRequest) GetDownload() *int32 {
	return s.Download
}

func (s *CreateQosRuleRequest) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *CreateQosRuleRequest) GetQosRuleName() *string {
	return s.QosRuleName
}

func (s *CreateQosRuleRequest) GetUpload() *int32 {
	return s.Upload
}

func (s *CreateQosRuleRequest) SetAuthAndroidId(v []*string) *CreateQosRuleRequest {
	s.AuthAndroidId = v
	return s
}

func (s *CreateQosRuleRequest) SetAuthDesktopId(v []*string) *CreateQosRuleRequest {
	s.AuthDesktopId = v
	return s
}

func (s *CreateQosRuleRequest) SetDownload(v int32) *CreateQosRuleRequest {
	s.Download = &v
	return s
}

func (s *CreateQosRuleRequest) SetNetworkPackageId(v string) *CreateQosRuleRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *CreateQosRuleRequest) SetQosRuleName(v string) *CreateQosRuleRequest {
	s.QosRuleName = &v
	return s
}

func (s *CreateQosRuleRequest) SetUpload(v int32) *CreateQosRuleRequest {
	s.Upload = &v
	return s
}

func (s *CreateQosRuleRequest) Validate() error {
	return dara.Validate(s)
}
