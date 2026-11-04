// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceUserResponseBody
	GetCode() *string
	SetData(v *QueryFaceUserResponseBodyData) *QueryFaceUserResponseBody
	GetData() *QueryFaceUserResponseBodyData
	SetErrorMessage(v string) *QueryFaceUserResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceUserResponseBody
	GetSuccess() *bool
}

type QueryFaceUserResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryFaceUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A13279A-5640-45E7-87AA-83750541AD0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceUserResponseBody) GetData() *QueryFaceUserResponseBodyData {
	return s.Data
}

func (s *QueryFaceUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceUserResponseBody) SetCode(v string) *QueryFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserResponseBody) SetData(v *QueryFaceUserResponseBodyData) *QueryFaceUserResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserResponseBody) SetErrorMessage(v string) *QueryFaceUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserResponseBody) SetRequestId(v string) *QueryFaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserResponseBody) SetSuccess(v bool) *QueryFaceUserResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceUserResponseBodyData struct {
	// The user ID in the caller\\"s system.
	//
	// example:
	//
	// TestCustomUserId
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// The list of face photo information.
	FacePicList []*QueryFaceUserResponseBodyDataFacePicList `json:"FacePicList,omitempty" xml:"FacePicList,omitempty" type:"Repeated"`
	// The username.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user parameters.
	//
	// example:
	//
	// {"age":28, "sex":"male"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The user ID.
	//
	// example:
	//
	// TestUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBodyData) GetCustomUserId() *string {
	return s.CustomUserId
}

func (s *QueryFaceUserResponseBodyData) GetFacePicList() []*QueryFaceUserResponseBodyDataFacePicList {
	return s.FacePicList
}

func (s *QueryFaceUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryFaceUserResponseBodyData) GetParams() *string {
	return s.Params
}

func (s *QueryFaceUserResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *QueryFaceUserResponseBodyData) SetCustomUserId(v string) *QueryFaceUserResponseBodyData {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetFacePicList(v []*QueryFaceUserResponseBodyDataFacePicList) *QueryFaceUserResponseBodyData {
	s.FacePicList = v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetName(v string) *QueryFaceUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetParams(v string) *QueryFaceUserResponseBodyData {
	s.Params = &v
	return s
}

func (s *QueryFaceUserResponseBodyData) SetUserId(v string) *QueryFaceUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryFaceUserResponseBodyData) Validate() error {
	if s.FacePicList != nil {
		for _, item := range s.FacePicList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFaceUserResponseBodyDataFacePicList struct {
	// The MD5 hash of the photo.
	//
	// example:
	//
	// TestFaceMd5
	FaceMd5 *string `json:"FaceMd5,omitempty" xml:"FaceMd5,omitempty"`
	// The URL of the photo.
	//
	// example:
	//
	// https://test.com/xxx.jpg
	FaceUrl *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	// The list of face feature values.
	FeatureDTOList []*QueryFaceUserResponseBodyDataFacePicListFeatureDTOList `json:"FeatureDTOList,omitempty" xml:"FeatureDTOList,omitempty" type:"Repeated"`
}

func (s QueryFaceUserResponseBodyDataFacePicList) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserResponseBodyDataFacePicList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBodyDataFacePicList) GetFaceMd5() *string {
	return s.FaceMd5
}

func (s *QueryFaceUserResponseBodyDataFacePicList) GetFaceUrl() *string {
	return s.FaceUrl
}

func (s *QueryFaceUserResponseBodyDataFacePicList) GetFeatureDTOList() []*QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	return s.FeatureDTOList
}

func (s *QueryFaceUserResponseBodyDataFacePicList) SetFaceMd5(v string) *QueryFaceUserResponseBodyDataFacePicList {
	s.FaceMd5 = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicList) SetFaceUrl(v string) *QueryFaceUserResponseBodyDataFacePicList {
	s.FaceUrl = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicList) SetFeatureDTOList(v []*QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) *QueryFaceUserResponseBodyDataFacePicList {
	s.FeatureDTOList = v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicList) Validate() error {
	if s.FeatureDTOList != nil {
		for _, item := range s.FeatureDTOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFaceUserResponseBodyDataFacePicListFeatureDTOList struct {
	// The name of the algorithm.
	//
	// example:
	//
	// TestAlgorithmName
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// The provider of the algorithm.
	//
	// example:
	//
	// TestAlgorithmProvider
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// The version of the algorithm.
	//
	// example:
	//
	// TestAlgorithmVersion
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The MD5 hash of the face.
	//
	// example:
	//
	// TestFaceMd5
	FaceMd5 *string `json:"FaceMd5,omitempty" xml:"FaceMd5,omitempty"`
}

func (s QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) GetAlgorithmVersion() *string {
	return s.AlgorithmVersion
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) GetFaceMd5() *string {
	return s.FaceMd5
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmName(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmName = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmProvider(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmProvider = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetAlgorithmVersion(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.AlgorithmVersion = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetErrorCode(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.ErrorCode = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetErrorMessage(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) SetFaceMd5(v string) *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList {
	s.FaceMd5 = &v
	return s
}

func (s *QueryFaceUserResponseBodyDataFacePicListFeatureDTOList) Validate() error {
	return dara.Validate(s)
}
