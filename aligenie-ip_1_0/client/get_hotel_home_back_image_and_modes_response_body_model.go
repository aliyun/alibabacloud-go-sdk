// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelHomeBackImageAndModesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetHotelHomeBackImageAndModesResponseBody
	GetCode() *int32
	SetMessage(v string) *GetHotelHomeBackImageAndModesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelHomeBackImageAndModesResponseBody
	GetRequestId() *string
	SetResult(v *GetHotelHomeBackImageAndModesResponseBodyResult) *GetHotelHomeBackImageAndModesResponseBody
	GetResult() *GetHotelHomeBackImageAndModesResponseBodyResult
}

type GetHotelHomeBackImageAndModesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 394450FC-9035-1B7C-8829-BC88832473FC
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelHomeBackImageAndModesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetHotelHomeBackImageAndModesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetHotelHomeBackImageAndModesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelHomeBackImageAndModesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelHomeBackImageAndModesResponseBody) GetResult() *GetHotelHomeBackImageAndModesResponseBodyResult {
	return s.Result
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetCode(v int32) *GetHotelHomeBackImageAndModesResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetMessage(v string) *GetHotelHomeBackImageAndModesResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetRequestId(v string) *GetHotelHomeBackImageAndModesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) SetResult(v *GetHotelHomeBackImageAndModesResponseBodyResult) *GetHotelHomeBackImageAndModesResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelHomeBackImageAndModesResponseBodyResult struct {
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/3d4fe6d66ec49d9789635f66627f0339/welcome_audios/976210a6532150f49c2677a8b7dbc105/l6fspbhn.jpg
	BackgroundImage *string `json:"BackgroundImage,omitempty" xml:"BackgroundImage,omitempty"`
	// example:
	//
	// 宣雍测试橙蜂酒店
	HotelName *string                                                    `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	ModeList  []*GetHotelHomeBackImageAndModesResponseBodyResultModeList `json:"ModeList,omitempty" xml:"ModeList,omitempty" type:"Repeated"`
}

func (s GetHotelHomeBackImageAndModesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) GetBackgroundImage() *string {
	return s.BackgroundImage
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) GetHotelName() *string {
	return s.HotelName
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) GetModeList() []*GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	return s.ModeList
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) SetBackgroundImage(v string) *GetHotelHomeBackImageAndModesResponseBodyResult {
	s.BackgroundImage = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) SetHotelName(v string) *GetHotelHomeBackImageAndModesResponseBodyResult {
	s.HotelName = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) SetModeList(v []*GetHotelHomeBackImageAndModesResponseBodyResultModeList) *GetHotelHomeBackImageAndModesResponseBodyResult {
	s.ModeList = v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResult) Validate() error {
	if s.ModeList != nil {
		for _, item := range s.ModeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotelHomeBackImageAndModesResponseBodyResultModeList struct {
	// example:
	//
	// 浪漫模式
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	// example:
	//
	// romantic
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingmoshi/langman.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
}

func (s GetHotelHomeBackImageAndModesResponseBodyResultModeList) String() string {
	return dara.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponseBodyResultModeList) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) GetCnName() *string {
	return s.CnName
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) GetCode() *string {
	return s.Code
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) GetIcon() *string {
	return s.Icon
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) SetCnName(v string) *GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	s.CnName = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) SetCode(v string) *GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	s.Code = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) SetIcon(v string) *GetHotelHomeBackImageAndModesResponseBodyResultModeList {
	s.Icon = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponseBodyResultModeList) Validate() error {
	return dara.Validate(s)
}
