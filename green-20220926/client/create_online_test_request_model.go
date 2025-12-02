// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOnlineTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v string) *CreateOnlineTestRequest
	GetDataId() *string
	SetResourceType(v string) *CreateOnlineTestRequest
	GetResourceType() *string
	SetServiceCode(v string) *CreateOnlineTestRequest
	GetServiceCode() *string
	SetUrl(v string) *CreateOnlineTestRequest
	GetUrl() *string
}

type CreateOnlineTestRequest struct {
	// Data ID
	//
	// example:
	//
	// xxxxxxx
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Resource Type
	//
	// example:
	//
	// video
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service Code
	//
	// example:
	//
	// VideoModeration
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Detection URL
	//
	// example:
	//
	// https://xxxxxxxxxx.com/data/data.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateOnlineTestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOnlineTestRequest) GoString() string {
	return s.String()
}

func (s *CreateOnlineTestRequest) GetDataId() *string {
	return s.DataId
}

func (s *CreateOnlineTestRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateOnlineTestRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *CreateOnlineTestRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateOnlineTestRequest) SetDataId(v string) *CreateOnlineTestRequest {
	s.DataId = &v
	return s
}

func (s *CreateOnlineTestRequest) SetResourceType(v string) *CreateOnlineTestRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateOnlineTestRequest) SetServiceCode(v string) *CreateOnlineTestRequest {
	s.ServiceCode = &v
	return s
}

func (s *CreateOnlineTestRequest) SetUrl(v string) *CreateOnlineTestRequest {
	s.Url = &v
	return s
}

func (s *CreateOnlineTestRequest) Validate() error {
	return dara.Validate(s)
}
