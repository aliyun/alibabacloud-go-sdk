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
	// The unique identifier of the client certificate or server-side certificate to query.
	//
	// > Call [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) to query the unique identifiers of all client certificates and server-side certificates.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the private CA instance to query.
	//
	// > After you purchase a private CA instance in the [CAS console](https://yundun.console.aliyun.com/?p=cas#/pca/rootlist), you can go to the **Private Certificates*	- page and view the **details*	- of the instance to obtain its ID.
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
