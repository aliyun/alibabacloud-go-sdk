// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetApplicationRequest
	GetAppId() *string
	SetAppName(v string) *GetApplicationRequest
	GetAppName() *string
	SetNamespaceId(v string) *GetApplicationRequest
	GetNamespaceId() *string
}

type GetApplicationRequest struct {
	// The application ID.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-shenzhen
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApplicationRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetApplicationRequest) SetAppId(v string) *GetApplicationRequest {
	s.AppId = &v
	return s
}

func (s *GetApplicationRequest) SetAppName(v string) *GetApplicationRequest {
	s.AppName = &v
	return s
}

func (s *GetApplicationRequest) SetNamespaceId(v string) *GetApplicationRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetApplicationRequest) Validate() error {
	return dara.Validate(s)
}
