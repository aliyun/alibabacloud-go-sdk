// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindK8sSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UnbindK8sSlbRequest
	GetAppId() *string
	SetClusterId(v string) *UnbindK8sSlbRequest
	GetClusterId() *string
	SetSlbName(v string) *UnbindK8sSlbRequest
	GetSlbName() *string
	SetType(v string) *UnbindK8sSlbRequest
	GetType() *string
}

type UnbindK8sSlbRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd-****-****-a286-781659d9f54c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the cluster. You can call the GetK8sCluster operation to query the cluster ID. For more information, see [GetK8sCluster](https://help.aliyun.com/document_detail/181437.html).
	//
	// example:
	//
	// 712082c3-****-****-9217-a947b5cde6ee
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the SLB instance.
	//
	// example:
	//
	// a2d4ab12666544a479cdd0711a97****
	SlbName *string `json:"SlbName,omitempty" xml:"SlbName,omitempty"`
	// The type of the SLB instance. Valid values:
	//
	// 	- **internet**: Internet-facing SLB instance
	//
	// 	- **intranet**: internal-facing SLB instance
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UnbindK8sSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindK8sSlbRequest) GoString() string {
	return s.String()
}

func (s *UnbindK8sSlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *UnbindK8sSlbRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UnbindK8sSlbRequest) GetSlbName() *string {
	return s.SlbName
}

func (s *UnbindK8sSlbRequest) GetType() *string {
	return s.Type
}

func (s *UnbindK8sSlbRequest) SetAppId(v string) *UnbindK8sSlbRequest {
	s.AppId = &v
	return s
}

func (s *UnbindK8sSlbRequest) SetClusterId(v string) *UnbindK8sSlbRequest {
	s.ClusterId = &v
	return s
}

func (s *UnbindK8sSlbRequest) SetSlbName(v string) *UnbindK8sSlbRequest {
	s.SlbName = &v
	return s
}

func (s *UnbindK8sSlbRequest) SetType(v string) *UnbindK8sSlbRequest {
	s.Type = &v
	return s
}

func (s *UnbindK8sSlbRequest) Validate() error {
	return dara.Validate(s)
}
