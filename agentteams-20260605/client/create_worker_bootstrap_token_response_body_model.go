// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkerBootstrapTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWorkerBootstrapTokenResponseBody
	GetCode() *string
	SetData(v *CreateWorkerBootstrapTokenResponseBodyData) *CreateWorkerBootstrapTokenResponseBody
	GetData() *CreateWorkerBootstrapTokenResponseBodyData
	SetHttpStatusCode(v int32) *CreateWorkerBootstrapTokenResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateWorkerBootstrapTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWorkerBootstrapTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkerBootstrapTokenResponseBody
	GetSuccess() *bool
}

type CreateWorkerBootstrapTokenResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateWorkerBootstrapTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWorkerBootstrapTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerBootstrapTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkerBootstrapTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkerBootstrapTokenResponseBody) GetData() *CreateWorkerBootstrapTokenResponseBodyData {
	return s.Data
}

func (s *CreateWorkerBootstrapTokenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateWorkerBootstrapTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkerBootstrapTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkerBootstrapTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkerBootstrapTokenResponseBody) SetCode(v string) *CreateWorkerBootstrapTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBody) SetData(v *CreateWorkerBootstrapTokenResponseBodyData) *CreateWorkerBootstrapTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBody) SetHttpStatusCode(v int32) *CreateWorkerBootstrapTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBody) SetMessage(v string) *CreateWorkerBootstrapTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBody) SetRequestId(v string) *CreateWorkerBootstrapTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBody) SetSuccess(v bool) *CreateWorkerBootstrapTokenResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkerBootstrapTokenResponseBodyData struct {
	BootstrapToken   *string                                        `json:"BootstrapToken,omitempty" xml:"BootstrapToken,omitempty"`
	Cms              *CreateWorkerBootstrapTokenResponseBodyDataCms `json:"Cms,omitempty" xml:"Cms,omitempty" type:"Struct"`
	InstanceId       *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name             *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkType      *string                                        `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	TokenFingerprint *string                                        `json:"TokenFingerprint,omitempty" xml:"TokenFingerprint,omitempty"`
}

func (s CreateWorkerBootstrapTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerBootstrapTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) GetBootstrapToken() *string {
	return s.BootstrapToken
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) GetCms() *CreateWorkerBootstrapTokenResponseBodyDataCms {
	return s.Cms
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) GetTokenFingerprint() *string {
	return s.TokenFingerprint
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) SetBootstrapToken(v string) *CreateWorkerBootstrapTokenResponseBodyData {
	s.BootstrapToken = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) SetCms(v *CreateWorkerBootstrapTokenResponseBodyDataCms) *CreateWorkerBootstrapTokenResponseBodyData {
	s.Cms = v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) SetInstanceId(v string) *CreateWorkerBootstrapTokenResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) SetName(v string) *CreateWorkerBootstrapTokenResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) SetNetworkType(v string) *CreateWorkerBootstrapTokenResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) SetTokenFingerprint(v string) *CreateWorkerBootstrapTokenResponseBodyData {
	s.TokenFingerprint = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBodyData) Validate() error {
	if s.Cms != nil {
		if err := s.Cms.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkerBootstrapTokenResponseBodyDataCms struct {
	Endpoint   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	Workspace  *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateWorkerBootstrapTokenResponseBodyDataCms) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerBootstrapTokenResponseBodyDataCms) GoString() string {
	return s.String()
}

func (s *CreateWorkerBootstrapTokenResponseBodyDataCms) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateWorkerBootstrapTokenResponseBodyDataCms) GetLicenseKey() *string {
	return s.LicenseKey
}

func (s *CreateWorkerBootstrapTokenResponseBodyDataCms) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateWorkerBootstrapTokenResponseBodyDataCms) SetEndpoint(v string) *CreateWorkerBootstrapTokenResponseBodyDataCms {
	s.Endpoint = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBodyDataCms) SetLicenseKey(v string) *CreateWorkerBootstrapTokenResponseBodyDataCms {
	s.LicenseKey = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBodyDataCms) SetWorkspace(v string) *CreateWorkerBootstrapTokenResponseBodyDataCms {
	s.Workspace = &v
	return s
}

func (s *CreateWorkerBootstrapTokenResponseBodyDataCms) Validate() error {
	return dara.Validate(s)
}
