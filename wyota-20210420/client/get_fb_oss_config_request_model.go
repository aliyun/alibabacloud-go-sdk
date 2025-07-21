// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFbOssConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaSite(v string) *GetFbOssConfigRequest
	GetAreaSite() *string
	SetDirPrefix(v string) *GetFbOssConfigRequest
	GetDirPrefix() *string
	SetIsDedicatedLine(v int32) *GetFbOssConfigRequest
	GetIsDedicatedLine() *int32
	SetRegion(v string) *GetFbOssConfigRequest
	GetRegion() *string
}

type GetFbOssConfigRequest struct {
	AreaSite        *string `json:"AreaSite,omitempty" xml:"AreaSite,omitempty"`
	DirPrefix       *string `json:"DirPrefix,omitempty" xml:"DirPrefix,omitempty"`
	IsDedicatedLine *int32  `json:"IsDedicatedLine,omitempty" xml:"IsDedicatedLine,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetFbOssConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFbOssConfigRequest) GoString() string {
	return s.String()
}

func (s *GetFbOssConfigRequest) GetAreaSite() *string {
	return s.AreaSite
}

func (s *GetFbOssConfigRequest) GetDirPrefix() *string {
	return s.DirPrefix
}

func (s *GetFbOssConfigRequest) GetIsDedicatedLine() *int32 {
	return s.IsDedicatedLine
}

func (s *GetFbOssConfigRequest) GetRegion() *string {
	return s.Region
}

func (s *GetFbOssConfigRequest) SetAreaSite(v string) *GetFbOssConfigRequest {
	s.AreaSite = &v
	return s
}

func (s *GetFbOssConfigRequest) SetDirPrefix(v string) *GetFbOssConfigRequest {
	s.DirPrefix = &v
	return s
}

func (s *GetFbOssConfigRequest) SetIsDedicatedLine(v int32) *GetFbOssConfigRequest {
	s.IsDedicatedLine = &v
	return s
}

func (s *GetFbOssConfigRequest) SetRegion(v string) *GetFbOssConfigRequest {
	s.Region = &v
	return s
}

func (s *GetFbOssConfigRequest) Validate() error {
	return dara.Validate(s)
}
