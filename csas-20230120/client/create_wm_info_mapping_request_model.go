// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmInfoMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWmInfoBytesB64(v string) *CreateWmInfoMappingRequest
	GetWmInfoBytesB64() *string
	SetWmInfoSize(v int64) *CreateWmInfoMappingRequest
	GetWmInfoSize() *int64
	SetWmType(v string) *CreateWmInfoMappingRequest
	GetWmType() *string
}

type CreateWmInfoMappingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmInfoMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWmInfoMappingRequest) GoString() string {
	return s.String()
}

func (s *CreateWmInfoMappingRequest) GetWmInfoBytesB64() *string {
	return s.WmInfoBytesB64
}

func (s *CreateWmInfoMappingRequest) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *CreateWmInfoMappingRequest) GetWmType() *string {
	return s.WmType
}

func (s *CreateWmInfoMappingRequest) SetWmInfoBytesB64(v string) *CreateWmInfoMappingRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmInfoMappingRequest) SetWmInfoSize(v int64) *CreateWmInfoMappingRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmInfoMappingRequest) SetWmType(v string) *CreateWmInfoMappingRequest {
	s.WmType = &v
	return s
}

func (s *CreateWmInfoMappingRequest) Validate() error {
	return dara.Validate(s)
}
