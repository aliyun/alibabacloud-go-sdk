// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstallPluginResponseBody
	GetCode() *string
	SetData(v *InstallPluginResponseBodyData) *InstallPluginResponseBody
	GetData() *InstallPluginResponseBodyData
	SetMessage(v string) *InstallPluginResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallPluginResponseBody
	GetRequestId() *string
}

type InstallPluginResponseBody struct {
	// example:
	//
	// Ok
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *InstallPluginResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 201CFCA0-3AF5-52D0-A0F3-FBA697AF55CB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InstallPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallPluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallPluginResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallPluginResponseBody) GetData() *InstallPluginResponseBodyData {
	return s.Data
}

func (s *InstallPluginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallPluginResponseBody) SetCode(v string) *InstallPluginResponseBody {
	s.Code = &v
	return s
}

func (s *InstallPluginResponseBody) SetData(v *InstallPluginResponseBodyData) *InstallPluginResponseBody {
	s.Data = v
	return s
}

func (s *InstallPluginResponseBody) SetMessage(v string) *InstallPluginResponseBody {
	s.Message = &v
	return s
}

func (s *InstallPluginResponseBody) SetRequestId(v string) *InstallPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallPluginResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallPluginResponseBodyData struct {
	InstallPluginResults []*InstallPluginResponseBodyDataInstallPluginResults `json:"installPluginResults,omitempty" xml:"installPluginResults,omitempty" type:"Repeated"`
}

func (s InstallPluginResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InstallPluginResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallPluginResponseBodyData) GetInstallPluginResults() []*InstallPluginResponseBodyDataInstallPluginResults {
	return s.InstallPluginResults
}

func (s *InstallPluginResponseBodyData) SetInstallPluginResults(v []*InstallPluginResponseBodyDataInstallPluginResults) *InstallPluginResponseBodyData {
	s.InstallPluginResults = v
	return s
}

func (s *InstallPluginResponseBodyData) Validate() error {
	if s.InstallPluginResults != nil {
		for _, item := range s.InstallPluginResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallPluginResponseBodyDataInstallPluginResults struct {
	// example:
	//
	// gw-d28mjcmm1hkub84mdbi0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// pl-cvs7gbum1hkhs6us6vbg
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s InstallPluginResponseBodyDataInstallPluginResults) String() string {
	return dara.Prettify(s)
}

func (s InstallPluginResponseBodyDataInstallPluginResults) GoString() string {
	return s.String()
}

func (s *InstallPluginResponseBodyDataInstallPluginResults) GetGatewayId() *string {
	return s.GatewayId
}

func (s *InstallPluginResponseBodyDataInstallPluginResults) GetPluginId() *string {
	return s.PluginId
}

func (s *InstallPluginResponseBodyDataInstallPluginResults) SetGatewayId(v string) *InstallPluginResponseBodyDataInstallPluginResults {
	s.GatewayId = &v
	return s
}

func (s *InstallPluginResponseBodyDataInstallPluginResults) SetPluginId(v string) *InstallPluginResponseBodyDataInstallPluginResults {
	s.PluginId = &v
	return s
}

func (s *InstallPluginResponseBodyDataInstallPluginResults) Validate() error {
	return dara.Validate(s)
}
