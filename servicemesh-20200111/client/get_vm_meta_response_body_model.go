// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVmMetaResponseBody
	GetRequestId() *string
	SetVmMetaInfo(v *GetVmMetaResponseBodyVmMetaInfo) *GetVmMetaResponseBody
	GetVmMetaInfo() *GetVmMetaResponseBodyVmMetaInfo
}

type GetVmMetaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9522f7c9-63a1-4603-b850-37d12a****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata that is required to add a non-containerized application to the ASM instance.
	VmMetaInfo *GetVmMetaResponseBodyVmMetaInfo `json:"VmMetaInfo,omitempty" xml:"VmMetaInfo,omitempty" type:"Struct"`
}

func (s GetVmMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVmMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetVmMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVmMetaResponseBody) GetVmMetaInfo() *GetVmMetaResponseBodyVmMetaInfo {
	return s.VmMetaInfo
}

func (s *GetVmMetaResponseBody) SetRequestId(v string) *GetVmMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVmMetaResponseBody) SetVmMetaInfo(v *GetVmMetaResponseBodyVmMetaInfo) *GetVmMetaResponseBody {
	s.VmMetaInfo = v
	return s
}

func (s *GetVmMetaResponseBody) Validate() error {
	if s.VmMetaInfo != nil {
		if err := s.VmMetaInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVmMetaResponseBodyVmMetaInfo struct {
	// The content of the EnvoyEnv file.
	//
	// example:
	//
	// ....
	EnvoyEnvContent *string `json:"EnvoyEnvContent,omitempty" xml:"EnvoyEnvContent,omitempty"`
	// The content of the hosts file.
	//
	// example:
	//
	// ....
	HostsContent *string `json:"HostsContent,omitempty" xml:"HostsContent,omitempty"`
	// The content of the Token file.
	//
	// example:
	//
	// ....
	TokenContent *string `json:"TokenContent,omitempty" xml:"TokenContent,omitempty"`
}

func (s GetVmMetaResponseBodyVmMetaInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVmMetaResponseBodyVmMetaInfo) GoString() string {
	return s.String()
}

func (s *GetVmMetaResponseBodyVmMetaInfo) GetEnvoyEnvContent() *string {
	return s.EnvoyEnvContent
}

func (s *GetVmMetaResponseBodyVmMetaInfo) GetHostsContent() *string {
	return s.HostsContent
}

func (s *GetVmMetaResponseBodyVmMetaInfo) GetTokenContent() *string {
	return s.TokenContent
}

func (s *GetVmMetaResponseBodyVmMetaInfo) SetEnvoyEnvContent(v string) *GetVmMetaResponseBodyVmMetaInfo {
	s.EnvoyEnvContent = &v
	return s
}

func (s *GetVmMetaResponseBodyVmMetaInfo) SetHostsContent(v string) *GetVmMetaResponseBodyVmMetaInfo {
	s.HostsContent = &v
	return s
}

func (s *GetVmMetaResponseBodyVmMetaInfo) SetTokenContent(v string) *GetVmMetaResponseBodyVmMetaInfo {
	s.TokenContent = &v
	return s
}

func (s *GetVmMetaResponseBodyVmMetaInfo) Validate() error {
	return dara.Validate(s)
}
