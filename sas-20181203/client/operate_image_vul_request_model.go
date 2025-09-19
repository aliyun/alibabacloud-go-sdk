// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateImageVulRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInfo(v string) *OperateImageVulRequest
	GetInfo() *string
	SetOperateType(v string) *OperateImageVulRequest
	GetOperateType() *string
	SetType(v string) *OperateImageVulRequest
	GetType() *string
}

type OperateImageVulRequest struct {
	// The information about the vulnerability. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- namespace: the namespace of the image
	//
	// 	- repoName: the name of the Container Registry repository
	//
	// 	- regionId: the region ID
	//
	// 	- instanceId: the ID of the Container Registry instance
	//
	// 	- repoId: the name of the repository
	//
	// 	- tag: the tad added to the image
	//
	// 	- digest: the digest of the image
	//
	// 	- newTag: the tag added to the image after the vulnerability is fixed
	//
	// 	- uuid: the UUID of the image
	//
	// 	- ids: the IDs of the vulnerability primary keys
	//
	// example:
	//
	// [{\\"namespace\\":\\"cloud_oa****\\",\\"repoName\\":\\"hybirdc****\\",\\"regionId\\":\\"cn-shanghai\\",\\"instanceId\\":\\"cri-rv4nvbv8iju4****\\",\\"repoId\\":\\"crr-2q7302qrofxg****\\",\\"tag\\":\\"hybird-cloud-web_fix_167115945****\\",\\"digest\\":\\"e1a4fd25884ca2ef8840bb252c9926e4f549df9e046500dd93539b2d458c****\\",\\"newTag\\":\\"hybird-cloud-web_fix_167115996****\\",\\"uuid\\":\\"4ad91dd8c0c02de6574fa98085d0****\\",\\"ids\\":[197540864,197540865,197540869]}]
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// If you want to fix the vulnerability, set the value to vul_fix.
	//
	// example:
	//
	// vul_fix
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The type of the vulnerability. Set the value to cve.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OperateImageVulRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateImageVulRequest) GoString() string {
	return s.String()
}

func (s *OperateImageVulRequest) GetInfo() *string {
	return s.Info
}

func (s *OperateImageVulRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateImageVulRequest) GetType() *string {
	return s.Type
}

func (s *OperateImageVulRequest) SetInfo(v string) *OperateImageVulRequest {
	s.Info = &v
	return s
}

func (s *OperateImageVulRequest) SetOperateType(v string) *OperateImageVulRequest {
	s.OperateType = &v
	return s
}

func (s *OperateImageVulRequest) SetType(v string) *OperateImageVulRequest {
	s.Type = &v
	return s
}

func (s *OperateImageVulRequest) Validate() error {
	return dara.Validate(s)
}
