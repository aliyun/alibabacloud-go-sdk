// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryRoomStatusResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryRoomStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRoomStatusResponseBody
	GetRequestId() *string
	SetResult(v *QueryRoomStatusResponseBodyResult) *QueryRoomStatusResponseBody
	GetResult() *QueryRoomStatusResponseBodyResult
	SetStatusCode(v int32) *QueryRoomStatusResponseBody
	GetStatusCode() *int32
}

type QueryRoomStatusResponseBody struct {
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
	// FAFCD152-4791-5F2F-B0BE-2DC06FD4F05B
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryRoomStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryRoomStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRoomStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryRoomStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRoomStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRoomStatusResponseBody) GetResult() *QueryRoomStatusResponseBodyResult {
	return s.Result
}

func (s *QueryRoomStatusResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRoomStatusResponseBody) SetCode(v int32) *QueryRoomStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRoomStatusResponseBody) SetMessage(v string) *QueryRoomStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRoomStatusResponseBody) SetRequestId(v string) *QueryRoomStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRoomStatusResponseBody) SetResult(v *QueryRoomStatusResponseBodyResult) *QueryRoomStatusResponseBody {
	s.Result = v
	return s
}

func (s *QueryRoomStatusResponseBody) SetStatusCode(v int32) *QueryRoomStatusResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryRoomStatusResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRoomStatusResponseBodyResult struct {
	// example:
	//
	// cf2446fc9d144c85aaee4f9ae20a96e7
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 1211
	RoomNo     *string                                        `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	SceneList  []*QueryRoomStatusResponseBodyResultSceneList  `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Repeated"`
	StatusList []*QueryRoomStatusResponseBodyResultStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s QueryRoomStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryRoomStatusResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *QueryRoomStatusResponseBodyResult) GetRoomNo() *string {
	return s.RoomNo
}

func (s *QueryRoomStatusResponseBodyResult) GetSceneList() []*QueryRoomStatusResponseBodyResultSceneList {
	return s.SceneList
}

func (s *QueryRoomStatusResponseBodyResult) GetStatusList() []*QueryRoomStatusResponseBodyResultStatusList {
	return s.StatusList
}

func (s *QueryRoomStatusResponseBodyResult) SetHotelId(v string) *QueryRoomStatusResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *QueryRoomStatusResponseBodyResult) SetRoomNo(v string) *QueryRoomStatusResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *QueryRoomStatusResponseBodyResult) SetSceneList(v []*QueryRoomStatusResponseBodyResultSceneList) *QueryRoomStatusResponseBodyResult {
	s.SceneList = v
	return s
}

func (s *QueryRoomStatusResponseBodyResult) SetStatusList(v []*QueryRoomStatusResponseBodyResultStatusList) *QueryRoomStatusResponseBodyResult {
	s.StatusList = v
	return s
}

func (s *QueryRoomStatusResponseBodyResult) Validate() error {
	if s.SceneList != nil {
		for _, item := range s.SceneList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StatusList != nil {
		for _, item := range s.StatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRoomStatusResponseBodyResultSceneList struct {
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s QueryRoomStatusResponseBodyResultSceneList) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomStatusResponseBodyResultSceneList) GoString() string {
	return s.String()
}

func (s *QueryRoomStatusResponseBodyResultSceneList) GetSceneName() *string {
	return s.SceneName
}

func (s *QueryRoomStatusResponseBodyResultSceneList) SetSceneName(v string) *QueryRoomStatusResponseBodyResultSceneList {
	s.SceneName = &v
	return s
}

func (s *QueryRoomStatusResponseBodyResultSceneList) Validate() error {
	return dara.Validate(s)
}

type QueryRoomStatusResponseBodyResultStatusList struct {
	StatusName  *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	StatusValue *string `json:"StatusValue,omitempty" xml:"StatusValue,omitempty"`
	// example:
	//
	// Thu Jan 09 13:56:51 CST 2025
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryRoomStatusResponseBodyResultStatusList) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomStatusResponseBodyResultStatusList) GoString() string {
	return s.String()
}

func (s *QueryRoomStatusResponseBodyResultStatusList) GetStatusName() *string {
	return s.StatusName
}

func (s *QueryRoomStatusResponseBodyResultStatusList) GetStatusValue() *string {
	return s.StatusValue
}

func (s *QueryRoomStatusResponseBodyResultStatusList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryRoomStatusResponseBodyResultStatusList) SetStatusName(v string) *QueryRoomStatusResponseBodyResultStatusList {
	s.StatusName = &v
	return s
}

func (s *QueryRoomStatusResponseBodyResultStatusList) SetStatusValue(v string) *QueryRoomStatusResponseBodyResultStatusList {
	s.StatusValue = &v
	return s
}

func (s *QueryRoomStatusResponseBodyResultStatusList) SetUpdateTime(v string) *QueryRoomStatusResponseBodyResultStatusList {
	s.UpdateTime = &v
	return s
}

func (s *QueryRoomStatusResponseBodyResultStatusList) Validate() error {
	return dara.Validate(s)
}
