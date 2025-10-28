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
	SetDeleteListener(v string) *UnbindSlbRequest
	GetDeleteListener() *string
	SetSlbId(v string) *UnbindSlbRequest
	GetSlbId() *string
	SetType(v string) *UnbindSlbRequest
	GetType() *string
}

type UnbindSlbRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// c627c157-560d-********************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to remove the configured listeners. Valid values:
	//
	// 	- true: removes the configured listeners.
	//
	// 	- false: does not remove the configured listeners.
	//
	// example:
	//
	// false
	DeleteListener *string `json:"DeleteListener,omitempty" xml:"DeleteListener,omitempty"`
	// The ID of the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-wz9vo49****************
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The network type of the SLB instance. Valid values:
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

func (s UnbindSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindSlbRequest) GoString() string {
	return s.String()
}

func (s *UnbindSlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *UnbindSlbRequest) GetDeleteListener() *string {
	return s.DeleteListener
}

func (s *UnbindSlbRequest) GetSlbId() *string {
	return s.SlbId
}

func (s *UnbindSlbRequest) GetType() *string {
	return s.Type
}

func (s *UnbindSlbRequest) SetAppId(v string) *UnbindSlbRequest {
	s.AppId = &v
	return s
}

func (s *UnbindSlbRequest) SetDeleteListener(v string) *UnbindSlbRequest {
	s.DeleteListener = &v
	return s
}

func (s *UnbindSlbRequest) SetSlbId(v string) *UnbindSlbRequest {
	s.SlbId = &v
	return s
}

func (s *UnbindSlbRequest) SetType(v string) *UnbindSlbRequest {
	s.Type = &v
	return s
}

func (s *UnbindSlbRequest) Validate() error {
	return dara.Validate(s)
}
