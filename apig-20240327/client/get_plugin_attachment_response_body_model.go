// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPluginAttachmentResponseBody
	GetCode() *string
	SetData(v *GetPluginAttachmentResponseBodyData) *GetPluginAttachmentResponseBody
	GetData() *GetPluginAttachmentResponseBodyData
	SetMessage(v string) *GetPluginAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPluginAttachmentResponseBody
	GetRequestId() *string
}

type GetPluginAttachmentResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response payload.
	Data *GetPluginAttachmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The status message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C61E30D3-579A-5B43-994E-31E02EDC9129
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPluginAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPluginAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetPluginAttachmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPluginAttachmentResponseBody) GetData() *GetPluginAttachmentResponseBodyData {
	return s.Data
}

func (s *GetPluginAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPluginAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPluginAttachmentResponseBody) SetCode(v string) *GetPluginAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetPluginAttachmentResponseBody) SetData(v *GetPluginAttachmentResponseBodyData) *GetPluginAttachmentResponseBody {
	s.Data = v
	return s
}

func (s *GetPluginAttachmentResponseBody) SetMessage(v string) *GetPluginAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetPluginAttachmentResponseBody) SetRequestId(v string) *GetPluginAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPluginAttachmentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPluginAttachmentResponseBodyData struct {
	// Indicates whether the plug-in is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The environment information.
	EnvironmentInfo *EnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty"`
	// The instance information.
	GatewayInfo *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// The information about the parent resource to which the plug-in is attached.
	ParentResourceInfo *ParentResourceInfo `json:"parentResourceInfo,omitempty" xml:"parentResourceInfo,omitempty"`
	// The attachment ID.
	//
	// example:
	//
	// pa-d05f1tmm1hku195dd8j0
	PluginAttachmentId *string `json:"pluginAttachmentId,omitempty" xml:"pluginAttachmentId,omitempty"`
	// The plug-in type information.
	PluginClassInfo *PluginClassInfo `json:"pluginClassInfo,omitempty" xml:"pluginClassInfo,omitempty"`
	// The Base64-encoded configurations of the plug-in.
	//
	// example:
	//
	// cHJlcGVuZDoKLSByb2xlOiBzeXN0ZW0KICBjb250ZW50OiDor7fkvb/nlKjoi7Hor63lm57nrZTpl67popgKYXBwZW5kOgotIHJvbGU6IHVzZXIKICBjb250ZW50OiDmr4/mrKHlm57nrZTlrozpl67popjvvIzlsJ3or5Xov5vooYzlj43pl64K
	PluginConfig *string `json:"pluginConfig,omitempty" xml:"pluginConfig,omitempty"`
	// The plug-in ID.
	//
	// example:
	//
	// pl-cvo8ub6m1hkvgv03r3k0
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// The resource details.
	ResourceInfos []*ResourceInfo `json:"resourceInfos,omitempty" xml:"resourceInfos,omitempty" type:"Repeated"`
}

func (s GetPluginAttachmentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPluginAttachmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPluginAttachmentResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetPluginAttachmentResponseBodyData) GetEnvironmentInfo() *EnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *GetPluginAttachmentResponseBodyData) GetGatewayInfo() *GatewayInfo {
	return s.GatewayInfo
}

func (s *GetPluginAttachmentResponseBodyData) GetParentResourceInfo() *ParentResourceInfo {
	return s.ParentResourceInfo
}

func (s *GetPluginAttachmentResponseBodyData) GetPluginAttachmentId() *string {
	return s.PluginAttachmentId
}

func (s *GetPluginAttachmentResponseBodyData) GetPluginClassInfo() *PluginClassInfo {
	return s.PluginClassInfo
}

func (s *GetPluginAttachmentResponseBodyData) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *GetPluginAttachmentResponseBodyData) GetPluginId() *string {
	return s.PluginId
}

func (s *GetPluginAttachmentResponseBodyData) GetResourceInfos() []*ResourceInfo {
	return s.ResourceInfos
}

func (s *GetPluginAttachmentResponseBodyData) SetEnable(v bool) *GetPluginAttachmentResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetPluginAttachmentResponseBodyData) SetEnvironmentInfo(v *EnvironmentInfo) *GetPluginAttachmentResponseBodyData {
	s.EnvironmentInfo = v
	return s
}

func (s *GetPluginAttachmentResponseBodyData) SetGatewayInfo(v *GatewayInfo) *GetPluginAttachmentResponseBodyData {
	s.GatewayInfo = v
	return s
}

func (s *GetPluginAttachmentResponseBodyData) SetParentResourceInfo(v *ParentResourceInfo) *GetPluginAttachmentResponseBodyData {
	s.ParentResourceInfo = v
	return s
}

func (s *GetPluginAttachmentResponseBodyData) SetPluginAttachmentId(v string) *GetPluginAttachmentResponseBodyData {
	s.PluginAttachmentId = &v
	return s
}

func (s *GetPluginAttachmentResponseBodyData) SetPluginClassInfo(v *PluginClassInfo) *GetPluginAttachmentResponseBodyData {
	s.PluginClassInfo = v
	return s
}

func (s *GetPluginAttachmentResponseBodyData) SetPluginConfig(v string) *GetPluginAttachmentResponseBodyData {
	s.PluginConfig = &v
	return s
}

func (s *GetPluginAttachmentResponseBodyData) SetPluginId(v string) *GetPluginAttachmentResponseBodyData {
	s.PluginId = &v
	return s
}

func (s *GetPluginAttachmentResponseBodyData) SetResourceInfos(v []*ResourceInfo) *GetPluginAttachmentResponseBodyData {
	s.ResourceInfos = v
	return s
}

func (s *GetPluginAttachmentResponseBodyData) Validate() error {
	if s.EnvironmentInfo != nil {
		if err := s.EnvironmentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.GatewayInfo != nil {
		if err := s.GatewayInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ParentResourceInfo != nil {
		if err := s.ParentResourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PluginClassInfo != nil {
		if err := s.PluginClassInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceInfos != nil {
		for _, item := range s.ResourceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
