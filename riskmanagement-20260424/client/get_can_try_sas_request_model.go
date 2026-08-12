// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCanTrySasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetCanTrySasRequest
	GetRegionId() *string
	SetSdkRequest(v *GetCanTrySasRequestSdkRequest) *GetCanTrySasRequest
	GetSdkRequest() *GetCanTrySasRequestSdkRequest
}

type GetCanTrySasRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-guangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request parameters.
	SdkRequest *GetCanTrySasRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s GetCanTrySasRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasRequest) GoString() string {
	return s.String()
}

func (s *GetCanTrySasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCanTrySasRequest) GetSdkRequest() *GetCanTrySasRequestSdkRequest {
	return s.SdkRequest
}

func (s *GetCanTrySasRequest) SetRegionId(v string) *GetCanTrySasRequest {
	s.RegionId = &v
	return s
}

func (s *GetCanTrySasRequest) SetSdkRequest(v *GetCanTrySasRequestSdkRequest) *GetCanTrySasRequest {
	s.SdkRequest = v
	return s
}

func (s *GetCanTrySasRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCanTrySasRequestSdkRequest struct {
	// Specifies whether the request is from the ECS console. Valid values:
	//
	// - **true**: The request is from the ECS console.
	//
	// - **false**: The request is not from the ECS console.
	//
	// example:
	//
	// true
	FromEcs *bool `json:"FromEcs,omitempty" xml:"FromEcs,omitempty"`
	// The language type for the request and response messages. Default value: zh. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetCanTrySasRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *GetCanTrySasRequestSdkRequest) GetFromEcs() *bool {
	return s.FromEcs
}

func (s *GetCanTrySasRequestSdkRequest) GetLang() *string {
	return s.Lang
}

func (s *GetCanTrySasRequestSdkRequest) SetFromEcs(v bool) *GetCanTrySasRequestSdkRequest {
	s.FromEcs = &v
	return s
}

func (s *GetCanTrySasRequestSdkRequest) SetLang(v string) *GetCanTrySasRequestSdkRequest {
	s.Lang = &v
	return s
}

func (s *GetCanTrySasRequestSdkRequest) Validate() error {
	return dara.Validate(s)
}
