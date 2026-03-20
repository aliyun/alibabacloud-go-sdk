// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceInstanceTokenResponseBody
	GetRequestId() *string
	SetStreamlogUrl(v string) *CreateServiceInstanceTokenResponseBody
	GetStreamlogUrl() *string
	SetToken(v string) *CreateServiceInstanceTokenResponseBody
	GetToken() *string
	SetUrl(v string) *CreateServiceInstanceTokenResponseBody
	GetUrl() *string
	SetWorkbenchUrl(v string) *CreateServiceInstanceTokenResponseBody
	GetWorkbenchUrl() *string
}

type CreateServiceInstanceTokenResponseBody struct {
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ***
	StreamlogUrl *string `json:"StreamlogUrl,omitempty" xml:"StreamlogUrl,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// ***
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
	WorkbenchUrl *string `json:"WorkbenchUrl,omitempty" xml:"WorkbenchUrl,omitempty"`
}

func (s CreateServiceInstanceTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceInstanceTokenResponseBody) GetStreamlogUrl() *string {
	return s.StreamlogUrl
}

func (s *CreateServiceInstanceTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *CreateServiceInstanceTokenResponseBody) GetUrl() *string {
	return s.Url
}

func (s *CreateServiceInstanceTokenResponseBody) GetWorkbenchUrl() *string {
	return s.WorkbenchUrl
}

func (s *CreateServiceInstanceTokenResponseBody) SetRequestId(v string) *CreateServiceInstanceTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceInstanceTokenResponseBody) SetStreamlogUrl(v string) *CreateServiceInstanceTokenResponseBody {
	s.StreamlogUrl = &v
	return s
}

func (s *CreateServiceInstanceTokenResponseBody) SetToken(v string) *CreateServiceInstanceTokenResponseBody {
	s.Token = &v
	return s
}

func (s *CreateServiceInstanceTokenResponseBody) SetUrl(v string) *CreateServiceInstanceTokenResponseBody {
	s.Url = &v
	return s
}

func (s *CreateServiceInstanceTokenResponseBody) SetWorkbenchUrl(v string) *CreateServiceInstanceTokenResponseBody {
	s.WorkbenchUrl = &v
	return s
}

func (s *CreateServiceInstanceTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
