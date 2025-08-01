// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *RemoveApplicationRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *RemoveApplicationRequest
	GetAppId() *string
	SetAppName(v string) *RemoveApplicationRequest
	GetAppName() *string
	SetNamespace(v string) *RemoveApplicationRequest
	GetNamespace() *string
	SetRegion(v string) *RemoveApplicationRequest
	GetRegion() *string
}

type RemoveApplicationRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// abcde@12345
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// prod
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s RemoveApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationRequest) GoString() string {
	return s.String()
}

func (s *RemoveApplicationRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *RemoveApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RemoveApplicationRequest) GetAppName() *string {
	return s.AppName
}

func (s *RemoveApplicationRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *RemoveApplicationRequest) GetRegion() *string {
	return s.Region
}

func (s *RemoveApplicationRequest) SetAcceptLanguage(v string) *RemoveApplicationRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *RemoveApplicationRequest) SetAppId(v string) *RemoveApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RemoveApplicationRequest) SetAppName(v string) *RemoveApplicationRequest {
	s.AppName = &v
	return s
}

func (s *RemoveApplicationRequest) SetNamespace(v string) *RemoveApplicationRequest {
	s.Namespace = &v
	return s
}

func (s *RemoveApplicationRequest) SetRegion(v string) *RemoveApplicationRequest {
	s.Region = &v
	return s
}

func (s *RemoveApplicationRequest) Validate() error {
	return dara.Validate(s)
}
