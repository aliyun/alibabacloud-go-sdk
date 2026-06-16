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
	RegionId   *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	FromEcs *bool   `json:"FromEcs,omitempty" xml:"FromEcs,omitempty"`
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
