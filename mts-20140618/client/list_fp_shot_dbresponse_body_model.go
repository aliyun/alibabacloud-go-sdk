// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFpShotDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFpShotDBList(v *ListFpShotDBResponseBodyFpShotDBList) *ListFpShotDBResponseBody
	GetFpShotDBList() *ListFpShotDBResponseBodyFpShotDBList
	SetNonExistIds(v *ListFpShotDBResponseBodyNonExistIds) *ListFpShotDBResponseBody
	GetNonExistIds() *ListFpShotDBResponseBodyNonExistIds
	SetRequestId(v string) *ListFpShotDBResponseBody
	GetRequestId() *string
}

type ListFpShotDBResponseBody struct {
	// The media fingerprint libraries.
	FpShotDBList *ListFpShotDBResponseBodyFpShotDBList `json:"FpShotDBList,omitempty" xml:"FpShotDBList,omitempty" type:"Struct"`
	// The IDs of the media fingerprint libraries that do not exist.
	NonExistIds *ListFpShotDBResponseBodyNonExistIds `json:"NonExistIds,omitempty" xml:"NonExistIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFpShotDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotDBResponseBody) GoString() string {
	return s.String()
}

func (s *ListFpShotDBResponseBody) GetFpShotDBList() *ListFpShotDBResponseBodyFpShotDBList {
	return s.FpShotDBList
}

func (s *ListFpShotDBResponseBody) GetNonExistIds() *ListFpShotDBResponseBodyNonExistIds {
	return s.NonExistIds
}

func (s *ListFpShotDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFpShotDBResponseBody) SetFpShotDBList(v *ListFpShotDBResponseBodyFpShotDBList) *ListFpShotDBResponseBody {
	s.FpShotDBList = v
	return s
}

func (s *ListFpShotDBResponseBody) SetNonExistIds(v *ListFpShotDBResponseBodyNonExistIds) *ListFpShotDBResponseBody {
	s.NonExistIds = v
	return s
}

func (s *ListFpShotDBResponseBody) SetRequestId(v string) *ListFpShotDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFpShotDBResponseBody) Validate() error {
	if s.FpShotDBList != nil {
		if err := s.FpShotDBList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistIds != nil {
		if err := s.NonExistIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFpShotDBResponseBodyFpShotDBList struct {
	FpShotDB []*ListFpShotDBResponseBodyFpShotDBListFpShotDB `json:"FpShotDB,omitempty" xml:"FpShotDB,omitempty" type:"Repeated"`
}

func (s ListFpShotDBResponseBodyFpShotDBList) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotDBResponseBodyFpShotDBList) GoString() string {
	return s.String()
}

func (s *ListFpShotDBResponseBodyFpShotDBList) GetFpShotDB() []*ListFpShotDBResponseBodyFpShotDBListFpShotDB {
	return s.FpShotDB
}

func (s *ListFpShotDBResponseBodyFpShotDBList) SetFpShotDB(v []*ListFpShotDBResponseBodyFpShotDBListFpShotDB) *ListFpShotDBResponseBodyFpShotDBList {
	s.FpShotDB = v
	return s
}

func (s *ListFpShotDBResponseBodyFpShotDBList) Validate() error {
	if s.FpShotDB != nil {
		for _, item := range s.FpShotDB {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFpShotDBResponseBodyFpShotDBListFpShotDB struct {
	// The description of the media fingerprint library.
	//
	// example:
	//
	// The library is a text fingerprint library.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the media fingerprint library.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	FpDBId *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	// The model ID of the media fingerprint library. A value of **11*	- indicates that the library is a text fingerprint library. A value of **12*	- indicates that the library is a video fingerprint library. A value of **13*	- indicates that the library is an audio fingerprint library. A value of **14*	- indicates that the library is an image fingerprint library.
	//
	// example:
	//
	// 11
	ModelId *int32 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The name of the media fingerprint library.
	//
	// example:
	//
	// test-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the media fingerprint library. Default value: **offline**. ****Valid values:
	//
	// 	- **offline**: The media fingerprint library is offline.
	//
	// 	- **active**: The media fingerprint library is online.
	//
	// 	- **paused**: The media fingerprint library is paused.
	//
	// 	- **deleted**: The media fingerprint library is deleted.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFpShotDBResponseBodyFpShotDBListFpShotDB) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotDBResponseBodyFpShotDBListFpShotDB) GoString() string {
	return s.String()
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) GetDescription() *string {
	return s.Description
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) GetFpDBId() *string {
	return s.FpDBId
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) GetModelId() *int32 {
	return s.ModelId
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) GetName() *string {
	return s.Name
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) GetStatus() *string {
	return s.Status
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) SetDescription(v string) *ListFpShotDBResponseBodyFpShotDBListFpShotDB {
	s.Description = &v
	return s
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) SetFpDBId(v string) *ListFpShotDBResponseBodyFpShotDBListFpShotDB {
	s.FpDBId = &v
	return s
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) SetModelId(v int32) *ListFpShotDBResponseBodyFpShotDBListFpShotDB {
	s.ModelId = &v
	return s
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) SetName(v string) *ListFpShotDBResponseBodyFpShotDBListFpShotDB {
	s.Name = &v
	return s
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) SetStatus(v string) *ListFpShotDBResponseBodyFpShotDBListFpShotDB {
	s.Status = &v
	return s
}

func (s *ListFpShotDBResponseBodyFpShotDBListFpShotDB) Validate() error {
	return dara.Validate(s)
}

type ListFpShotDBResponseBodyNonExistIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListFpShotDBResponseBodyNonExistIds) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotDBResponseBodyNonExistIds) GoString() string {
	return s.String()
}

func (s *ListFpShotDBResponseBodyNonExistIds) GetString_() []*string {
	return s.String_
}

func (s *ListFpShotDBResponseBodyNonExistIds) SetString_(v []*string) *ListFpShotDBResponseBodyNonExistIds {
	s.String_ = v
	return s
}

func (s *ListFpShotDBResponseBodyNonExistIds) Validate() error {
	return dara.Validate(s)
}
