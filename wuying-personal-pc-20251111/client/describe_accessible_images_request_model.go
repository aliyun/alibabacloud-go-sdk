// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessibleImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribeAccessibleImagesRequest
	GetApiKey() *string
	SetBizSource(v string) *DescribeAccessibleImagesRequest
	GetBizSource() *string
	SetDesktopId(v string) *DescribeAccessibleImagesRequest
	GetDesktopId() *string
	SetDesktopType(v string) *DescribeAccessibleImagesRequest
	GetDesktopType() *string
	SetImageType(v string) *DescribeAccessibleImagesRequest
	GetImageType() *string
	SetResourceId(v string) *DescribeAccessibleImagesRequest
	GetResourceId() *string
	SetScene(v string) *DescribeAccessibleImagesRequest
	GetScene() *string
	SetSearchKey(v string) *DescribeAccessibleImagesRequest
	GetSearchKey() *string
	SetSessionId(v string) *DescribeAccessibleImagesRequest
	GetSessionId() *string
}

type DescribeAccessibleImagesRequest struct {
	// This parameter is required.
	ApiKey      *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	BizSource   *string `json:"BizSource,omitempty" xml:"BizSource,omitempty"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	ImageType   *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Scene       *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SearchKey   *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	SessionId   *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribeAccessibleImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessibleImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessibleImagesRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeAccessibleImagesRequest) GetBizSource() *string {
	return s.BizSource
}

func (s *DescribeAccessibleImagesRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeAccessibleImagesRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeAccessibleImagesRequest) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeAccessibleImagesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeAccessibleImagesRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeAccessibleImagesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeAccessibleImagesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeAccessibleImagesRequest) SetApiKey(v string) *DescribeAccessibleImagesRequest {
	s.ApiKey = &v
	return s
}

func (s *DescribeAccessibleImagesRequest) SetBizSource(v string) *DescribeAccessibleImagesRequest {
	s.BizSource = &v
	return s
}

func (s *DescribeAccessibleImagesRequest) SetDesktopId(v string) *DescribeAccessibleImagesRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeAccessibleImagesRequest) SetDesktopType(v string) *DescribeAccessibleImagesRequest {
	s.DesktopType = &v
	return s
}

func (s *DescribeAccessibleImagesRequest) SetImageType(v string) *DescribeAccessibleImagesRequest {
	s.ImageType = &v
	return s
}

func (s *DescribeAccessibleImagesRequest) SetResourceId(v string) *DescribeAccessibleImagesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeAccessibleImagesRequest) SetScene(v string) *DescribeAccessibleImagesRequest {
	s.Scene = &v
	return s
}

func (s *DescribeAccessibleImagesRequest) SetSearchKey(v string) *DescribeAccessibleImagesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAccessibleImagesRequest) SetSessionId(v string) *DescribeAccessibleImagesRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeAccessibleImagesRequest) Validate() error {
	return dara.Validate(s)
}
