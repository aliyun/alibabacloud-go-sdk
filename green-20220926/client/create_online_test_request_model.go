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
	// The data ID.
	//
	// example:
	//
	// 4f27b8cc7c4544cb90b41882a5b36326
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The resource type. Valid values: voice (audio), image (image), file (document), text (text), and video (video).
	//
	// example:
	//
	// video
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service code. Specify a valid service code. The valid service code for online video testing is videoDetection. VideoModeration is not a valid value and returns a 406 online.test.failed error.
	//
	// example:
	//
	// videoDetection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The detection URL.
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
