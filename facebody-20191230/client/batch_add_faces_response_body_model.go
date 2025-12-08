// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddFacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BatchAddFacesResponseBodyData) *BatchAddFacesResponseBody
	GetData() *BatchAddFacesResponseBodyData
	SetRequestId(v string) *BatchAddFacesResponseBody
	GetRequestId() *string
}

type BatchAddFacesResponseBody struct {
	Data *BatchAddFacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C46A46D0-3263-181A-A1EE-0901E4595390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchAddFacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponseBody) GetData() *BatchAddFacesResponseBodyData {
	return s.Data
}

func (s *BatchAddFacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchAddFacesResponseBody) SetData(v *BatchAddFacesResponseBodyData) *BatchAddFacesResponseBody {
	s.Data = v
	return s
}

func (s *BatchAddFacesResponseBody) SetRequestId(v string) *BatchAddFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddFacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchAddFacesResponseBodyData struct {
	FailedFaces   []*BatchAddFacesResponseBodyDataFailedFaces   `json:"FailedFaces,omitempty" xml:"FailedFaces,omitempty" type:"Repeated"`
	InsertedFaces []*BatchAddFacesResponseBodyDataInsertedFaces `json:"InsertedFaces,omitempty" xml:"InsertedFaces,omitempty" type:"Repeated"`
}

func (s BatchAddFacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponseBodyData) GetFailedFaces() []*BatchAddFacesResponseBodyDataFailedFaces {
	return s.FailedFaces
}

func (s *BatchAddFacesResponseBodyData) GetInsertedFaces() []*BatchAddFacesResponseBodyDataInsertedFaces {
	return s.InsertedFaces
}

func (s *BatchAddFacesResponseBodyData) SetFailedFaces(v []*BatchAddFacesResponseBodyDataFailedFaces) *BatchAddFacesResponseBodyData {
	s.FailedFaces = v
	return s
}

func (s *BatchAddFacesResponseBodyData) SetInsertedFaces(v []*BatchAddFacesResponseBodyDataInsertedFaces) *BatchAddFacesResponseBodyData {
	s.InsertedFaces = v
	return s
}

func (s *BatchAddFacesResponseBodyData) Validate() error {
	if s.FailedFaces != nil {
		for _, item := range s.FailedFaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InsertedFaces != nil {
		for _, item := range s.InsertedFaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchAddFacesResponseBodyDataFailedFaces struct {
	// example:
	//
	// ClientError.IllegalArgument
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/imgsearch/demo/xxxx.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// not found the db=test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s BatchAddFacesResponseBodyDataFailedFaces) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesResponseBodyDataFailedFaces) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) GetCode() *string {
	return s.Code
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) GetImageURL() *string {
	return s.ImageURL
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) GetMessage() *string {
	return s.Message
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) SetCode(v string) *BatchAddFacesResponseBodyDataFailedFaces {
	s.Code = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) SetImageURL(v string) *BatchAddFacesResponseBodyDataFailedFaces {
	s.ImageURL = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) SetMessage(v string) *BatchAddFacesResponseBodyDataFailedFaces {
	s.Message = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) Validate() error {
	return dara.Validate(s)
}

type BatchAddFacesResponseBodyDataInsertedFaces struct {
	// example:
	//
	// 16350536
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/imgsearch/demo/xxxx.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 99.79581
	QualitieScore *float32 `json:"QualitieScore,omitempty" xml:"QualitieScore,omitempty"`
}

func (s BatchAddFacesResponseBodyDataInsertedFaces) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesResponseBodyDataInsertedFaces) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) GetFaceId() *string {
	return s.FaceId
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) GetImageURL() *string {
	return s.ImageURL
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) GetQualitieScore() *float32 {
	return s.QualitieScore
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) SetFaceId(v string) *BatchAddFacesResponseBodyDataInsertedFaces {
	s.FaceId = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) SetImageURL(v string) *BatchAddFacesResponseBodyDataInsertedFaces {
	s.ImageURL = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) SetQualitieScore(v float32) *BatchAddFacesResponseBodyDataInsertedFaces {
	s.QualitieScore = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) Validate() error {
	return dara.Validate(s)
}
