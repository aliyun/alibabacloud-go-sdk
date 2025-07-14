// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UnbindSlbRequest
	GetAppId() *string
	SetInternet(v bool) *UnbindSlbRequest
	GetInternet() *bool
	SetIntranet(v bool) *UnbindSlbRequest
	GetIntranet() *bool
}

type UnbindSlbRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to disassociate the Internet-facing SLB instance. Valid values:
	//
	// 	- **true**: dissociates the Internet-facing SLB instance.
	//
	// 	- **false**: does not dissociate the Internet-facing SLB instance.
	//
	// example:
	//
	// true
	Internet *bool `json:"Internet,omitempty" xml:"Internet,omitempty"`
	// Specifies whether to disassociate the internal-facing SLB instance. Valid values:
	//
	// 	- **true**: dissociates the internal-facing SLB instance.
	//
	// 	- **false**: does not dissociate the internal-facing SLB instance.
	//
	// example:
	//
	// true
	Intranet *bool `json:"Intranet,omitempty" xml:"Intranet,omitempty"`
}

func (s UnbindSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindSlbRequest) GoString() string {
	return s.String()
}

func (s *UnbindSlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *UnbindSlbRequest) GetInternet() *bool {
	return s.Internet
}

func (s *UnbindSlbRequest) GetIntranet() *bool {
	return s.Intranet
}

func (s *UnbindSlbRequest) SetAppId(v string) *UnbindSlbRequest {
	s.AppId = &v
	return s
}

func (s *UnbindSlbRequest) SetInternet(v bool) *UnbindSlbRequest {
	s.Internet = &v
	return s
}

func (s *UnbindSlbRequest) SetIntranet(v bool) *UnbindSlbRequest {
	s.Intranet = &v
	return s
}

func (s *UnbindSlbRequest) Validate() error {
	return dara.Validate(s)
}
