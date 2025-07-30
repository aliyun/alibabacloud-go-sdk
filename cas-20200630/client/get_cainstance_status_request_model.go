// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCAInstanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *GetCAInstanceStatusRequest
	GetIdentifier() *string
	SetInstanceId(v string) *GetCAInstanceStatusRequest
	GetInstanceId() *string
}

type GetCAInstanceStatusRequest struct {
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the private CA instance.
	//
	// >  After you purchase a private CA instance by using the [SSL Certificates Service console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist), you can click **Details*	- for the private CA instance on the **Private Certificates*	- page to query the ID of the private CA instance.
	//
	// example:
	//
	// cas-member-0hmi****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCAInstanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCAInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCAInstanceStatusRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetCAInstanceStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCAInstanceStatusRequest) SetIdentifier(v string) *GetCAInstanceStatusRequest {
	s.Identifier = &v
	return s
}

func (s *GetCAInstanceStatusRequest) SetInstanceId(v string) *GetCAInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCAInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
