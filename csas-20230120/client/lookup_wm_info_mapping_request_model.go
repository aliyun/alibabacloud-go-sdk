// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupWmInfoMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWmInfoSize(v int64) *LookupWmInfoMappingRequest
	GetWmInfoSize() *int64
	SetWmInfoUint(v string) *LookupWmInfoMappingRequest
	GetWmInfoUint() *string
	SetWmType(v string) *LookupWmInfoMappingRequest
	GetWmType() *string
}

type LookupWmInfoMappingRequest struct {
	// Bit width of the watermark information. Default value: 32. This parameter must match the bit width used when embedding or generating a transparent image. Valid values: 32 to 64. Use the same value as when you created the mapping. Otherwise, the mapping cannot be found.
	//
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// Numeric-formatted watermark information. Value source:
	//
	// - [CreateWmInfoMapping](~~CreateWmInfoMapping~~): The **WmInfoUint*	- return value from the CreateWmInfoMapping API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// Watermark type. Valid values:
	//
	// - **PureWebappInvisible**: Webpage watermark.
	//
	// - **PureAppInvisible**: App watermark.
	//
	// - **PureScreenInvisible**: Screen watermark.
	//
	// - **PureDocument**: Document watermark.
	//
	// - **PureImage**: Image watermark.
	//
	// - **PureAudio**: Audio watermark.
	//
	// - **PureVideo**: Video watermark.
	//
	// - **AigcWebappInvisible**: AIGC webpage watermark.
	//
	// - **AigcAppInvisible**: AIGC app watermark.
	//
	// - **AigcScreenInvisible**: AIGC screen watermark.
	//
	// - **AigcDocument**: AIGC document watermark.
	//
	// - **AigcImage**: AIGC image watermark.
	//
	// - **AigcAudio**: AIGC audio watermark.
	//
	// - **AigcVideo**: AIGC video watermark.
	//
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s LookupWmInfoMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s LookupWmInfoMappingRequest) GoString() string {
	return s.String()
}

func (s *LookupWmInfoMappingRequest) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *LookupWmInfoMappingRequest) GetWmInfoUint() *string {
	return s.WmInfoUint
}

func (s *LookupWmInfoMappingRequest) GetWmType() *string {
	return s.WmType
}

func (s *LookupWmInfoMappingRequest) SetWmInfoSize(v int64) *LookupWmInfoMappingRequest {
	s.WmInfoSize = &v
	return s
}

func (s *LookupWmInfoMappingRequest) SetWmInfoUint(v string) *LookupWmInfoMappingRequest {
	s.WmInfoUint = &v
	return s
}

func (s *LookupWmInfoMappingRequest) SetWmType(v string) *LookupWmInfoMappingRequest {
	s.WmType = &v
	return s
}

func (s *LookupWmInfoMappingRequest) Validate() error {
	return dara.Validate(s)
}
