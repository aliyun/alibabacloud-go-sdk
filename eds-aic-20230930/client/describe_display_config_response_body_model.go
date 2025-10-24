// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisplayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayConfigModel(v []*DescribeDisplayConfigResponseBodyDisplayConfigModel) *DescribeDisplayConfigResponseBody
	GetDisplayConfigModel() []*DescribeDisplayConfigResponseBodyDisplayConfigModel
	SetRequestId(v string) *DescribeDisplayConfigResponseBody
	GetRequestId() *string
}

type DescribeDisplayConfigResponseBody struct {
	DisplayConfigModel []*DescribeDisplayConfigResponseBodyDisplayConfigModel `json:"DisplayConfigModel,omitempty" xml:"DisplayConfigModel,omitempty" type:"Repeated"`
	// example:
	//
	// FFEF7EFE-1E36-56D1-B5BF-5BACE43B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDisplayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisplayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisplayConfigResponseBody) GetDisplayConfigModel() []*DescribeDisplayConfigResponseBodyDisplayConfigModel {
	return s.DisplayConfigModel
}

func (s *DescribeDisplayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDisplayConfigResponseBody) SetDisplayConfigModel(v []*DescribeDisplayConfigResponseBodyDisplayConfigModel) *DescribeDisplayConfigResponseBody {
	s.DisplayConfigModel = v
	return s
}

func (s *DescribeDisplayConfigResponseBody) SetRequestId(v string) *DescribeDisplayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisplayConfigResponseBody) Validate() error {
	if s.DisplayConfigModel != nil {
		for _, item := range s.DisplayConfigModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDisplayConfigResponseBodyDisplayConfigModel struct {
	// example:
	//
	// cpn-jewjt8xryuituz4qn-****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// example:
	//
	// 240
	Dpi *int32 `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	// example:
	//
	// null
	Fps *int32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s DescribeDisplayConfigResponseBodyDisplayConfigModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisplayConfigResponseBodyDisplayConfigModel) GoString() string {
	return s.String()
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) GetDpi() *int32 {
	return s.Dpi
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) GetFps() *int32 {
	return s.Fps
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) GetLockResolution() *string {
	return s.LockResolution
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetAndroidInstanceId(v string) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetDpi(v int32) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.Dpi = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetFps(v int32) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.Fps = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetLockResolution(v string) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.LockResolution = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetResolutionHeight(v int32) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) SetResolutionWidth(v int32) *DescribeDisplayConfigResponseBodyDisplayConfigModel {
	s.ResolutionWidth = &v
	return s
}

func (s *DescribeDisplayConfigResponseBodyDisplayConfigModel) Validate() error {
	return dara.Validate(s)
}
